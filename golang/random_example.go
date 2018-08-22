package main

// Example program demonstrating the usage of Random.
import (
	"encoding/hex"
	"fmt"
	"testing"

	// import the go-cryptoauthlib package into element
	element "github.com/riddleandcode/go-cryptoauthlib"
)

func main(t *testing.T) {
	// element.Random() returns 32 bytes of random data
	randomBytes := element.Random()
	// print the data to the console
	fmt.Printf("randomBytes:\n%s\n", hex.Dump(randomBytes))
}
