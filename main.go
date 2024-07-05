package main

import (
	"fmt"
	"go-api-rest/database"
	"go-api-rest/routes"
)

func main() {
	database.ConnectDatabase()
	fmt.Println("Iniciando servidor Go")
	routes.HandleRequest()
}
