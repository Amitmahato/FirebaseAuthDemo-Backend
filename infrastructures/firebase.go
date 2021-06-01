package infrastructures

import (
	"context"
	"path/filepath"

	firebase "firebase.google.com/go/v4" //upgrade to v4
	"google.golang.org/api/option"
)

// InitializeFirebase initlise the firebase auth client
func InitializeFirebase() *firebase.App {
	ctx := context.Background()
	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKey.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic("Firebase load error")
	}
	return app
}
