package main

import (
	"os"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	router := gin.Default()
    router.GET("/", homePage)

	log.Printf("Listening on port %s\n\n", port)
    router.Run(":"+port)
}

func homePage(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, gin.H{"message": "Elastic Beanstalk is working!"})
}