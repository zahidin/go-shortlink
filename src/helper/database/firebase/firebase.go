package firebase

import (
	"context"
	"fmt"
	"shortlink/src/config"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

func NewFirebase(ctx context.Context) (*firebase.App, error) {
	opt := option.WithCredentialsFile(config.GetConfig("FILE_NAME_CREDENTIAL"))
	config := &firebase.Config{
		DatabaseURL: config.GetConfig("DATABASE_URL"),
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return app, nil
}

func FirebaseDatabase(ctx context.Context) (*db.Client, error) {
	firebase, err := NewFirebase(ctx)

	if err != nil {
		return nil, err
	}

	client, err := firebase.Database(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}
