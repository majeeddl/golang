package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	/*
		For example, rand.Intn returns a random int n, 0 <= n < 100.
	*/
	fmt.Println(rand.IntN(100), ",") // 81,
	fmt.Println(rand.IntN(100))      // 87
	fmt.Println()

	/*
		rand.Flost64 returns a random float64 f, 0.0 <= f < 1.0.
	*/
	fmt.Println(rand.Float64()) // 0.6645600532184904

	/*
		This can be used to generate random floats in other ranges, for example 5.0 <= f' < 10.0.
	*/
	fmt.Println((rand.Float64()*5)+5, ",") // 7.188570935934901,
	fmt.Println((rand.Float64() * 5) + 5)  // 7.123187485356329
	fmt.Println()

	/*
		If you want a known seed, create a new rand.
		Source and pass it into the New constructor.
		NewPCG creates a new PCG source that requires a seed of two uint64 numbers.
	*/
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
