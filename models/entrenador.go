package models

import "time"

type Entrenador struct {
	IDEntrenadores      int        `json:"id_entrenadores"`
	IDAdministrador     int        `json:"id_adminitrador"`
	Cedula              string     `json:"cedula"`
	Descripcion         string     `json:"descripcion"`
	Especialidad        string     `json:"especialidad"`
	GimnasioID          int        `json:"gimnasio_id"`
	AniosExperiencia    int        `json:"anios_experiencia"`
	FotoURL             string     `json:"foto_url"`
	AprovacionEntrenador bool      `json:"aprovacion_entrenador"`
	HojaVida            string     `json:"hoja_vida"`
	CalificacionProm    float64    `json:"calificacion_prom"`
	Activo              bool       `json:"activo"`
	FechaModificacion   *time.Time `json:"fecha_modificacion"`
	FechaCreacion       *time.Time `json:"fecha_creacion"`
}
