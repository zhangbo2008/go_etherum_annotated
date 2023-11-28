package ethereum

import (
	"testing"
	"unsafe"
)

func Test1(t *testing.T) {
	a := "dafasdf"
	b := a[:]

	println(unsafe.Pointer(&a)) //0xc000063e78
	println(unsafe.Pointer(&b)) //0xc000063e68

}
