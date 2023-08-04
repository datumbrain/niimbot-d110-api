package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const port = ":8769"
const serviceName = "Rasp-API"

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		for {
			select {
			case sig := <-sigs:
				log.Printf("Received signal: %v. Stopping daemon.", sig)
				done <- true
			}
		}
	}()

	// Run the HTTP server continuously
	server()

	done <- true
	<-done
}

func server() {
	os.Mkdir("images", 0750)

	http.Handle("/print", withLogging(printHandler))

	log.Println("listening at port", port)
	log.Println(http.ListenAndServe(port, nil))
}
