package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {

	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)

	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{"Bob", 31},
		Person{"John", 42},
		Person{"Michael", 17},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})

	fmt.Println(people)
}
