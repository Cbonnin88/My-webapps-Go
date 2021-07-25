package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Welcome to my first Go Web app")
	})

	http.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w,r, "https://tutorialedge." +
			"net/golang/basic-rest-api-go-fiber/", 301)
	})

	http.HandleFunc("/video", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w,r, "https://www.youtube." +
			"com/watch?v=ldCD9yLa6tM&t=480s", 301)
	})

	log.Fatal(http.ListenAndServe(":1988", nil))
}

// Test
// Test two
