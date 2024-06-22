package storage

import (
	"net/url"
	"sync"
)

var (
	urlStore     = make(map[string]string)
	domainCounts = make(map[string]int)
	mu           sync.Mutex
)

func SaveURL(originalURL, shortURL string) {
	mu.Lock()
	defer mu.Unlock()

	urlStore[shortURL] = originalURL

	u, _ := url.Parse(originalURL)
	domainCounts[u.Host]++
}

func GetOriginalURL(shortURL string) string {
	mu.Lock()
	defer mu.Unlock()

	return urlStore[shortURL]
}

func GetTopDomains(n int) map[string]int {
	mu.Lock()
	defer mu.Unlock()

	// Here you would implement logic to return the top n domains by count
	// For simplicity, returning the entire map
	return domainCounts
}
