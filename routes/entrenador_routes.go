package routes

import (
	"ENTRENADORES_GO_CRUD/controllers"
	"github.com/gorilla/mux"
)

func RegisterEntrenadorRoutes(r *mux.Router) {
	r.HandleFunc("/entrenadores", controllers.GetAllEntrenadores).Methods("GET")
	r.HandleFunc("/entrenadores/{id}", controllers.GetEntrenadorByID).Methods("GET")
	r.HandleFunc("/entrenadores", controllers.CreateEntrenador).Methods("POST")
	r.HandleFunc("/entrenadores/{id}", controllers.UpdateEntrenador).Methods("PUT")
	r.HandleFunc("/entrenadores/{id}", controllers.DeleteEntrenador).Methods("DELETE")
}
