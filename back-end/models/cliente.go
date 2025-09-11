package models

type Cliente struct {
	Telefone        string `db:"telefonecliente" json:"telefoneCliente"`
	NomeCliente     string `db:"nomecliente" json:"nomeCliente"`
	DataNascimento  string `db:"datanascimentocliente" json:"datanascimentoCliente"`
	Genero          string `db:"generocliente" json:"generoCliente"`
	NivelTecnico    string `db:"niveltecnicocliente" json:"nivelTecnicoCliente"`
	Interesse       string `db:"interessecliente" json:"interesseCliente"`
	Disponibilidade string `db:"disponibilidadecliente" json:"disponibilidadeCliente"`
}
