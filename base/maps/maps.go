package maps

import "fmt"

func PrintMaps() {

	fmt.Println("*******************")
	fmt.Println("***  begin maps ***")

	m := map[string]string{
		"a" : "1",
		"b" : "2",
	}

	fmt.Println(m)

	fmt.Println("=> add new item to maps")

	m["c"] = "3"

	fmt.Println(m)

	fmt.Println("=> delete item from maps")

	delete(m,"c")

	fmt.Println(m)

	fmt.Println("=> loops")

	i := 0

	// for{
	// 	i++
	// 	fmt.Println(i)
	// }

	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("=>for range")

	for index,value := range m {
		fmt.Println(index,value)
	}

	fmt.Println("***  end maps ***")
	fmt.Println("*******************")
}