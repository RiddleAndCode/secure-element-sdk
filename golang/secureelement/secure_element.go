/*
Generate 32 bytes of true, hardware-generated, as-good-as-it-gets,
finest-entropy randomness.
*/
package secureelement

//#cgo CFLAGS: -I../../inc
//#cgo LDFLAGS: -lcryptoauth
//#include "./wrapper.h"
import "C"
import (
	"unsafe"
)

/*
Generate 32 bytes of true, hardware-generated, as-good-as-it-gets,
finest-entropy randomness.
*/
func Random() []byte {
	randomBytes := make([]byte, C.RANDOM_NUM_SIZE)
	C.getRandomNumber((*C.uint8_t)(unsafe.Pointer(&randomBytes[0])))
	return randomBytes
}
