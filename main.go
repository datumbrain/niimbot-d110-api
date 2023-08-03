package main

import (
	"log"
	"net/http"
	"os"
)

const port = ":8769"

func main() {
	os.Mkdir("images", 0750)

	http.Handle("/print", withLogging(printHandler))

	log.Println("listening at port", port)
	log.Println(http.ListenAndServe(port, nil))
}
