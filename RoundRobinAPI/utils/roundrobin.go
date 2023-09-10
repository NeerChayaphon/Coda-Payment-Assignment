package utils

import (
	"fmt"
	"net"
	"net/url"
	"sync"
	"time"
)

type BackendServerInfo struct {
	Address      string
	IsSlow       bool
	ResponseTime time.Duration
}

type RoundRobinLoadBalancer struct {
	backendServers []BackendServerInfo
	currentIndex   int
	mu             sync.Mutex
}

// NewRoundRobinLoadBalancer creates a new RoundRobinLoadBalancer instance.
func NewRoundRobinLoadBalancer(backendServers []string) *RoundRobinLoadBalancer {
	lb := &RoundRobinLoadBalancer{}
	for _, server := range backendServers {
		lb.backendServers = append(lb.backendServers, BackendServerInfo{
			Address:      server,
			IsSlow:       false,
			ResponseTime: 0,
		})
	}
	return lb
}

// GetHealthyBackend selects the next healthy backend server in a round-robin fashion.
func (lb *RoundRobinLoadBalancer) GetHealthyBackend() (string, error) {
	// Lock to prevent concurrent access to currentIndex
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// Iterate through backend servers in a round-robin fashion
	for i := 0; i < len(lb.backendServers); i++ {
		if lb.backendServers[lb.currentIndex].IsSlow {
			lb.currentIndex = (lb.currentIndex + 1) % len(lb.backendServers)
			continue
		}

		backendURL := lb.backendServers[lb.currentIndex]
		lb.currentIndex = (lb.currentIndex + 1) % len(lb.backendServers)

		// Check if the backend server is healthy (responds to health check)
		if isHealthy(backendURL.Address) {
			return backendURL.Address, nil
		}
	}

	// if all backend servers are slow, we will use the slow server for round robin
	// It best to return slow response than no response
	for i := 0; i < len(lb.backendServers); i++ {
		backendURL := lb.backendServers[lb.currentIndex]
		lb.currentIndex = (lb.currentIndex + 1) % len(lb.backendServers)

		if isHealthy(backendURL.Address) {
			return backendURL.Address, nil
		}
	}

	// No healthy backend server found
	return "", fmt.Errorf("no healthy backend server available")
}

func isHealthy(backendURL string) bool {
	u, err := url.Parse(backendURL)
	if err != nil {
		return false
	}

	hostPort := u.Hostname() + ":" + u.Port()

	conn, err := net.DialTimeout("tcp", hostPort, 2*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}
