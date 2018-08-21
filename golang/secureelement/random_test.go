package secureelement_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	element "github.com/riddleandcode/secure-element-sdk/golang/secureelement"
)

func TestRandom(t *testing.T) {
	zeroBytes := make([]byte, 32)
	randomBytes := element.Random()

	// fmt.Printf("zeroBytes: %d\n%s\n", len(zeroBytes), hex.Dump(zeroBytes))
	// fmt.Printf("randomBytes:\n%s\n", hex.Dump(randomBytes))

	if len(randomBytes) != 32 {
		t.Errorf("The length is incorrect, got: %d, want: %d.", len(randomBytes), 32)
	}

	if bytes.Equal(zeroBytes, randomBytes) {
		t.Errorf("The random bytes are all zero. Not very random, is it?\n%s", hex.Dump(randomBytes))
	}

	// test stuff here...
}
