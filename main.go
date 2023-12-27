package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// Structure to store product information
type Product struct {
	Name  string  `json:"title"`
	Price float64 `json:"price"`
}

// APIResponse represents the API response structure
type APIResponse struct {
	Products []Product `json:"products"`
}

// Function to fetch product information
func fetchProduct(url string, wg *sync.WaitGroup, productsDetails chan APIResponse, mutex *sync.Mutex) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error crawling URL:", url, err)
		return
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	var apiResponse APIResponse
	if err := json.Unmarshal([]byte(body), &apiResponse); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Protect shared product channel with a mutex
	mutex.Lock()
	productsDetails <- apiResponse // Send product information to the channel
	mutex.Unlock()
}

func main() {
	productURLs := []string{
		"https://dummyjson.com/products/category/smartphones",
		"https://dummyjson.com/products/category/laptops",
	}

	products := make(chan APIResponse)
	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	// Spawn multiple goroutines to crawl product information concurrently
	for _, url := range productURLs {
		go fetchProduct(url, wg, products, mutex)
	}

	var prod []Product
	totalResponseCount := 0

	for {
		product := <-products // Receive product information from the channel
		totalResponseCount++
		prod = append(prod, product.Products...)
		fmt.Println("products are as follows : ")
		for _, p := range prod {
			fmt.Println(p)
		}
		if totalResponseCount == 2 {
			fmt.Println("All responses received, closing channel...")
			close(products)
			break
		}
	}

}
