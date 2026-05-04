package models

import "time"

type SolicitudEntrenador struct {
	ID               int        `json:"id"`
	UsuarioID        int        `json:"usuario_id"`
	GimnasioID       int        `json:"gimnasio_id"`
	NombresApellidos string     `json:"nombres_apellidos"`
	Cedula           string     `json:"cedula"`
	Experiencia      string     `json:"experiencia"`
	SobreMi          string     `json:"sobre_mi"`
	Especialidad     string     `json:"especialidad"`
	Whatsapp         string     `json:"whatsapp"`
	Correo           string     `json:"correo"`
	Direccion        string     `json:"direccion"`
	Estado           string     `json:"estado"`
	RevisadoPor      bool       `json:"revisado_por"`
	RevisadoEn       *time.Time `json:"revisado_en"`
	CreadoEn         *time.Time `json:"creado_en"`
	Activo           bool       `json:"activo"`
	FechaModificacion *time.Time `json:"fecha_modificacion"`
	FechaCreacion    *time.Time `json:"fecha_creacion"`
}
