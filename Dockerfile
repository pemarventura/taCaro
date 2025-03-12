# Use uma imagem oficial do Go como base
FROM golang:1.24

# Defina o diretório de trabalho
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./

# Run `go mod download` before copying the rest of the files
RUN go mod download
# Copie os arquivos do projeto para o container
COPY . .

# Roda os testes antes de compilar a aplicação
RUN go test ./...

# Compile a aplicação (altere "minha_app" para o nome desejado)
RUN go build -o taCaro-backend .

# Exponha a porta que sua aplicação utilizará (ajuste se necessário)
EXPOSE 8080

# Comando para iniciar a aplicação
CMD ["./taCaro-backend"]
