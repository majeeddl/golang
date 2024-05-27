package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}

	fmt.Printf("%v\n", p)       // {1 2}
	fmt.Printf("%+v\n", p)      // {x:1 y:2}
	fmt.Printf("%#v\n", p)      // main.point{x:1, y:2}
	fmt.Printf("type: %T\n", p) // type: main.point

	// ***************************
	fmt.Printf("%t\n", true) // true
	fmt.Printf("%d\n", 123)  // 123

	fmt.Printf("%b\n", 14) // 1110
	fmt.Printf("%c\n", 33) // \! (character representation)

	fmt.Printf("%x\n", 456)  // 1c8 (hex encoding)
	fmt.Printf("%f\n", 78.9) // 78.900000 (default width, precision)

	fmt.Printf("%e\n", 123400000.0) // 1.234000e+08 (scientific notation)
	fmt.Printf("%E\n", 123400000.0) // 1.234000E+08 (scientific notation)

	fmt.Printf("%s\n", "\"string\"") // "string" (string representation)
	fmt.Printf("%q\n", "\"string\"") // "\"string\"" (double-quoted string)

	fmt.Printf("%x\n", "hex this") // 6865782074686973 (hex encoding)

	fmt.Printf("%p\n", &p) // 0xc0000140f0 (pointer)

	fmt.Printf("|%6d|%6d|\n", 12, 345)       // |    12|   345| (width 6, right justified)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) // |  1.20|  3.45| (width 6, precision 2)

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // |1.20  |3.45  | (width 6, precision 2, left justified)

	fmt.Printf("|%6s|%6s|\n", "foo", "b") // |   foo|     b| (width 6, right justified)

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // |foo   |b     | (width 6, left justified)

	s := fmt.Sprintf("a %s", "string") // a string
	fmt.Println(s)                     // a string

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error") // io: an error
}
