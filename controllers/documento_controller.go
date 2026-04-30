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


func GetDocumentoByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var d models.EntrenadorDocumento

	err := config.DB.QueryRow(
		`SELECT id, entrenador_id, identificacion, nombre_archivo, creado_en,
		activo, fecha_modificacion, fecha_creacion FROM entrenador_documentos WHERE id=$1`, id,
	).Scan(
		&d.ID, &d.EntrenadorID, &d.Identificacion, &d.NombreArchivo,
		&d.CreadoEn, &d.Activo, &d.FechaModificacion, &d.FechaCreacion,
	)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Documento no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, d)
}


func CreateDocumento(w http.ResponseWriter, r *http.Request) {
	var d models.EntrenadorDocumento
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err := config.DB.QueryRow(
		`INSERT INTO entrenador_documentos (entrenador_id, identificacion, nombre_archivo, activo)
		VALUES ($1,$2,$3,$4) RETURNING id`,
		d.EntrenadorID, d.Identificacion, d.NombreArchivo, d.Activo,
	).Scan(&d.ID)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, d)
}
