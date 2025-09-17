package models

type Imovel struct {
	IdImovel     int    `db:"idimovel" json:"idImovel"`
	TipoImovel   string `db:"tipoimovel" json:"tipo_do_imovel"`
	CidadeImovel string `db:"cidadeimovel" json:"cidade_de_interesse"`
	LinkIPTU     string `db:"linkiptu" json:"linkIPTU"`
}

type ImovelVenda struct {
	Fk_Imovel_idImovel int    `db:"fk_imovel_idimovel" json:"fk_imovel_idimovel"`
	FinanciadoQuitado  string `db:"financiadoquitado" json:"financiado_quitado"`
	DocEmDia           string `db:"docemdia" json:"documentacao_dia"`
	EstaHabitado       string `db:"estahabitado" json:"local_habitado"`
}
