package main

import (
	"database/sql"
	db2 "github.com/aryells/hexagonal-aluno/adapters/db"
	"github.com/aryells/hexagonal-aluno/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	productService.Create("Product Test 1", 10)
	productService.Create("Product Test 2", 20)
	product, _ := productService.Create("Product Test 3", 30)

	productService.Enable(product)
}
