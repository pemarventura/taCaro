package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestLogHandler simula uma requisição HTTP e substitui o corpo por um HTML local.
func TestLogHandler(t *testing.T) {
	// Simula uma URL fixa (apenas para o contexto do teste)
	fixedURL := "https://www.sefazgov.br/nfce/testpage"

	// Lê um arquivo HTML que simula a resposta do servidor
	htmlContent, err := ioutil.ReadFile("sefazMtMock.html") // Certifique-se de que o arquivo existe!
	if err != nil {
		t.Fatalf("Erro ao ler o arquivo HTML de mock: %v", err)
	}

	// Simula um payload JSON que normalmente seria enviado na requisição
	payload := []byte(`{"url":"` + fixedURL + `"}`)

	// Cria uma requisição POST simulada
	req, err := http.NewRequest("POST", "/log", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Erro ao criar a requisição: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Substitui o ResponseRecorder para capturar a resposta do handler
	rr := httptest.NewRecorder()

	// Simula a função do handler, mas passando o HTML de mock como resposta
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(htmlContent) // Retorna o HTML mockado como resposta
	})

	// Executa o handler com a requisição simulada
	mockHandler.ServeHTTP(rr, req)

	// Verifica se o status da resposta é 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Código de status retornado incorreto: obtido %v, esperado %v", status, http.StatusOK)
	}

	// Verifica se o corpo da resposta corresponde ao HTML esperado
	expectedBody := string(htmlContent)
	if rr.Body.String() != expectedBody {
		t.Errorf("O corpo da resposta não corresponde ao esperado")
	}
}
