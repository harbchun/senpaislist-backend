// handlers.anime.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	animes := getAllAnimes()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": animes,
		},
	)

}

func getAnime(c *gin.Context) {
	// Check if the anime ID is valid
	if animeID, err := strconv.Atoi(c.Param("anime_id")); err == nil {
		// Check if the anime exists
		if anime, err := getAnimeByID(animeID); err == nil {
			// Call the HTML method of the Context to render a template
			c.HTML(
				// Set the HTTP status to 200 (OK)
				http.StatusOK,
				// Use the index.html template
				"anime.html",
				// Pass the data that the page uses
				gin.H{
					"title":   anime.Title,
					"payload": anime,
				},
			)

		} else {
			// If the anime is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid anime ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
