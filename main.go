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
	"os"

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

	// Lê a variável de ambiente PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor padrão se não definido
	}

	fmt.Printf("Servidor rodando na porta %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
