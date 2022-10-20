package models

type Descricao struct {
	Placa     string `json:"Placa,omitempty"  bson:"Placa,omitempty"`
	Marca     string `json:"Marca,omitempty"  bson:"Marca,omitempty"`
	Estado    string `json:"Estado,omitempty"  bson:"Estado"`
	Municipio string `json:"Municipio,omitempty"  bson:"Municipio,omitempty"`
}
