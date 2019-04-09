package main

import (
	"fmt"
	"net/http"
	"time"
)

// CustomRoundTripper is a custom implementation of http.RoundTripper
type CustomRoundTripper struct {
	Transport http.RoundTripper
}

func (c *CustomRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()

	fmt.Printf("Requesting %s %s\n", req.Method, req.URL)

	resp, err := c.Transport.RoundTrip(req)

	elapsed := time.Since(start)
	fmt.Printf("Received response in %v\n", elapsed)

	return resp, err
}

func main() {
	client := &http.Client{
		Transport: &CustomRoundTripper{
			Transport: http.DefaultTransport,
		},
	}

	resp, err := client.Get("https://www.example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
}
