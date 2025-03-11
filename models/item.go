package models

import (
	"strconv"
	"strings"
)

// Item representa um objeto com as propriedades extraídas do HTML.
type Item struct {
	TxtTit string  `json:"txtTit"`
	RCod   string  `json:"rCod"`
	Valor  float64 `json:"valor"`
	Qtde   float64 `json:"qtde"`
	Unit   Unit    `json:"unidade"`
}

// NewItem constrói um Item processando os valores:
// - RCod: remove todos os espaços.
// - Valor: substitui vírgula por ponto e converte para float64.
func NewItem(txtTit, rCod, valor string, qtde string, unit string) Item {
	rCod = strings.ReplaceAll(rCod, " ", "")
	valor = strings.ReplaceAll(valor, ",", ".")
	f, err := strconv.ParseFloat(valor, 64)
	if err != nil {
		f = 0.0 // Você pode tratar o erro conforme necessário
	}
	// parsedQtd := strings.ReplaceAll(valor, ",", ".")
	c, err := strconv.ParseFloat(qtde, 64)
	if err != nil {

		c = 0.0 // Default to 0.0 if parsing fails.
	}

	// Convert the unit string to the corresponding Unit enum.
	var u Unit
	switch strings.ToLower(strings.TrimSpace(unit)) {
	case "kg":
		u = UnitKG
	case "liter":
		u = UnitLiter
	case "un":
		u = UnitUnity
	default:
		u = UnitUnknown
	}

	return Item{
		TxtTit: txtTit,
		RCod:   rCod,
		Valor:  f,
		Qtde:   c,
		Unit:   u,
	}
}
