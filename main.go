// @title taCaro Backend API
// @version 1.0
// @description API que processa notas fiscais e extrai informações.
// @host localhost:8080
// @BasePath /
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// "strings"

	selector "taCaro-backend/extractors"

	// Swagger docs
	_ "taCaro-backend/docs"

	// Swagger UI handler
	httpSwagger "github.com/swaggo/http-swagger"
)

// RequestData represents the payload for the log endpoint.
// swagger:model
type RequestData struct {
	// URL to be fetched and processed.
	URL string `json:"url"`
}

// processarQRCode receives a URL, performs a GET request to it, and extracts information from the HTML.
// @Summary      Log and extract info from URL
// @Description  Receives a URL, performs a GET request and extracts HTML info.
// @Tags         log
// @Accept       json
// @Produce      json
// @Param        request  body      RequestData  true  "Request payload"
// @Success      204      {string}  string       "No Content"
// @Failure      400      {string}  string       "Bad Request"
// @Failure      405      {string}  string       "Method Not Allowed"
// @Failure      500      {string}  string       "Internal Server Error"
// @Router       /processarQRCode [post]
func processarQRCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Apenas POST é permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decode the incoming JSON payload
	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Log the received URL
	fmt.Printf("URL recebida: %s\n", data.URL)

	// Perform the GET request to the received URL
	resp, err := http.Get(data.URL)
	if err != nil {
		fmt.Printf("Erro ao requisitar a URL: %v\n", err)
		http.Error(w, "Erro ao requisitar a URL", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erro ao ler a resposta: %v\n", err)
		http.Error(w, "Erro ao ler a resposta", http.StatusInternalServerError)
		return
	}

	// Select the extractor based on the URL
	extractorInstance := selector.SelectExtractor(data.URL)

	// Extract information from the HTML
	arr := extractorInstance.ExtractInfo(string(body))
	if len(arr) > 0 {
		for _, item := range arr {
			fmt.Printf("TxtTit: %s, RCod: %s, Qtde: %.2f, Valor: %.2f, UN: %s \n",
				item.TxtTit, item.RCod, item.Qtde, item.Valor, item.Unit)
			// Here you can marshal and write a response if needed.
			w.WriteHeader(http.StatusNoContent)
		}
	} else {
		// If no items are found, return 400 Bad Request.
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	// Garante um caminho base e redirecionamento para o Swagger UI
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusSeeOther)
	})

	// Expose o endpoint /log
	http.HandleFunc("/log", processarQRCode)

	// Serve o Swagger UI em /swagger/index.html
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
