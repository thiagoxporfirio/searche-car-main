package models

type Car struct {
	MarcaEModelo string `json:"marcaEModelo,omitempty"  bson:"marcaEModelo,omitempty"`
	Nome         string `json:"nome,omitempty"  bson:"nome,omitempty"`
	Cor          string `json:"cor,omitempty"  bson:"cor,omitempty"`
	Municipio    string `json:"municipio,omitempty"  bson:"municipio,omitempty"`
	AnoDoCarro   string `json:"anoDoCarro,omitempty"  bson:"anoDoCarro,omitempty"`
	Placa        string `json:"placa,omitempty"  bson:"placa,omitempty"`
	State        string `json:"state,omitempty"  bson:"state"`
	UserId       string `json:"userId,omitempty" bson:"userId"`
	Renavam      string `json:"renavam,omitempty" bson:"renavam"`
	Chassi       string `json:"chassi,omitempty" bson:"chassi"`
}
