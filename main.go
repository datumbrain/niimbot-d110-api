package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const port = ":8769"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	os.Mkdir("images", 0750)

	http.Handle("/print", HandleAuth(withLogging(printHandler)))

	log.Println("listening at port", port)
	log.Println(http.ListenAndServe(port, nil))
}
