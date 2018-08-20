package main
//#cgo CFLAGS: -I../inc
//#cgo LDFLAGS: -L../lib -Wl,-rpath -Wl,./ -lcryptoauth
//#include "wrapper.h"
import "C"
import (
        "encoding/hex"
        "fmt"
        "unsafe"
)


func main() {
	var res C.int
	rnd_number := make([]byte, C.RANDOM_NUM_SIZE)
	res = C.getRandomNumber((*C.uint8_t)(unsafe.Pointer(&rnd_number[0])))
	fmt.Printf("Here's your random: %d\n", res)
	fmt.Printf("%s\n", hex.Dump(rnd_number))


	digest := make([]byte, C.RANDOM_NUM_SIZE)
	res = C.getRandomNumber((*C.uint8_t)(unsafe.Pointer(&digest[0])))
	fmt.Printf("Here's your digest: %d\n", res)
	fmt.Printf("%s\n", hex.Dump(digest))


	const pub_keysize = 64
	pub_key := make([]byte, C.PUBLIC_KEY_SIZE)
	res = C.getPublicKey((*C.uint8_t)(unsafe.Pointer(&pub_key[0])))
	fmt.Printf("Here's your public key: %d\n", res)
	fmt.Printf("%s\n", hex.Dump(pub_key))

	signature :=make([]byte, C.SIGNATURE_SIZE)

	res = C.sign( (*C.uint8_t)(unsafe.Pointer(&digest[0])), (*C.uint8_t)(unsafe.Pointer(&signature[0])))
	fmt.Printf("Here's your signature: %d\n", res)
	fmt.Printf("%s\n", hex.Dump(signature))
	var verified C.bool

	res = C.verify_extern( (*C.uint8_t)(unsafe.Pointer(&digest[0])), (*C.uint8_t)(unsafe.Pointer(&signature[0])), (*C.uint8_t)(unsafe.Pointer(&pub_key[0])), &verified )
	fmt.Printf("Here's your verification: %d\n", res)
	fmt.Printf("%t \n",verified )



}
