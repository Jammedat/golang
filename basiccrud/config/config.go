package config


import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	
)



var DB *sql.DB

func ConnectDB() {
	
	dsn := "user=postgres password=<password> database=cars_db sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		fmt.Println("error opening database:", err)
		panic(err)

	
	}

	if err := db.Ping(); err != nil {
		fmt.Println("error connecting database", err)
		panic(err)
	}

	fmt.Println("successfuly connected to the database")
	DB = db

}