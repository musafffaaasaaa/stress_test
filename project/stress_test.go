package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Define the number of concurrent users and requests
	numUsers := 10
	numRequests := 100

	// Create a WaitGroup to wait for all Goroutines to finish
	var wg sync.WaitGroup

	// Start stress testing
	for i := 0; i < numUsers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stressTest(numRequests)
		}()
	}

	// Wait for all Goroutines to finish
	wg.Wait()
	fmt.Println("Stress testing completed.")
}

func stressTest(numRequests int) {
	// Replace with your shop website's URL
	targetURL := "https://your-shop-website.com"

	for i := 0; i < numRequests; i++ {
		startTime := time.Now()

		// Send an HTTP GET request to the website
		resp, err := http.Get(targetURL)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		defer resp.Body.Close()

		// Calculate and print response time
		elapsedTime := time.Since(startTime)
		fmt.Printf("Request #%d - Status: %s, Elapsed Time: %v\n", i+1, resp.Status, elapsedTime)

		// You can add additional logic to parse the response, check for expected content, etc.
	}
}
