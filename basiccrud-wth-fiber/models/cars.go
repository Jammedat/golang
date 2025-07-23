package models

import (
	"cars/config"
	"database/sql"
	"fmt"
)

type Car struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Model string    `json:"model"`
	Brand string    `json:"brand"`
	Year  int       `json:"year"`
	Price float64   `json:"price"`
}

var Cars = make(map[int]Car)

func (c *Car) Insert() {
	query := `INSERT INTO cars (name, model, brand, year, price) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := config.DB.QueryRow(query, c.Name, c.Model, c.Brand, c.Year, c.Price).
		Scan(&c.ID)
	if err != nil {
		fmt.Printf("error inserting car: %v\n", err)
	}

}

func (c *Car) Get() error {
	query := `SELECT name, model, brand, year, price FROM cars WHERE id = $1`
	err := config.DB.QueryRow(query, c.ID).
		Scan(&c.Name, &c.Model, &c.Brand, &c.Year, &c.Price)
	if err != nil {

		if err == sql.ErrNoRows {
			fmt.Printf("error getting car: %v\n", err)
			return err

		}
	}
	return nil
}

func (c *Car) Update() {
	// query := `INSERT INTO cars (name, model, brand, year, price) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	query := `UPDATE cars SET name = $1, model = $2, brand = $3, year = $4, price = $5 WHERE id = $6`
	_, err := config.DB.Exec(query, c.Name, c.Model, c.Brand, c.Year, c.Price, c.ID)
	if err != nil {
		fmt.Printf("error updating car: %v", err)
	}

}

func (c *Car) Delete() {
	query := `DELETE FROM cars where id = $1`
	_, err := config.DB.Exec(query, c.ID)
	if err != nil {
		fmt.Printf("error deleting the car with the id: %v, error: %v\n",c.ID, err)
	}
}