package routes

import (
	"github.com/gorilla/mux"
	"github.com/medali/ramo-server/controllers"
)



func Helper(router *mux.Router){
	router.HandleFunc("/helper/openai", controllers.HelperController).Methods("POST")
	router.HandleFunc("/helper/gemeni", controllers.GemeniHelperController).Methods("POST")
	
} 