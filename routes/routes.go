package routes

import (
	"go-api-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.Personalities).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreatePersonalities).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.ShowPersonalities).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonalities).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.UpdatePersonalities).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
