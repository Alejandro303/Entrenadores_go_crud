package routes

import (
	"API_ENTRENADORES/controllers"
	"github.com/gorilla/mux"
)

func RegisterDocumentoRoutes(r *mux.Router) {
	r.HandleFunc("/documentos", controllers.GetAllDocumentos).Methods("GET")
	r.HandleFunc("/documentos/{id}", controllers.GetDocumentoByID).Methods("GET")
	r.HandleFunc("/documentos", controllers.CreateDocumento).Methods("POST")
	r.HandleFunc("/documentos/{id}", controllers.UpdateDocumento).Methods("PUT")
	r.HandleFunc("/documentos/{id}", controllers.DeleteDocumento).Methods("DELETE")
}
