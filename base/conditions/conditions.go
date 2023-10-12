package conditions

import "fmt"

func PrintConditions() {

	fmt.Println("*******************")
	fmt.Println("***  begin conditions ***")

	fmt.Println("***  if conditions ***")
	x := 10

	if x == 5 {
		fmt.Println(" x = 4")
	} else {
		fmt.Println(" x = 10")
	}

	fmt.Println("***  switch conditions ***")

	switch x {
	case 10:
		fmt.Println("x = 10")
	default:
		fmt.Println(" x is default")
	}

	fmt.Println("***  end conditions ***")
	fmt.Println("*******************")

}
