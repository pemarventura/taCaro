package tests

import (
	"testing"

	"taCaro-backend/models" // Import from the correct path
)

func TestNewItem(t *testing.T) {
	item := models.NewItem("Test Title", " 123 456 ", "10,99", "2", "kg")

	if item.TxtTit != "Test Title" {
		t.Errorf("Expected TxtTit to be 'Test Title', got '%s'", item.TxtTit)
	}
	if item.RCod != "123456" {
		t.Errorf("Expected RCod to be '123456', got '%s'", item.RCod)
	}
	if item.Valor != 10.99 {
		t.Errorf("Expected Valor to be 10.99, got '%f'", item.Valor)
	}
	if item.Qtde != 2 {
		t.Errorf("Expected Qtde to be 2, got '%f'", item.Qtde)
	}
	if item.Unit != models.UnitKG {
		t.Errorf("Expected Unit to be UnitKG, got '%v'", item.Unit)
	}
}
