// models/unit.go
package models

// Unit representa o tipo de unidade de um item.
type Unit int

const (
	// UnitUnknown indica que a unidade não foi definida.
	UnitUnknown Unit = iota
	// UnitKG representa a unidade "KG".
	UnitKG
	// UnitLiter representa a unidade "Liter".
	UnitLiter
	// UnitUnity representa a unidade "Unity".
	UnitUnity
)

// String retorna a representação em string de Unit.
func (u Unit) String() string {
	switch u {
	case UnitKG:
		return "KG"
	case UnitLiter:
		return "Liter"
	case UnitUnity:
		return "UN"
	default:
		return "Unknown"
	}
}
