package main

import (
	"context"
	"log"

	"github.com/tdzdslippen/suplan/internal/app"
)

func main() {
	a, err := app.New(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
