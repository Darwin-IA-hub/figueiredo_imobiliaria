package models

type Financiamento struct {
	IdFinanciamento        int    `db:"idfinanciamento" json:"idFinanciamento"`
	DescricaoFinanciamento string `db:"descricaofinanciamento" json:"tipoFinanciamento"`
	Fk_Cliente_telefone    string `db:"fk_cliente_telefone" json:"fk_Cliente_telefone"`
}
