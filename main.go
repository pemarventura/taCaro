// @title taCaro Backend API
// @version 1.0
// @description API que processa notas fiscais e extrai informações.
// @host localhost:8080
// @BasePath /api-v1
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "taCaro-backend/docs" // Importa a documentação gerada pelo swag
	"taCaro-backend/endpoints"

	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Cria um novo router com Gorilla Mux
	router := mux.NewRouter()

	// Rota raiz: redireciona para o Swagger UI
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusSeeOther)
	}).Methods("GET")

	// Rota para /api-v1/processarQRCode (chama a função exportada no package endpoints)
	router.HandleFunc("/api-v1/processarQRCode", endpoints.ProcessarQRCode)

	// Rota para servir o Swagger UI (todos os caminhos que iniciem com /swagger/ serão atendidos)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Handler customizado para endpoints não encontrados
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Endereço não encontrado", http.StatusNotFound)
	})

	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
