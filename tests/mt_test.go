package tests

import (
	"io/ioutil"
	"testing"

	"taCaro-backend/extractors/mt"
	"taCaro-backend/models"
)

// readTestHTML reads a test HTML file from disk.
func readTestHTML(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func TestExtractInfo_FromHTMLFile(t *testing.T) {
	// Load the mock HTML file
	htmlContent, err := readTestHTML("../sefazMtMock.html") // Ensure this file exists!
	if err != nil {
		t.Fatalf("Failed to read HTML file: %v", err)
	}

	// Create an Extractor instance
	extractor := mt.Extractor{}
	// Extract items from the HTML content
	items := extractor.ExtractInfo(htmlContent)

	// Ensure that at least one item is extracted
	if len(items) == 0 {
		t.Fatal("Expected at least one item to be extracted, but got none")
	}

	// Check the first extracted item
	item := items[0]

	if item.TxtTit == "" {
		t.Errorf("Expected TxtTit to be populated, but got empty")
	}
	if item.RCod == "" {
		t.Errorf("Expected RCod to be populated, but got empty")
	}
	if item.Valor == 0 {
		t.Errorf("Expected Valor to be a non-zero number, but got 0")
	}
	if item.Qtde == 0 {
		t.Errorf("Expected Qtde to be a non-zero number, but got 0")
	}
	if item.Unit == models.UnitUnknown {
		t.Errorf("Expected Unit to be identified, but got UnitUnknown")
	}
}
