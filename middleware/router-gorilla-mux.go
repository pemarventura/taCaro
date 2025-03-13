package middleware

import (
	// "fmt"
	"net/http"
)

// // Handler para a rota inicial
// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Bem-vindo à página inicial!")
// }

// // Handler para um endpoint de exemplo
// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Olá, mundo!")
// }

// Handler para endereços não encontrados
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Endereço não encontrado", http.StatusNotFound)
}
