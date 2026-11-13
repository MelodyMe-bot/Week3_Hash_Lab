package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"password-cracker/utils/crack"
)

func main() {
	targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
	wordlistPath := "wordlists/nord_vpn.txt"

	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false
	count := 0

	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		count++

		hash := crack.HashSHA512(word)
		fmt.Printf("Trying: %-20s => %s\n", word, hash)

		if hash == targetHash {
			fmt.Printf("\n[+] PASSWORD FOUND: %s\n", word)
			found = true
			break
		}
	}

	fmt.Println("Total words tried:", count)

	if !found {
		fmt.Println("\n[-] Password not found in wordlist.")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
