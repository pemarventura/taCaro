package mt

import (
	"fmt"
	"log"
	"strings"

	extractor "taCaro-backend/extractors/base_extractor" // importando o pacote que contém o BaseExtractor
	"taCaro-backend/models"

	"golang.org/x/net/html"
)

// Extractor é o extractor específico para "mt" e incorpora BaseExtractor.
type Extractor struct {
	extractor.BaseExtractor
}

// ExtractInfo extrai informações do HTML utilizando os métodos do BaseExtractor.
func (e *Extractor) ExtractInfo(htmlContent string) []models.Item {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatalf("Erro ao parsear HTML: %v", err)
	}

	// Procura pelo contêiner: busca uma <table> com id "tabResult"
	var container *html.Node
	var findContainer func(*html.Node)
	findContainer = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {
			if id, ok := e.GetAttr(n, "id"); ok && id == "tabResult" {
				fmt.Println("Achou tabResult")
				container = n
				return
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findContainer(c)
			if container != nil {
				return
			}
		}
	}
	findContainer(doc)
	if container == nil {
		log.Println("Nenhum contêiner (<tbody> ou <table> com id 'tabResult') encontrado.")
		return nil
	}

	var items []models.Item

	// Itera pelos filhos do container procurando por <tbody>
	for _, tbody := range e.ElementChildren(container) {
		if tbody.Data != "tbody" {
			continue
		}
		// Itera pelos elementos <tr> dentro do <tbody>
		for _, tr := range e.ElementChildren(tbody) {
			if tr.Data != "tr" {
				continue
			}
			var txtTit, rCod, valor, qtde, unit string

			// Itera por cada <td> do <tr>
			for _, td := range e.ElementChildren(tr) {
				if td.Data != "td" {
					continue
				}
				// Itera sobre os spans de cada <td>
				for _, span := range e.ElementChildren(td) {
					if span.Data != "span" {
						continue
					}
					if class, ok := e.GetAttr(span, "class"); ok {
						if strings.Contains(class, "txtTit") {
							txtTit = e.GetText(span)
						} else if strings.Contains(class, "RCod") {
							rCod = e.GetText(span)
						} else if strings.Contains(class, "RvlUnit") {
							// Procura entre os filhos deste span o span com a classe "valor"
							for _, sp := range e.ElementChildren(span) {
								if sp.Data == "span" {
									if childClass, ok := e.GetAttr(sp, "class"); ok {
										if strings.Contains(childClass, "valor") {
											valor = e.GetText(sp)
											break
										}
									}
								}
							}
						} else if strings.Contains(class, "Rqtd") {
							qtyText := e.GetText(span) // Exemplo: "Qtde.: 1"
							parts := strings.SplitN(qtyText, ":", 2)
							if len(parts) == 2 {
								qtde = strings.TrimSpace(parts[1])
							}
						} else if strings.Contains(class, "valor") {
							// Caso o span "valor" esteja diretamente no td
							valor = e.GetText(span)
						} else if strings.Contains(class, "RUN") {
							// Processa o span que contém o unit information.
							uText := e.GetText(span) // Ex.: "UN: Kg"
							uText = strings.TrimPrefix(uText, "UN:")
							unit = strings.TrimSpace(uText)
							fmt.Println("Unidade extraída:", unit)
						}
					}
				}
			}

			// Se algum campo foi preenchido, cria o item e adiciona ao slice.
			if txtTit != "" || rCod != "" || valor != "" || qtde != "" || unit != "" {
				item := models.NewItem(txtTit, rCod, valor, qtde, unit)
				items = append(items, item)
			}
		}
	}

	return items
}
