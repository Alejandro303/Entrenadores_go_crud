package main

import (
	"log"
	"net/http"

	"ENTRENADORES_GO_CRUD/config"
	"ENTRENADORES_GO_CRUD/routes"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	config.ConnectDB()

	r := mux.NewRouter()

	routes.RegisterEntrenadorRoutes(r)
	routes.RegisterDocumentoRoutes(r)
	routes.RegisterSolicitudRoutes(r)

	log.Println("Servidor corriendo en el puerto 8083")
	http.ListenAndServe(":8083", enableCORS(r))
}