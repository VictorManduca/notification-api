package main

import (
	"fmt"
	"log"
	"net/http"

	config "notification-api/src/configuration"
	"notification-api/src/controllers"
)

func main() {
	fmt.Println("Starting on port: " + config.Env().PORT)

	controllers.Routes()
	log.Fatal(http.ListenAndServe(":"+config.Env().PORT, nil))
}
