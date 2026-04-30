package routes

import (
	"API_ENTRENADORES/controllers"
	"github.com/gorilla/mux"
)

func RegisterSolicitudRoutes(r *mux.Router) {
	r.HandleFunc("/solicitudes", controllers.GetAllSolicitudes).Methods("GET")
	r.HandleFunc("/solicitudes/{id}", controllers.GetSolicitudByID).Methods("GET")
	r.HandleFunc("/solicitudes", controllers.CreateSolicitud).Methods("POST")
	r.HandleFunc("/solicitudes/{id}", controllers.UpdateSolicitud).Methods("PUT")
	r.HandleFunc("/solicitudes/{id}", controllers.DeleteSolicitud).Methods("DELETE")
}
