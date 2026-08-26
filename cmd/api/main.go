package main

// server is build
import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

// error 404 bec no end point is defined.
