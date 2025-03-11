package extractor

import (
	"strings"

	"golang.org/x/net/html"
)

type BaseExtractor struct{}

// elementChildren retorna todos os filhos de 'n' que são ElementNode.
func (b *BaseExtractor) ElementChildren(n *html.Node) []*html.Node {
	var children []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			children = append(children, c)
		}
	}
	return children
}

// getAttr retorna o valor do atributo especificado em 'n', se existir.
func (b *BaseExtractor) GetAttr(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

// GetText retorna todo o texto contido em 'n' e seus descendentes.
func (b *BaseExtractor) GetText(n *html.Node) string {
	if n == nil {
		return ""
	}
	var sb strings.Builder
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			sb.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return strings.TrimSpace(sb.String())
}
