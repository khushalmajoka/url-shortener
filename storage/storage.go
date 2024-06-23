package storage

import (
	"net/url"
	"sort"
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

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv

	for k, v := range domainCounts {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	topDomains := make(map[string]int)
	for i := 0; i < n && i < len(ss); i++ {
		topDomains[ss[i].Key] = ss[i].Value
	}

	return topDomains
}
