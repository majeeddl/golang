package main

import "fmt"

type bot interface{
	getGreeding() string
}

type englishBot struct{}
type spanishBot struct{}


func main() {
	eb := englishBot{}
	sb := spanishBot{}


	printGreeing(eb)
	printGreeing(sb)
}


func printGreeing (b bot){
	fmt.Println(b.getGreeding())
}

func (englishBot) getGreeding() string{
	return "Hi English Bot"
}

func (spanishBot) getGreeding() string{
	return "Hi Spanish Bot"
}
