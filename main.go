package main

import (
	"forum/auth"
	"forum/post"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	auth.RegisterRoutes()
	post.RegisterRoutes()
	port := os.Getenv("PORT")
	log.Println("Starting server on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
