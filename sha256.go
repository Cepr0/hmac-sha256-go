package main

import (
	"os"
	"fmt"
	"crypto/sha256"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: sha256 <1st Arg> <2nd Arg>")
		return
	}
	arg1 := os.Args[1]
	arg2 := os.Args[2]

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s%s", arg1, arg2)))

	fmt.Printf("%x", h.Sum(nil))
}
