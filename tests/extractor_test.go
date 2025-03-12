package tests

import (
	"strings"
	extractor "taCaro-backend/extractors/base_extractor"
	"testing"

	"golang.org/x/net/html"
)

func TestElementChildren(t *testing.T) {
	htmlData := `<div><p>Test</p><span>Ignore</span></div>`
	doc, err := html.Parse(strings.NewReader(htmlData))
	if err != nil {
		t.Fatalf("Error parsing HTML: %v", err)
	}

	// Locate the <body> tag first
	var body = FindBody((doc))

	if body == nil {
		t.Fatal("Failed to locate <body> element")
	}

	// Locate the <div> inside <body>
	var div *html.Node
	for c := body.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "div" {
			div = c
			break
		}
	}

	if div == nil {
		t.Fatal("Failed to locate <div> element")
	}

	// Now extract children from the <div>
	extractor := extractor.BaseExtractor{}
	children := extractor.ElementChildren(div)

	// Debug children count
	t.Logf("Children found: %d", len(children))

	if len(children) != 2 {
		t.Errorf("Expected 2 children, got %d", len(children))
	}

	if len(children) >= 2 {
		if children[0].Data != "p" || children[1].Data != "span" {
			t.Errorf("Unexpected children node names: got [%s, %s]", children[0].Data, children[1].Data)
		}
	}
}

func TestGetAttr(t *testing.T) {
	htmlData := `<div id="test" class="example"></div>`
	doc, err := html.Parse(strings.NewReader(htmlData))
	if err != nil {
		t.Fatalf("Error parsing HTML: %v", err)
	}

	extractor := extractor.BaseExtractor{}

	// Get <body> node
	body := FindBody(doc)
	if body == nil {
		t.Fatal("Failed to find <body> element")
	}

	// Ensure <div> is correctly selected
	targetNode := body.FirstChild
	if targetNode == nil || targetNode.Data != "div" {
		t.Fatal("Expected <div> inside <body>, but found none")
	}

	// Test `id` attribute
	attr, exists := extractor.GetAttr(targetNode, "id")
	if !exists || attr != "test" {
		t.Errorf("Expected id='test', got '%s'", attr)
	}

	// Test `class` attribute
	classAttr, classExists := extractor.GetAttr(targetNode, "class")
	if !classExists || classAttr != "example" {
		t.Errorf("Expected class='example', got '%s'", classAttr)
	}

	// Test a non-existent attribute
	nonExistentAttr, nonExistent := extractor.GetAttr(targetNode, "style")
	if nonExistent {
		t.Errorf("Expected 'style' attribute to not exist, but got '%s'", nonExistentAttr)
	}
}

func TestGetText(t *testing.T) {
	htmlData := `<div><p>Hello <b>World</b></p></div>`
	doc, _ := html.Parse(strings.NewReader(htmlData))

	// Get <body> node
	body := FindBody(doc)
	if body == nil {
		t.Fatal("Failed to find <body> element")
	}

	extractor := extractor.BaseExtractor{}
	text := extractor.GetText(body.FirstChild.FirstChild)

	expectedText := "Hello World"
	if text != expectedText {
		t.Errorf("Expected '%s', got '%s'", expectedText, text)
	}
}

func FindBody(n *html.Node) *html.Node {
	if n.Type == html.ElementNode && n.Data == "body" {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if body := FindBody(c); body != nil {
			return body
		}
	}
	return nil
}
