package tests

import (
	"testing"

	"taCaro-backend/models" // Import from the correct path
)

func TestUnitString(t *testing.T) {
	tests := []struct {
		unit     models.Unit
		expected string
	}{
		{models.UnitKG, "KG"},
		{models.UnitLiter, "Liter"},
		{models.UnitUnity, "UN"},
		{models.UnitUnknown, "Unknown"},
	}

	for _, test := range tests {
		result := test.unit.String()
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}
