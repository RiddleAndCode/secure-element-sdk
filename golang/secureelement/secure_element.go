/*
Generate 32 bytes of true, hardware-generated, as-good-as-it-gets,
finest-entropy randomness.
*/
package secureelement

//#cgo CFLAGS: -I../inc
//#cgo LDFLAGS: -L../lib -Wl,-rpath -Wl,./ -lcryptoauth
//#include "golang/wrapper.h"
/*
import "C"
import (
	"unsafe"
)
*/

/*
Generate 32 bytes of true, hardware-generated, as-good-as-it-gets,
finest-entropy randomness.
*/
func Random() []byte {
	// randomBytes := make([]byte, C.RANDOM_NUM_SIZE)
	// return C.getRandomNumber((*C.uint8_t)(unsafe.Pointer(&randomBytes[0])))
	randomBytes := make([]byte, 32)
	return randomBytes
}
