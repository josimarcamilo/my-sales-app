package main

import (
	"api-rest/models"
	"api-rest/routes"
	"fmt"
)

func main() {
	models.Companys = []models.Company{
		{Description: "Josimar SA"},
		{Description: "Josimar Camilo SA"},
	}

	fmt.Println("api rodando")
	routes.HandleRequest()
}
