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
		fmt.Println("Usage: hmac <message> <secret>")
	} else {
		fmt.Println(computeHmac(args[1], args[2]))
	}
}

func computeHmac(message string, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(message))
	return hex.EncodeToString(hash.Sum(nil))
}
