package main

import (
	"fmt"
	"log"
	"net/http"
	database "taCaro-backend/databases" // novo pacote criado
	_ "taCaro-backend/docs"             // Documentação do swag
	"taCaro-backend/endpoints"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	if err := database.Connect("mongodb://pemar:123456@172.18.0.2:27017/?authSource=admin"); err != nil {
		log.Fatal("Não foi possível estabelecer conexão com o MongoDB:", err)
	}

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

	fmt.Println("Servidor rodando na porta 9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
