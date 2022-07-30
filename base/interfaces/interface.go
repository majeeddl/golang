package interfaces

import "fmt"

type IProduct interface {
	list()
	add()
	counter(i int) int
}

type handler struct{
	DB string
	num int
}

func (h handler ) list(){
	fmt.Println("print list")
}

func (h handler) add(){

}

func (h handler)counter(a int) int{
	a++
	return a
}

func RunInterfaces() {

	fmt.Println("*******************")
	fmt.Println("***  begin interfaces ***")

	var product IProduct
	product = new(handler)
	// product = &handler{}
	// product = handler{}
	product.list()

	a := 10

	b := product.counter(a)

	fmt.Println(b)

	fmt.Println("*******************")
	fmt.Println("***  end interfaces ***")
}
