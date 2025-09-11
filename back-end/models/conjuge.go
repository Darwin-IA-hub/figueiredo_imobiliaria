package models

type Conjuge struct{
	IdConjuge int `db:"idconjuge" json:"idConjuge"`
	RendaBrutaMensalConjuge float64 `db:"rendabrutamensalconjuge" json:"renda_bruta_do_cliente_conjuge"`
	DataNascimentoConjuge string `db:"datanascimentoconjuge" json:"data_de_nascimento_conjuge"`
	Fk_Cliente_telefone string `db:"fk_cliente_telefone" json:"fk_Cliente_telefone"`
}
