package main

import (
	"fmt"
	"go-api-rest/models"
	"go-api-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Name 1", History: "History 1"},
		{Id: 2, Name: "Name 2", History: "History 2"},
	}
	fmt.Println("Iniciando servidor Go")
	routes.HandleRequest()
}
