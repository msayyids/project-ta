package main

import (
	"fmt"
	"log"
	"net/http"
	"project-ta/router"
)

func main() {
	apps := router.NewRouter()

	port := 8080

	serverAddress := fmt.Sprintf(":%d", port)
	log.Printf("server is running on %s", serverAddress)
	if err := http.ListenAndServe(serverAddress, apps); err != nil {
		log.Fatal(err)
	}
}
