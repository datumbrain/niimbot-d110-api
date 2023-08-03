package main

import (
	"log"
	"net/http"
)

const port = ":8769"

func main() {
	http.HandleFunc("/print", printHandler)

	log.Println("listening at port", port)
	log.Println(http.ListenAndServe(port, nil))
}
