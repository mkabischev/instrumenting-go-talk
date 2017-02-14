package main

import "github.com/Sirupsen/logrus"

func main() {
	base := logrus.New()
	base.Formatter = &logrus.JSONFormatter{}

	// START OMIT
	logger := logrus.New().
		WithField("customer_id", "73be6b5f-9d78-43a4-991b-b5a5cb9630d3").
		WithField("provider_id", "0a170dfd-52d9-494a-88e7-fe73b48f900f")

	// ..

	logger = logger.WithField("ride_id", "5abc0868-8376-49ea-a66a-d42bc3937852")

	// ..

	logger.Info("ride started")
	// END OMIT
}
