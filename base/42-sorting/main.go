package main

import (
	"fmt"
	"slices"
)

func main() {

	sts := []string{"c", "a", "b"}

	slices.Sort(sts)
	fmt.Println(sts)

	ints := []int{3, 1, 2}

	slices.Sort(ints)
	fmt.Println(ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted:", s)

}
