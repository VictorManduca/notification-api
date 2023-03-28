package main

import (
	config "notification-api/src/configuration"
)

func main() {
	config.ConnectDb().Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(100), email VARCHAR(100));")
}
