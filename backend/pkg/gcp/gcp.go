package gcp

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/manaaan/ekolivs-oms/backend/pkg/env"
)

func InitFirestore() *firestore.Client {
	ctx := context.Background()
	projectId := env.Required("GOOGLE_CLOUD_PROJECT")
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalln("could not init firestore client", err.Error())
	}
	return client
}
