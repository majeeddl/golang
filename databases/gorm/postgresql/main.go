package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	dsn := "user=postgres dbname=postgres sslmode=disable password=postgis@2022 host=localhost"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	user := User{Name: "Majeed", Age: 38}

	result := db.Create(&user)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("User created successfully", result.RowsAffected)

	getUsers := db.Take(&user)

	if getUsers.Error != nil {
		panic(getUsers.Error)
	}

	fmt.Println("User retrieved successfully", user)

	db.First(&user)

	user.Name = "Masoumeh"
	user.Age = 30
	db.Save(&user)
}
