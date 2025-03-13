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

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Redireciona a raiz para o Swagger UI
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusSeeOther)
	})

	// Rota para /processarQRCode (chama a função exportada no package endpoints)
	http.HandleFunc("/api-v1/processarQRCode", endpoints.ProcessarQRCode)

	// Rota para o Swagger UI
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
