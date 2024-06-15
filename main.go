package main

import (
	"log"
)

func main() {

	db, err := NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", db)
	server.Serve()
}
