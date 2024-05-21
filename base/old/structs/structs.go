package structs

import (
	"base/structs/models"
	"fmt"
)



func UseStructs() {

	fmt.Println("*******************")
	fmt.Println("***  begin structs ***")

	newUser := models.User{
		Name:   "majeed",
		Age:    36,
		Gender: "male",
		City:   "Istanbul",
		State:  true,
	}

	fmt.Println(newUser)

	var newUserTwo models.User

	newUserTwo.Name = "John"
	newUserTwo.Age = 21

	fmt.Println(newUserTwo)

	var newProduct models.Product

	newProduct.Name = "pen"
	newProduct.Price = 10

	fmt.Println(newProduct)


	var productList []models.Product;

	productList = append(productList, newProduct)

	fmt.Println(productList)
	fmt.Println("len of product list is : ",len(productList))


	for _,product := range productList{
		fmt.Println(product.Name)
	}


	fmt.Println("***  end structs ***")
	fmt.Println("*******************")
}
