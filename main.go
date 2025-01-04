package main

import (
    "fmt"
    "log"
    "your_project/urlshortener"
)

func main() {
    // Create a new URL shortener instance
    shortener := urlshortener.New()

    // Shorten a URL
    shortURL, err := shortener.CreateShortURL("https://example.com")
    if err != nil {
        log.Fatalf("Error creating short URL: %v", err)
    }
    fmt.Println("Short URL:", shortURL)

    // Retrieve the original URL
    originalURL, err := shortener.GetOriginalURL(shortURL)
    if err != nil {
        log.Fatalf("Error retrieving original URL: %v", err)
    }
    fmt.Println("Original URL:", originalURL)
}
