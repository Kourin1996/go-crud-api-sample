package main

import (
	"log"

	v1 "github.com/Kourin1996/go-crud-api-sample/internal/api/v1"
)

func main() {
	if err := v1.Start(); err != nil {
		log.Fatalf("error happened: %s", err.Error())
		return
	}
}
