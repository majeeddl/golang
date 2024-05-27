package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var println = fmt.Println

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	/*
		First we’ll look at encoding basic data types to JSON strings.
		Here are some examples for atomic values.
	*/
	bolB, _ := json.Marshal(true)
	println(string(bolB)) // true

	intB, _ := json.Marshal(1)
	println(string(intB)) // 1

	fltB, _ := json.Marshal(2.34)
	println(string(fltB)) // 2.34

	strB, _ := json.Marshal("gopher")
	println(string(strB)) // "gopher"

	/*
		And here are some for slices and maps, which encode to JSON arrays and objects as you’d expect.
	*/
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	println(string(slcB)) // ["apple","peach","pear"]

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	println(string(mapB)) // {"apple":5,"lettuce":7}

	/*
		The JSON package can automatically encode your custom data types.
		It will only include exported fields in the encoded output and will by default use those names as the JSON keys.
	*/
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	println(string(res1B)) // {"Page":1,"Fruits":["apple","peach","pear"]}

	/*
		You can use tags on struct field declarations to customize the encoded JSON key names.
		Check the definition of response2 above to see an example of such tags.
	*/
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	println(string(res2B)) // {"page":1,"fruits":["apple","peach","pear"]}

	/*
		Now let’s look at decoding JSON data into Go values.
		Here’s an example for a generic data structure.
	*/
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	// Here’s the actual decoding, and a check for associated errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	println(dat) // map[num:6.13 strs:[a b]]

	/*
		In order to use the values in the decoded map, we'll need to convert them to their appropriate type.
		For example here we convert the value in num to the expected float64 type.
	*/
	num := dat["num"].(float64)
	println(num) // 6.13

	/*
		Accessing nested data required a series of conversions.
	*/
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	println(str1) // a

	/*
		We can also decode JSON into custom data types.
		This has the advantages of adding additional type-safety to our programs and eliminating the need for type assertions when accessing the decoded data.
	*/
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	println(res)           // {1 [apple peach]}
	println(res.Fruits[0]) // apple

	/*
		In the examples above we always used bytes and strings as intermediates between the data and JSON representation on standard out.
		We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	*/
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d) // {"apple":5,"lettuce":7}
}
