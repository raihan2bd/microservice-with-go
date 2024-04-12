package main

import (
	"context"
	"fmt"

	"github.com/raihan2bd/microservice-with-go/application"
)

func main() {
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("error starting server: %v", err)
	}
}
