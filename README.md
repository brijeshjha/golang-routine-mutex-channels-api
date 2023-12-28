# Golang Concurrency Example: Fetching Product Information

## Overview

This Golang app fetches product details concurrently from two URLs (smartphones and laptops) using goroutines and channels. It ensures efficient data retrieval through parallel processing.

## Features

- **Concurrent Fetching:** Utilizes goroutines for parallel fetching of product information.
- **Channel Communication:** Ensures safe data exchange through channels with mutex protection.
- **JSON Decoding:** Structured decoding of JSON responses for easy data manipulation.

## Usage

1. **Clone Repository:**
   ```bash
   git clone https://github.com/brijeshjha/golang-routine-mutex-channels-api.git
   cd golang-routine-mutex-channels-api
   go run .
