package main

// Example program demonstrating the usage of Random.
import (
	"encoding/hex"
	"fmt"
	"testing"

	// import the go-cryptoauthlib package into element
	element "github.com/riddleandcode/go-cryptoauthlib"
)

func main() {
	// element.Random() returns 32 bytes of random data
	randomBytes := element.Random()
	// print the data to the console
	fmt.Printf("randomBytes:\n%s\n", hex.Dump(randomBytes))



	var res int
	var random_number = element.Random()
	fmt.Printf("Here's your random: %d\n", res)
	fmt.Printf("%s\n", hex.Dump(random_number))


	var digest = element.Random()
	fmt.Printf("Here's your digest: %d\n", res)
	fmt.Printf("%s\n", hex.Dump(digest))

	var public_key []byte
	res, public_key = element.GetPublicKey()
	fmt.Printf("Here's your public key: %d\n", res)
	fmt.Printf("%s\n", hex.Dump(public_key))

	var signature []byte
	res, signature = element.SignDigest(digest)
	fmt.Printf("Here's your signature: %d\n", res)
	fmt.Printf("%s\n", hex.Dump(signature))

	var verified bool
	res, verified = element.VerifySignedDigest( digest, signature, public_key )
	fmt.Printf("Here's your verification: %d\n", res)
	fmt.Printf("%t \n",verified )

}
