package models

type Cliente struct {
	Telefone              string  `db:"telefone" json:"telefone_do_cliente"`
	NomeCliente           string  `db:"nomecliente" json:"nome_do_cliente"`
	DataNascimentoCliente string  `db:"datanascimentocliente" json:"data_de_nascimento"`
	RendaBrutaCliente     float64 `db:"rendabrutacliente" json:"renda_bruta_do_cliente"`
	QuantidadeFilhos      int     `db:"quantidadefilhos" json:"filhos_quantidade"`
	AnosCarteiraAssinada  int     `db:"anoscarteiraassinada" json:"quantos_anos_de_carteira"`
	TeveSubsidio          string  `db:"tevesubsidio" json:"tem_outros_subsidios"`
	VaiUsarFGTS           string  `db:"vaiusarfgts" json:"fgts"`
	PossuiFinanciamento   string  `db:"possuifinanciamento" json:"possui_financiamento"`
}
