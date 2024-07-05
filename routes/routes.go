package routes

import (
	"go-api-rest/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.Personalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
