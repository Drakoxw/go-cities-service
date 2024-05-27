package models

type City struct {
	Nombre       string `json:"nombre"`
	CodigoDANE   string `json:"codigodane"`
	Departamento string `json:"departamento"`
}

var CitiesList []City
