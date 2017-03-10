package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// The two functions are used for the convertion between tring and []byte
func main() {
	str := "Test for convertion"
	fmt.Println(S2B(str))
	fmt.Println(B2S(S2B(str)))
}

func B2S(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

func S2B(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}
