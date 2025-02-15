package fireauth

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() error {
	env := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	fmt.Println("enviroment")
	fmt.Println(env)
	opt := option.WithCredentialsFile(env)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	App = app

	return err
}

func VerifyIDToken(idToken string) (*auth.Token, error) {
	ctx := context.Background()
	client, err := App.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)
	return token, err
}
