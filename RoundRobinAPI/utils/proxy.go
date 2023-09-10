package utils

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func NewReverseProxyHandler(lb *RoundRobinLoadBalancer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Choose a healthy backend server using the load balancing algorithm
		backendURL, err := lb.GetHealthyBackend()
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "No healthy backend available"})
			return
		}

		// Measure the response time of the backend server
		startTime := time.Now()
		target, err := url.Parse(backendURL)
		if err != nil {
			log.Printf("Error parsing URL: %v", err)
			c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to parse backend URL"})
			return
		}

		// Create a reverse proxy handler
		proxy := httputil.NewSingleHostReverseProxy(target)

		// Update the request to use the selected backend server
		c.Request.URL.Scheme = target.Scheme
		c.Request.URL.Host = target.Host
		c.Request.Header.Set("X-Forwarded-Host", c.Request.Header.Get("Host"))

		// Serve the request via the reverse proxy
		proxy.ServeHTTP(c.Writer, c.Request)

		// Calculate the response time
		responseTime := time.Since(startTime)

		// Update the backend server's response time and mark it as slow if needed
		for i, server := range lb.backendServers {
			if server.Address == backendURL {
				lb.mu.Lock()
				lb.backendServers[i].ResponseTime = responseTime
				lb.backendServers[i].IsSlow = responseTime > 5*time.Second // if response time > 5 seconds, mark as slow
				lb.mu.Unlock()
				break
			}

		}
	}
}
