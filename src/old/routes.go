// routes.go

package main

func initializeRoutes() {
	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.

	// Handle the index route
	router.GET("/", showIndexPage)

	// Handle GET requests at /anime/view/some_article_id
	router.GET("/anime/view/:anime_id", getAnime)
}
