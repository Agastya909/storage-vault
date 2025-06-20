package main

import (
	"storage/handlers/ping"

	"github.com/go-chi/chi/v5"
)

func InitRoutes(r *chi.Mux, handlers Handlers) {
	r.Get("/ping", ping.Ping)

	r.Route("/aws", func(r chi.Router) {
		r.Post("/bucket/{bucketName}", handlers.AwsHandler.CreateBucket)
		r.Get("/bucket/{bucketName}", handlers.AwsHandler.GetBucketInfo)

		r.Get("/url/presigned/{bucketName}/{objectKey}", handlers.AwsHandler.GetPresignedUrl)
	})
}
