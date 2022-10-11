package main

import (
	"api-rest/models"
	"api-rest/routes"
	"fmt"
)

func main() {
	models.Companys = []models.Company{
		{Id: 1, Description: "Josimar SA"},
		{Id: 2, Description: "Josimar Camilo SA"},
	}

	fmt.Println("api rodando")
	routes.HandleRequest()
}
