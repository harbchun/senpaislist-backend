package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	// years := [3]string{"2019", "2020", "2020"}
	// seasons := [4]string{"winter", "spring", "summer", "fall"}
	// client := &http.Client{}

	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 4; j++ {

	// 		req, _ := http.NewRequest("GET", "https://kitsu.io/api/edge/anime/?filter[seasonYear]="+years[i]+"&filter[season]="+seasons[j], nil)

	// 		resp, err := client.Do(req)

	// 		if err != nil {
	// 			fmt.Println("Errored when sending request to the server")
	// 			return
	// 		}

	// 		defer resp.Body.Close()
	// 		resp_body, _ := ioutil.ReadAll(resp.Body)

	// 		fmt.Println(resp.Status)
	// 		fmt.Println(string(resp_body))
	// 	}
	// }
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://kitsu.io/api/edge/anime?filter[seasonYear]=2020&filter[season]=winter", nil)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(resp.Status)
	// fmt.Println(string(resp_body))
	fmt.Fprintf(w, string(resp_body))
}

func main() {
	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("../templates/*")

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.
	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	})

	// Start serving the application
	router.Run()
}
