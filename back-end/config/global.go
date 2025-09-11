package config

import (
	"strings"
	"time"
)

var planos = map[string]int{"bronze": 200, "prata": 1200, "ouro": 10000}

var Mensagem string
var Respondendo bool
var Counter int


const (
	baseUrl     = "http://147.93.10.167:8012"
	plano_atual = "bronze"
)

func GetPlanoAtual() int {
	return planos[plano_atual]
}

func PadronizaTelefone(telefone string) string {
	if len(telefone) > 13 && !strings.HasSuffix(telefone, "@lid") {
		telefone += "@lid"
	} else if !strings.HasSuffix(telefone, "@s.whatsapp.net") && !strings.HasSuffix(telefone, "@lid") {
		telefone += "@s.whatsapp.net"
	}
	return telefone
}

func DentroHorario() bool {
	now := time.Now()
	horaAtual := now.Hour()
	diaSemana := now.Weekday()
	// Horário comercial: das 8h às 18h
	//TODO refatorar baseado no documento
	if (horaAtual >= 8 && horaAtual < 18) {
		if(diaSemana != 0 && diaSemana != 7){
			return true 
		}
	}
	return false
}