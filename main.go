package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is not found in the environement.")
	} else {
		fmt.Println("Port:", portString)
	}

	router := chi.NewRouter()

	log.Printf("Server starting on port %v", portString)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
