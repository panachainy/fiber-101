package it

import "fiber-101/database"

func CleanDatabase() {
	database.DBConn.Exec("DELETE FROM products")
}
