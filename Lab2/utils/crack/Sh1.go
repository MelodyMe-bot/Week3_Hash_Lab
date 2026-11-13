package crack

import (
	"crypto/sha1"
	"fmt"
)

func HashSHA1(word string) string {
	hashBytes := sha1.Sum([]byte(word)) // hash the word (as bytes)
	return fmt.Sprintf("%x", hashBytes) // convert bytes to readable hex
}
