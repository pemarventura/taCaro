package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"taCaro-backend/models"

	selector "taCaro-backend/extractors"
)

// ProcessarQRCode receives a URL, performs a GET request, and extracts information from the HTML.
//
// @Summary      Log and extract info from URL
// @Description  Receives a URL!, performs a GET request and extracts HTML info.
// @Tags         api-v1
// @Accept       json
// @Produce      json
// @Param        request  body      models.RequestData  true  "Request payload"
// @Success      204      {string}  string       "No Content"
// @Failure      400      {string}  string       "Bad Request"
// @Failure      405      {string}  string       "Method Not Allowed"
// @Failure      500      {string}  string       "Internal Server Error"
// @Router       /processarQRCode [post]
func ProcessarQRCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Apenas POST é permitido", http.StatusMethodNotAllowed)
		return
	}

	var data models.RequestData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("URL recebida: %s\n", data.URL)

	resp, err := http.Get(data.URL)
	if err != nil {
		fmt.Printf("Erro ao requisitar a URL: %v\n", err)
		http.Error(w, "Erro ao requisitar a URL", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erro ao ler a resposta: %v\n", err)
		http.Error(w, "Erro ao ler a resposta", http.StatusInternalServerError)
		return
	}

	extractorInstance := selector.SelectExtractor(data.URL)
	arr := extractorInstance.ExtractInfo(string(body))

	if len(arr) > 0 {
		for _, item := range arr {
			fmt.Printf("TxtTit: %s, RCod: %s, Qtde: %.2f, Valor: %.2f, UN: %s \n",
				item.TxtTit, item.RCod, item.Qtde, item.Valor, item.Unit)
			w.WriteHeader(http.StatusNoContent)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
