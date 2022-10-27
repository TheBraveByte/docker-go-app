package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(wr http.ResponseWriter, rq *http.Request) {
	fmt.Fprintln(wr, "Hello guys welcome to U-library")
}

func main() {
	fmt.Println("Staring a simple web application")
	
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", homeHandler)

	srv := http.Server{
		Addr:    ":4040",
		Handler: mux,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Println("error starting go-app server")
	}
}
