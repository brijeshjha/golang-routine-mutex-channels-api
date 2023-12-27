# Golang Concurrency Example: Fetching Product Information

This Golang code demonstrates how to fetch product information from multiple URLs concurrently using channels and mutexes. It retrieves product lists for smartphones and laptops and merges the responses into a single list.

**Code Overview**
The main function initializes two URLs for fetching product information: https://dummyjson.com/products/category/smartphones and https://dummyjson.com/products/category/laptops. It creates a channel products to receive product information and a mutex mutex to ensure data integrity. Additionally, it declares a wait group wg to manage goroutine synchronization.

Multiple goroutines are spawned using a for loop, each responsible for fetching product information from one of the URLs. The fetchProduct function handles the fetching process: it sends an HTTP GET request to the specified URL, reads the response body, decodes the JSON data into an APIResponse struct, and sends the response to the products channel.

The main function retrieves product information from the channel using a loop. It appends each received product to a prod slice. Once two responses are received (corresponding to the smartphones and laptops categories), it closes the products channel and prints the merged product list.

**Channel and Mutex Usage**
The products channel serves as a communication medium between the fetchProduct goroutines and the main function. It allows the goroutines to send product information asynchronously, while the main function retrieves the information in a controlled manner.

The mutex mutex ensures that only one goroutine can access the products channel at a time, preventing data races and maintaining the integrity of the product list. This is crucial as the main function needs to merge the received responses without any conflicts.

This Golang code demonstrates the effectiveness of concurrency techniques in efficient data retrieval. By utilizing channels and mutexes, it effectively fetches product information from multiple URLs simultaneously, ensuring data consistency and streamlined output.