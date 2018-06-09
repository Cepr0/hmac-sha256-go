package main

import (
	"os"
	"fmt"
	"crypto/sha256"
	"crypto/hmac"
	"encoding/hex"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: hmac <message> <secret>\nIf an argument contains spaces then use double quotes around it.\n")
	} else {
		fmt.Println(computeHmac(trimQuotes(args[1]), trimQuotes(args[2])))
	}
}

func computeHmac(message string, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(message))
	return hex.EncodeToString(hash.Sum(nil))
}

func trimQuotes(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}