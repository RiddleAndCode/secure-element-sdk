# SDK for I²C communication with the RIDDLE&CODE Secure Element

This SDK includes precompiled cryptoauth libraries, header files, and wrapper code for several languages to connect to the RIDDLE&CODE Secure Element via the I²C protocol on various devices.

The SDK has the following structure

- inc           The include files needed to utilize the library.
- lib            precompiled libraries (for several architectures: e.g. arm raspberry pi).
- golang    The wrapper for the Go Language as well as some golang sample code.

## The C library
The C libraries are available as shared objects and can be loaded into several languages. The corresponding header files are within the inc folder.

## Go wrapper
Using the go wrapper comes with two installation steps

- Copy the c library from the lib folder to a path contained in LDPATH of the system or to the golang folder if the command is run from there. 
- Add the pkg to your GOPATH 

The golang/install.sh script just executes these two steps as shown below
```go
sudo cp ../bin/libcryptoauth.so /usr/lib
cp -R src $GOPATH/.
cp -R pkg $GOPATH/.
```

Thereafter, the library can be utilized as needed. 
The concept of the provided Go API is shown in the golang/random_example.go example file.


### Example

```go
package main

// Example program demonstrating the usage of Random.
import (
	"encoding/hex"
	"fmt"
	"testing"

	// import the go-cryptoauthlib package into element
	element "github.com/riddleandcode/cryptoauthlib"
)

func main(t *testing.T) {
	// element.Random() returns 32 bytes of random data
	randomBytes := element.Random()
	// print the data to the console
	fmt.Printf("randomBytes:\n%s\n", hex.Dump(randomBytes))
}
```
