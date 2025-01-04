# GoLang-URLShortener
Simple bare bones URL shortener written in Go

**Short URL Generation**
The generateShortHash method generates a SHA-256 hash of the original URL and encodes it using Base64. It takes the first 8 characters of the encoded string as the shortened URL.

**Collision Detection**
The CreateShortURL method checks for collisions in the in-memory database. If a collision occurs, it regenerates the short URL by hashing a combination of the existing short URL and the original URL until a unique short URL is found.

**Thread Safety**
The package uses a read-write mutex (sync.RWMutex) to handle concurrent access to the in-memory map.
