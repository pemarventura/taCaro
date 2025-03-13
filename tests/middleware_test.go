package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// dummyProcessarQRCode simula o endpoint, aceitando apenas POST.
func dummyProcessarQRCode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Processado"))
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()

	// Rota para /api-v1/processarQRCode aceita somente POST.
	router.HandleFunc("/api-v1/processarQRCode", dummyProcessarQRCode).Methods("POST")

	// Define um handler customizado para métodos não permitidos.
	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	return router
}

func TestWrongMethodForCorrectURL(t *testing.T) {
	router := setupRouter()
	// Para a rota /api-v1/processarQRCode, que aceita apenas POST, enviaremos um GET.
	req, err := http.NewRequest("GET", "/api-v1/processarQRCode", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Verifica se o status code é 405.
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Esperado status code 405, mas obteve %d", rr.Code)
	}

	// Verifica se a mensagem de erro contém "Method Not Allowed".
	if !strings.Contains(rr.Body.String(), "Method Not Allowed") {
		t.Errorf("Mensagem de erro esperada 'Method Not Allowed', mas obteve: %s", rr.Body.String())
	}
}
