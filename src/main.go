package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Senpai's List")
}

func main() {
	fmt.Println("Senpai's List")

	http.HandleFunc("/", homepage)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
