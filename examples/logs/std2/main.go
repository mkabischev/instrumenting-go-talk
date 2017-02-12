package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/pborman/uuid"
)

func doSmth() error {
	if rand.Float64() < 0.5 {
		return errors.New("some error")
	}

	return nil
}

// LOG1 OMIT
func contextWithLogger(ctx context.Context, l *log.Logger) context.Context {
	return context.WithValue(ctx, "logger", l)
}

func loggerFromContext(ctx context.Context) *log.Logger {
	return ctx.Value("logger").(*log.Logger)
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := log.New(os.Stderr, uuid.New() + " ", log.LstdFlags) // HLx

		ctx := contextWithLogger(r.Context(), l) // HLx
	 	next(w, r.WithContext(ctx)) // HLx
	}
}

func main() {
	http.HandleFunc("/", loggerMiddleware(handler)) // HLx
	http.ListenAndServe(":8080", nil)
}

// LOG2 OMIT

// START OMIT
func handler(w http.ResponseWriter, r *http.Request) {
	l := loggerFromContext(r.Context()) // HLx

	l.Printf("request received: %v", r.URL) // HLx

	if err := doSmth(); err != nil {
		l.Printf("error occured: %v", err) // HLx
	}
}



// END OMIT
