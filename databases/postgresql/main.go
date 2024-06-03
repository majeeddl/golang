package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Test struct {
	Name string `db:"name"`
}

func main() {
	//connect to a PostgreSQL database
	// Replace the connection details (user, dbname, password, host) with your own
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable password=postgis@2022 host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	test := Test{}

	// Query the database
	rows, _ := db.Queryx("SELECT name FROM test")

	for rows.Next() {
		err := rows.StructScan(&test)
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Println(test)

}
