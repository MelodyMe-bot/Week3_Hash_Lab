//go:build ignore

package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

type hashFunc func([]byte) string

type algo struct {
	name string
	fn   hashFunc
}

var algos = []algo{
	{"MD5", func(b []byte) string { h := md5.Sum(b); return hex.EncodeToString(h[:]) }},
	{"SHA1", func(b []byte) string { h := sha1.Sum(b); return hex.EncodeToString(h[:]) }},

	{"SHA256", func(b []byte) string { h := sha256.Sum256(b); return hex.EncodeToString(h[:]) }},
	{"SHA512", func(b []byte) string { h := sha512.Sum512(b); return hex.EncodeToString(h[:]) }},
	{"SHA3", func(b []byte) string { h := sha3.Sum256(b); return hex.EncodeToString(h[:]) }},
}

func proofMe(txt1, txt2 string) {
	for _, a := range algos {
		outputA := a.fn([]byte(txt1))
		outputB := a.fn([]byte(txt2))

		result := "No Match!"
		if outputA == outputB {
			result = "Match!"
		}
		fmt.Printf("Hash (%s) : Output A / Output B => %s\n", a.name, result)
		fmt.Printf("   A: %s\n   B: %s\n", outputA, outputB)
	}
}

func readLine(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	line, _ := reader.ReadString('\n')
	return strings.TrimRight(line, "\r\n")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("========Name + Hashing Program========")
	input1 := readLine(reader, "Please input value 1: ")
	input2 := readLine(reader, "Please input value 2: ")

	proofMe(input1, input2)
}
