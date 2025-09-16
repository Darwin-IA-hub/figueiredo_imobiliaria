package models

type Lancamento struct {
	IdLancamento     int    `db:"idlancamento" json:"idLancamento"`
	CidadeLancamento string `db:"cidadelancamento" json:"cidadeLancamento"`
	NomeLancamento   string `db:"nomelancamento" json:"nomeLancamento"`
	Detalhes         string `db:"detalhes" json:"detalhes"`
}
