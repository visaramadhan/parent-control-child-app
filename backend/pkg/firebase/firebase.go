package firebase

import (
    "context"
    "firebase.google.com/go/v4"
    "firebase.google.com/go/v4/auth"
    "google.golang.org/api/option"
    "log"
)

var AuthClient *auth.Client

func InitFirebaseApp() {
    ctx := context.Background()
    opt := option.WithCredentialsFile("path/to/serviceAccountKey.json") // Ganti dengan path ke file kunci layanan Anda
    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
        log.Fatalf("error initializing app: %v\n", err)
    }

    AuthClient, err = app.Auth(ctx)
    if err != nil {
        log.Fatalf("error getting Auth client: %v\n", err)
    }
}
