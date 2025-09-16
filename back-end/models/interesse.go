package models

type Interesse struct {
	IdInteresse                int    `db:"idinteresse" json:"idInteresse"`
	InteresseAtual             string `db:"interesseatual" json:"fluxo_do_cliente"`
	CidadeInteresse            string `db:"cidadeinteresse" json:"cidade_de_interesse"`
	IntervaloPreco             string `db:"intervalopreco" json:"preco_do_cliente"`
	Observacao                 string `db:"observacao" json:"observacao"`
	TipoImovelInteresse        string `db:"tipoimovelinteresse" json:"tipoImovelInteresse"`
	Fk_Cliente_telefone        string `db:"fk_cliente_telefone" json:"telefone_do_cliente"`
	Fk_Imovel_idImovel         int    `db:"fk_imovel_idimovel" json:"fk_imovel_idimovel"`
	Fk_Lancamento_idLancamento int    `db:"fk_lancamento_idlancamento" json:"fk_lancamento_idlancamento"`
}
