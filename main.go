package main

import (
	"log"
	"user/vault/cmd"
)

func main() {
	// create + run new APIServer
	server := cmd.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
