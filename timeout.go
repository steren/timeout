package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("I received a new request and will now print a log entry every 10 seconds until the request times out.")

	i := 0

	// Print logs forever
	for {
		time.Sleep(10 * time.Second)
		i++
		log.Print(i)
	}
}

func main() {
	log.Print("I started. Waiting for a request on /")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
