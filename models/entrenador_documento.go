package models

import "time"

type EntrenadorDocumento struct {
	ID                int        `json:"id"`
	EntrenadorID      int        `json:"entrenador_id"`
	Identificacion    string     `json:"identificacion"`
	NombreArchivo     string     `json:"nombre_archivo"`
	CreadoEn          *time.Time `json:"creado_en"`
	Activo            bool       `json:"activo"`
	FechaModificacion *time.Time `json:"fecha_modificacion"`
	FechaCreacion     *time.Time `json:"fecha_creacion"`
}
