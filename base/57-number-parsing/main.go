package main

import (
	"fmt"
	"strconv"
)

var println = fmt.Println

func main() {

	/*
		With ParseFloat, this 64 tells how many bits of precision to parse.
	*/
	f, _ := strconv.ParseFloat("1.234", 64)
	println(f) // 1.234

	/*
		For ParseInt, the 0 means infer the base from the string. 64 requires that the result fit in 64 bits.
	*/
	i, _ := strconv.ParseInt("123", 0, 64)
	println(i) // 123

	/*
		ParseInt will recognize hex-formatted numbers.
	*/
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	println(d) // 456

	/*
		A ParseUint is also available.
	*/
	u, _ := strconv.ParseUint("789", 0, 64)
	println(u) // 789

	/*
		Atoi is a convenience function for basic base-10 int parsing.
	*/
	k, _ := strconv.Atoi("135")
	println(k) // 135

	/*
		Parse function return an error on bad input.
	*/
	_, e := strconv.Atoi("wat")
	println(e) // strconv.Atoi: parsing "wat": invalid syntax
}
