package main

import (
	"fmt"
	s "strings"
)

var print = fmt.Println

func main() {

	print("Contains:  ", s.Contains("test", "es"))
	print("Count:     ", s.Count("test", "t"))
	print("HasPrefix: ", s.HasPrefix("test", "te"))
	print("HasSuffix: ", s.HasSuffix("test", "st"))
	print("Index:     ", s.Index("test", "e"))
	print("Join:      ", s.Join([]string{"a", "b"}, "-"))
	print("Repeat:    ", s.Repeat("a", 5))
	print("Replace:   ", s.Replace("foo", "o", "0", -1))
	print("Replace:   ", s.Replace("foo", "o", "0", 1))
	print("Split:     ", s.Split("a-b-c-d-e", "-"))
	print("ToLower:   ", s.ToLower("TEST"))
	print("ToUpper:   ", s.ToUpper("test"))
	print()

	print("Len: ", len("hello"))
	print("Char: ", "hello"[1])
}
