package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestLogHandler envia uma requisição POST com uma URL fixa e verifica a resposta do handler.
func TestLogHandler(t *testing.T) {
	// URL fixa para teste
	fixedURL := "https://www.sefaz.mt.gov.br/nfce/consultanfce?p=51250373909400000350650110003660441894058178|2|1|1|7da4581730a059c3a5f2b22ef6ff21e192722257"
	// Cria o payload JSON com a URL fixa
	payload := []byte(`{"url":"` + fixedURL + `"}`)

	// Cria uma nova requisição POST para a rota /log
	req, err := http.NewRequest("POST", "/log", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Erro ao criar a requisição: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Cria um ResponseRecorder para capturar a resposta
	rr := httptest.NewRecorder()

	// Executa o handler logHandler com a requisição e o ResponseRecorder
	handler := http.HandlerFunc(logHandler)
	handler.ServeHTTP(rr, req)

	// Verifica se o status da resposta é 200 OK
	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("Código de status retornado incorreto: obtido %v, esperado %v", status, http.StatusOK)
	// }

	// Verifica se o corpo da resposta é o esperado
	// expected := "URL logada com sucesso"
	// if rr.Body.String() != expected {
	// 	t.Errorf("Resposta inesperada: obtido '%v', esperado '%v'", rr.Body.String(), expected)
	// }
}
