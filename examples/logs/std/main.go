package main

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
)

func doSmth() error {
	if rand.Float64() < 0.5 {
		return errors.New("some error")
	}

	return nil
}

// START OMIT
func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request received: %v", r.URL) // HLx

	if err := doSmth(); err != nil {
		log.Printf("error occured: %v", err) // HLx
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// END OMIT
