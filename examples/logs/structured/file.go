package main

import (
	"context"

	"junolab.net/ms/core/logger"
)

type Logger interface{}

// START OMIT
const key int = 0

<<<<<<< HEAD
func ToContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, key, l)
}

=======
>>>>>>> 550e9efcfdbc7871354677fb09fdae3e976cc41a
func FromContext(ctx context.Context) Logger {
	return ctx.Value(key).(Logger)
}

<<<<<<< HEAD
=======
func ToContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, key, l)
}

>>>>>>> 550e9efcfdbc7871354677fb09fdae3e976cc41a
// END OMIT

type Request struct{}
type Response struct{}

// HANDLER_START OMIT

func SomeHandler(ctx context.Context, r Request) (*Response, error) {
	l := logger.FromContext(ctx)  // HLx
	l.Debug("something happened") // HLx

	// ...

}

// HANDLER_END OMIT

// PREPARE_START OMIT
l = l.WithField("request_id", NewRequestID()) // HLx
ctx := logger.WithContext(ctx, l) // HLx

// ...

res, err := SomeHandler(ctx, req)

// PREPARE_END OMIT
