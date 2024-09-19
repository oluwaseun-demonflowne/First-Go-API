package main

import (
	"log"

	"github.com/oluwaseun-demonflowne/ecom/cmd/api"
)

func main() {
	server := api.NewApiServer(":8080",nil)
	// server.Run()
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}