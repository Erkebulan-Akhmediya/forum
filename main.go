package main

import (
	"forum/auth"
	"forum/comment"
	"forum/file"
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
	file.RegisterRoutes()
	comment.RegisterRoutes()
	port := os.Getenv("PORT")
	log.Println("Starting server on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
