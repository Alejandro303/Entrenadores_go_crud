package controllers

import (
	"API_ENTRENADORES/config"
	"API_ENTRENADORES/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL documentos (filtro: entrenador_id)
func GetAllDocumentos(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, entrenador_id, identificacion, nombre_archivo, creado_en,
		activo, fecha_modificacion, fecha_creacion FROM entrenador_documentos WHERE 1=1`

	entrenadorID := r.URL.Query().Get("entrenador_id")
	if entrenadorID != "" {
		query += " AND entrenador_id=" + entrenadorID
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var lista []models.EntrenadorDocumento
	for rows.Next() {
		var d models.EntrenadorDocumento
		rows.Scan(
			&d.ID, &d.EntrenadorID, &d.Identificacion, &d.NombreArchivo,
			&d.CreadoEn, &d.Activo, &d.FechaModificacion, &d.FechaCreacion,
		)
		lista = append(lista, d)
	}
	respondJSON(w, 200, lista)
}