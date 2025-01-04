// Package urlshortener provides a simple URL reduction algorithm with an in-memory database.
package urlshortener

import (
    "crypto/sha256"
    "encoding/base64"
    "errors"
    "sync"
)

// Shortener is the main struct that manages URL shortening with an in-memory map.
type Shortener struct {
    mu    sync.RWMutex
    store map[string]string // Maps short URLs to original URLs
}

// New creates and returns a new instance of Shortener.
func New() *Shortener {
    return &Shortener{
        store: make(map[string]string),
    }
}

// CreateShortURL generates a short URL from an original URL and stores it in the database.
// It handles collision by adjusting the hash in case of a conflict.
func (s *Shortener) CreateShortURL(originalURL string) (string, error) {
    shortURL := s.generateShortHash(originalURL)

    s.mu.Lock()
    defer s.mu.Unlock()

    // Handle collision by checking if the generated short URL already exists
    for {
        if _, exists := s.store[shortURL]; !exists {
            // No collision; store the short URL and break the loop
            s.store[shortURL] = originalURL
            break
        }
        // In case of collision, regenerate the hash with additional characters
        shortURL = s.generateShortHash(shortURL + originalURL)
    }

    return shortURL, nil
}

// GetOriginalURL retrieves the original URL corresponding to a given short URL.
// Returns an error if the short URL is not found.
func (s *Shortener) GetOriginalURL(shortURL string) (string, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    originalURL, exists := s.store[shortURL]
    if !exists {
        return "", errors.New("short URL not found")
    }

    return originalURL, nil
}

// generateShortHash creates a base64-encoded hash from the input string and returns the first 8 characters.
func (s *Shortener) generateShortHash(input string) string {
    hash := sha256.Sum256([]byte(input))
    return base64.URLEncoding.EncodeToString(hash[:])[:8]
}
