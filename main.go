package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	selector "taCaro-backend/extractors"
)

// RequestData representa a estrutura para receber os dados enviados pelo app.
type RequestData struct {
	URL string `json:"url"`
}

// logHandler recebe uma URL, faz uma requisição GET para ela e extrai informações do HTML.
func logHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Apenas POST é permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodifica o JSON recebido
	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Loga a URL recebida
	fmt.Printf("URL recebida: %s\n", data.URL)

	// Realiza a requisição GET para a URL recebida
	resp, err := http.Get(data.URL)
	if err != nil {
		fmt.Printf("Erro ao requisitar a URL: %v\n", err)
		http.Error(w, "Erro ao requisitar a URL", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erro ao ler a resposta: %v\n", err)
		http.Error(w, "Erro ao ler a resposta", http.StatusInternalServerError)
		return
	}

	// Loga o corpo da resposta (HTML)
	// fmt.Printf("Resposta da requisição:\n%s\n", string(body))

	// Seleciona o extractor com base na URL
	extractorInstance := selector.SelectExtractor(data.URL)

	// Utiliza a função de extração do package extractor para processar o HTML.
	arr := extractorInstance.ExtractInfo(string(body))
	for _, item := range arr {
		fmt.Printf("TxtTit: %s, RCod: %s, Qtde: %.2f, Valor: %.2f, UN: %s \n", item.TxtTit, item.RCod, item.Qtde, item.Valor, item.Unit)
	}

	// Retorna um status 204 No Content, pois nenhuma resposta ao cliente é necessária.
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/log", logHandler)
	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
