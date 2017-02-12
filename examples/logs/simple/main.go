package main

import (
	"log"
	"math/rand"
	"net/http"
)

func func1() int { return rand.Int() }
func func2() int { return rand.Int() }

// START OMIT
func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("request received: ", r.URL) // HLx
	// ...

	w.WriteHeader(200)
	// ...
	log.Print("request processed") // HLx
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// END OMIT
