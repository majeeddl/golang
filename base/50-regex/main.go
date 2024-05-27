package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// This tests whether a pattern matches a string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) // true

	/*
		Above we used a string pattern directly,
		but for other regexp tasks you’ll need to Compile an optimized Regexp struct.
	*/
	r, _ := regexp.Compile("p([a-z]+)ch") // Compile the pattern

	/*
		Many methods are available on these structs.
		Here’s a match test like we saw earlier
	*/
	fmt.Println(r.MatchString("peach")) // true

	/*
		This finds the match for the regexp.
	*/
	fmt.Println(r.FindString("peach punch")) // peach

	/*
		This also finds the first match but returns the start and end indexes for the match instead of the matching text.
	*/
	fmt.Println("idx:", r.FindStringIndex("peach punch")) // [0 5]

	/*
		The Submatch variants include information about both the whole-pattern matches and the submatches within those matches.
		For example this will return information for both p([a-z]+)ch and ([a-z]+).
	*/
	fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]

	/*
		Similarly this will return information about the indexes of matches and submatches.
	*/
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

	/*
		The All variants of these functions apply to all matches in the input, not just the first.
		For example to find all matches for a regexp.
	*/
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]

	/*
		These All variants are available for the other functions we saw above as well.
	*/
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1)) // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	/*
		Providing a non-negative integer as the second argument to these functions will limit the number of matches.
	*/
	fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]

	/*
		Our examples above had string arguments and used names like MatchString.
		We can also provide []byte arguments and drop String from the function name.
	*/
	fmt.Println(r.Match([]byte("peach"))) // true

	/*
		When creating global variables with regular expressions you can use the MustCompile variation of Compile.
		MustCompile panics instead of returning an error, which makes it safer to use for global variables.
	*/
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r) // regexp: p([a-z]+)ch

	/*
		The regexp package can also be used to replace subsets of strings with other values.
	*/
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

	/*
		The Func variant allows you to transform matched text with a given function.
	*/
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
