package main

import config "notification-api/src/configuration"

func main() {
	config.ConnectDb().Exec("DROP TABLE IF EXISTS users;")
}
