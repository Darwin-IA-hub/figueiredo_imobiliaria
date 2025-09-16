package models

type Interesse struct {
	IdInteresse                int    `db:"idinteresse" json:"idInteresse"`
	InteresseAtual             string `db:"interesseatual" json:"fluxo_do_cliente"`
	CidadeInteresse            string `db:"cidadeinteresse" json:"cidadeInteresse"`
	IntervaloPreco             string `db:"intervalopreco" json:"intervaloPreco"`
	Observacao                 string `db:"observacao" json:"observacao"`
	TipoImovelInteresse        string `db:"tipoimovelinteresse" json:"tipoImovelInteresse"`
	Fk_Cliente_telefone        string `db:"fk_cliente_telefone" json:"fk_Cliente_telefone"`
	Fk_Imovel_idImovel         int    `db:"fk_imovel_idimovel" json:"Fk_Imovel_idImovel"`
	Fk_Lancamento_idLancamento int    `db:"fk_lancamento_idlancamento" json:"fk_Lancamento_idLancamento"`
}
