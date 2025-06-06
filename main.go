package main

import (
	"fmt"
	"log"
	"net/http"
	"user/vault/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// create file server
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// create router
	router := gin.New()
	router.GET("/api", api.Middleware)
	router.Run()

	// start server on local port
	fmt.Println("Starting server a port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
