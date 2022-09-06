package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct{
	FirstName string `json:"first_name"`
	Age int `json:"age"`
}

func main()  {
	

	myJson := `[
		{
		"first_name": "majeed",
		"age" : 36
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson),&unmarshalled)


	if err != nil {
		log.Println("Error unmarshalling json",err)
	}

	log.Printf("unmarshalled : %v", unmarshalled)


	//write json from a strcut
	var mySlice []Person

	var m1 Person
	m1.FirstName ="John"
	m1.Age = 22


	mySlice = append(mySlice, m1)

	newJson,err := json.MarshalIndent(mySlice,"", "    ")

	if err != nil {
		log.Println("error marshalling ", err)
	}

	fmt.Println(string(newJson))
}