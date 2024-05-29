package main

import (
	"crypto/sha256"
	"fmt"
)

var println = fmt.Println

func main() {
	s := "sha256 this string"

	h := sha256.New()

	/*
		Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.
	*/
	h.Write([]byte(s))

	/*
		This gets the finalized hash result as a byte slice.
		 The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed.
	*/
	bs := h.Sum(nil)

	println(s)
	println(bs)
}
