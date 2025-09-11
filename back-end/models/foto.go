package models

type Foto struct {
	IdFoto              int    `db:"idfoto" json:"idFoto"`
	LinkFoto            string `db:"linkfoto" json:"linkFoto"`
	Fk_Cliente_telefone string `db:"fk_cliente_telefone" json:"fk_Cliente_telefone"`
}
