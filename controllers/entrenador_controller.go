package controllers

import (
	"API_ENTRENADORES/config"
	"API_ENTRENADORES/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Helper JSON
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}


func GetAllEntrenadores(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id_entrenadores, id_adminitrador, cedula, descripcion, especialidad,
		gimnasio_id, anios_experiencia, foto_url, aprovacion_entrenador, hoja_vida,
		calificacion_prom, activo, fecha_modificacion, fecha_creacion
		FROM entrenadores WHERE 1=1`

	cedula       := r.URL.Query().Get("cedula")
	especialidad := r.URL.Query().Get("especialidad")

	if cedula != "" {
		query += " AND cedula ILIKE '%" + cedula + "%'"
	}
	if especialidad != "" {
		query += " AND especialidad ILIKE '%" + especialidad + "%'"
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var lista []models.Entrenador
	for rows.Next() {
		var e models.Entrenador
		rows.Scan(
			&e.IDEntrenadores, &e.IDAdministrador, &e.Cedula, &e.Descripcion,
			&e.Especialidad, &e.GimnasioID, &e.AniosExperiencia, &e.FotoURL,
			&e.AprovacionEntrenador, &e.HojaVida, &e.CalificacionProm,
			&e.Activo, &e.FechaModificacion, &e.FechaCreacion,
		)
		lista = append(lista, e)
	}
	respondJSON(w, 200, lista)
}
