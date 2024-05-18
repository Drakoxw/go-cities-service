package models

type City struct {
	ID           int    `json:"id,omitempty"`
	Nombre       string `json:"nombre"`
	CodigoDANE   string `json:"codigodane"`
	Departamento string `json:"departamento"`
}
