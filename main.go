package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/takama/daemon"
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

	if len(os.Args) > 1 {
		cmd := os.Args[1]
		switch cmd {
		case "install":
			d, err := daemon.New(serviceName, "Your daemon description", daemon.SystemDaemon)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			_, err = d.Install()
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			log.Println("Daemon installed successfully.")
			os.Exit(0)
		case "remove":
			d, err := daemon.New(serviceName, "Your daemon description", daemon.SystemDaemon)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			_, err = d.Remove()
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			log.Println("Daemon removed successfully.")
			os.Exit(0)
		case "start":
			d, err := daemon.New(serviceName, "Your daemon description", daemon.SystemDaemon)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			_, err = d.Start()
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			log.Println("Daemon started successfully.")
			os.Exit(0)
		case "stop":
			d, err := daemon.New(serviceName, "Your daemon description", daemon.SystemDaemon)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			_, err = d.Stop()
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			log.Println("Daemon stopped successfully.")
			os.Exit(0)
		case "status":
			d, err := daemon.New(serviceName, "Your daemon description", daemon.SystemDaemon)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			status, err := d.Status()
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			log.Printf("Daemon status: %v", status)
			os.Exit(0)
		}
	}

	// This part will be executed when running the binary normally without any command.
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
