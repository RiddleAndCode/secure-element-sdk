package main

import (
	"C"
	"encoding/hex"
	"fmt"

	element "github.com/riddleandcode/secure-element-sdk/golang/secureelement"
)

func main() {
	var res C.int
	randomBytes := element.Random()
	fmt.Print("Here are your random bytes:\n")
	fmt.Printf("%s\n", hex.Dump(randomBytes))
}
