package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/harrisonwjs/senpaislist-backend/db/sqlc"
)

// Server serves HTTP requests for anime bookmarking service
// perhaps more in the future ...
type Server struct {
	store  db.Store    // interact with the db
	router *gin.Engine // send each API request to the correct handler for processing
}

// func NewServer(store *db.Store) *Server {
// 	server := &Server{store: store} // new Server object
// 	router := gin.Default()         // new router

// 	// ROUTES
// 	router.POST("/anime", server.createAnime)

// 	// assign the router object to the server.router
// 	server.router = router
// 	return server
// }
