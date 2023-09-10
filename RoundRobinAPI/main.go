package main

import (
	"log"
	"os"

	"github.com/NeerChayaphon/CodaAssignment/RoundRobinAPI/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	router := gin.Default()

	// List of backend servers
	backendServers := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}

	// Load balancing algorithm with health checking and timeout handling
	lb := utils.NewRoundRobinLoadBalancer(backendServers)

	//  Route for the reverse proxy
	router.POST("/", utils.NewReverseProxyHandler(lb))

	log.Fatal(router.Run(":" + port))
}
