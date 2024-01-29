package main

import (
	"context"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"github.com/rs/zerolog/log"
)

func main() {
	// Set up Google Cloud Storage client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create GCS client")
	}

	// GCS bucket and object name
	bucketName := "your-gcs-bucket-name"
	objectName := "logs.log"

	// Open the local log file
	file, err := os.Open("logs.log")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}
	defer file.Close()

	// Upload log file to GCS
	wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		log.Fatal().Err(err).Msg("Failed to upload log file to GCS")
	}
	if err := wc.Close(); err != nil {
		log.Fatal().Err(err).Msg("Failed to close GCS writer")
	}

	log.Info().Msg("Logging to GCS...")
}
