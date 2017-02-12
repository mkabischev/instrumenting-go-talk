package main

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/pborman/uuid"
)

func userID(r *http.Request) int {
	return rand.Int()
}

func doSmth() error {
	if rand.Float64() < 0.5 {
		return errors.New("some error")
	}

	return nil
}

func contextWithLogger(ctx context.Context, l *logrus.Entry) context.Context {
	return context.WithValue(ctx, "logger", l)
}

func loggerFromContext(ctx context.Context) *logrus.Entry {
	return ctx.Value("logger").(*logrus.Entry)
}

// USER1 OMIT
func userMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := loggerFromContext(r.Context()) // HLx
		l = l.WithField("userID", userID(r)) // HLx

		ctx := contextWithLogger(r.Context(), l)
		next(w, r.WithContext(ctx))
	}
}
// USER2 OMIT

// LOG1 OMIT
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logrus.New().WithField("requestID", uuid.New()) // HLx

		ctx := contextWithLogger(r.Context(), l)
		next(w, r.WithContext(ctx))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	l := loggerFromContext(r.Context())
	l.Infof("request received: %v", r.URL) // HLx

	if err := doSmth(); err != nil {
		l.Errorf("error occured: %v", err) // HLx
	}
}
// LOG2 OMIT

func main() {
	http.HandleFunc("/", loggerMiddleware(userMiddleware(handler))) // HLx
	http.ListenAndServe(":8080", nil)
}


// START OMIT


// END OMIT
