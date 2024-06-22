package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateShortURL(originalURL string) string {
	hash := sha256.Sum256([]byte(originalURL))
	return base64.URLEncoding.EncodeToString(hash[:6]) // Shortened to 6 characters
}
