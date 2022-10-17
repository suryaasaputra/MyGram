package main

import (
	"fmt"
	"mygram/database"
	"mygram/router"
)

func main() {
	r := router.StartApp()
	err := database.StartDB()
	if err != nil {
		fmt.Println("Error starting database: ", err)
		return
	}
	r.Run(":8080")
}
