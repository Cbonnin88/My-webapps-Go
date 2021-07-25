package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is my first Web app with Go")
	})

	http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request){
		http.Redirect(w, r, "https://tutorialedge." +
			"net/golang/creating-simple-web-server-with-golang/", 301 )
	})

	http.HandleFunc("/video", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w,r, "https://www.youtube." +
			"com/watch?v=ldCD9yLa6tM&t=480s", 301)

	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}