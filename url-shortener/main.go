package main

import (
	"fmt"
	"url-shortener/db"
	"url-shortener/routes"
)

func main() {

	db.InitDB()
	r := routes.InitRoutes()
	fmt.Println("Server is running at port 8080")
	r.Run(":8080")
}
