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
	r.HandleFunc("/companys", controllers.Companys)

	log.Fatal(http.ListenAndServe(":8000", r))
}
