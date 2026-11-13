package main

import (
	"bufio"
	"fmt"
	"os"

	"password-cracker/utils/crack"
)

func main() {
	targetHash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	wordlistPath := "wordlists/passwords/nord_vpn.txt"

	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false

	for scanner.Scan() {
		word := scanner.Text()
		hash := crack.HashSHA1(word)

		fmt.Printf("Trying: %-20s => %s\n", word, hash)

		if hash == targetHash {
			fmt.Printf("\n[+] PASSWORD FOUND: %s\n", word)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("\n[-] Password not found in wordlist.")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
