package routes

import (
	"github.com/blazte/VentasJunior/controllers"
	"github.com/gorilla/mux"
)

//SetLoginRouter router para login
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
