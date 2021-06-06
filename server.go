package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/harrisonwjs/senpaislist-backend/database"
	"github.com/harrisonwjs/senpaislist-backend/graph"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/airingInformation"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/anime"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/animesgenres"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/genre"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/season"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/statistic"
	"github.com/harrisonwjs/senpaislist-backend/graph/controller/year"
	"github.com/harrisonwjs/senpaislist-backend/graph/generated"
)

const defaultPort = "5001"

func main() {
	log.Printf("Starting up...")

	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:championsclub123@postgres:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Printf("Migrations failed...")
	} else {
		log.Printf("Migrations passed...")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := database.InitDB()
	defer db.Close()

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5001"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		AnimeController:             anime.Anime{DB: db},
		StatisticController:         statistic.Statistic{DB: db},
		AiringInformationController: airingInformation.AiringInformation{DB: db},
		GenreController:             genre.Genre{DB: db},
		YearController:              year.Year{DB: db},
		SeasonController:            season.Season{DB: db},
		AnimesGenresController:      animesgenres.AnimesGenres{DB: db},
	}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				log.Println(r.Host)
				return r.Host == "localhost:3000" || r.Host == "localhost:5001"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
