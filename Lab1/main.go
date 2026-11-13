package main

import (
	"bufio"
	"fmt"
	"os"

	"password-cracker/utils/crack"
)

func main() {
	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"
	wordlistPath := "wordlists/nord_vpn.txt"

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
		hash := crack.HashMD5(word)

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
