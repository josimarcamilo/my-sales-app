package routes

import (
	"api-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/companys", controllers.Companys).Methods("Get")
	r.HandleFunc("/companys/{id}", controllers.Company).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))
}
