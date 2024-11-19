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
	log.Println("server is running")
	http.ListenAndServe(serverAddress, apps)
}
