package main

import (
	"database/sql"
	"fmt"
)

const PORT = ":8001"

type App struct {
	DB *sql.DB
}

func main() {
	fmt.Printf("Starting auth service on port %v", PORT)
}
