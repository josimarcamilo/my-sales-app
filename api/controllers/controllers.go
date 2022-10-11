package controllers

import (
	"api-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Companys(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Companys)
}

func Company(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, company := range models.Companys {
		if strconv.Itoa(company.Id) == id {
			json.NewEncoder(w).Encode(company)
		}
	}
}
