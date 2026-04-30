package controllers

import (
	"API_ENTRENADORES/config"
	"API_ENTRENADORES/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GET ALL solicitudes (filtros: estado, especialidad)
func GetAllSolicitudes(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, usuario_id, gimnasio_id, nombres_apellidos, cedula, experiencia,
		sobre_mi, especialidad, whatsapp, correo, direccion, estado, revisado_por,
		revisado_en, creado_en, activo, fecha_modificacion, fecha_creacion
		FROM solicitudes_entrenador WHERE 1=1`

	estado       := r.URL.Query().Get("estado")
	especialidad := r.URL.Query().Get("especialidad")

	if estado != "" {
		query += " AND estado ILIKE '%" + estado + "%'"
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

	var lista []models.SolicitudEntrenador
	for rows.Next() {
		var s models.SolicitudEntrenador
		rows.Scan(
			&s.ID, &s.UsuarioID, &s.GimnasioID, &s.NombresApellidos,
			&s.Cedula, &s.Experiencia, &s.SobreMi, &s.Especialidad,
			&s.Whatsapp, &s.Correo, &s.Direccion, &s.Estado,
			&s.RevisadoPor, &s.RevisadoEn, &s.CreadoEn,
			&s.Activo, &s.FechaModificacion, &s.FechaCreacion,
		)
		lista = append(lista, s)
	}
	respondJSON(w, 200, lista)
}


func GetSolicitudByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var s models.SolicitudEntrenador

	err := config.DB.QueryRow(
		`SELECT id, usuario_id, gimnasio_id, nombres_apellidos, cedula, experiencia,
		sobre_mi, especialidad, whatsapp, correo, direccion, estado, revisado_por,
		revisado_en, creado_en, activo, fecha_modificacion, fecha_creacion
		FROM solicitudes_entrenador WHERE id=$1`, id,
	).Scan(
		&s.ID, &s.UsuarioID, &s.GimnasioID, &s.NombresApellidos,
		&s.Cedula, &s.Experiencia, &s.SobreMi, &s.Especialidad,
		&s.Whatsapp, &s.Correo, &s.Direccion, &s.Estado,
		&s.RevisadoPor, &s.RevisadoEn, &s.CreadoEn,
		&s.Activo, &s.FechaModificacion, &s.FechaCreacion,
	)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Solicitud no encontrada"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, s)
}

func CreateSolicitud(w http.ResponseWriter, r *http.Request) {
	var s models.SolicitudEntrenador
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err := config.DB.QueryRow(
		`INSERT INTO solicitudes_entrenador (usuario_id, gimnasio_id, nombres_apellidos,
		cedula, experiencia, sobre_mi, especialidad, whatsapp, correo, direccion,
		estado, revisado_por, activo)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING id`,
		s.UsuarioID, s.GimnasioID, s.NombresApellidos, s.Cedula,
		s.Experiencia, s.SobreMi, s.Especialidad, s.Whatsapp,
		s.Correo, s.Direccion, s.Estado, s.RevisadoPor, s.Activo,
	).Scan(&s.ID)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 201, s)
}