package crack

import (
	"crypto/md5"
	"fmt"
)

func HashMD5(word string) string {
	hashBytes := md5.Sum([]byte(word))  // hash the word (as bytes)
	return fmt.Sprintf("%x", hashBytes) // convert bytes to readable hex
}
