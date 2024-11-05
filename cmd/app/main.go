package main

import "github.com/Kei-K23/user-management-system-api/db"

func main() {
	db.ConnectDB()

	// Close the db when program stop
	defer db.Pool.Close()
}
