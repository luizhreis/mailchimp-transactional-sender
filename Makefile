# Makefile para Mailchimp Transactional Sender

.PHONY: help build run test clean install setup

# Variáveis
BINARY_NAME=mandrill-sender
MAIN_PATH=./cmd/main.go

help: ## Mostra esta ajuda
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

setup: ## Configura o ambiente de desenvolvimento
	@echo "🔧 Configurando ambiente..."
	@chmod +x scripts/setup.sh
	@./scripts/setup.sh

build: ## Compila o projeto
	@echo "🔨 Compilando..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "✅ Build concluído: $(BINARY_NAME)"

run: ## Executa a aplicação principal
	@echo "🚀 Executando aplicação..."
	@go run $(MAIN_PATH)

test: ## Executa todos os testes
	@echo "🧪 Executando testes..."
	@go test ./...

test-cover: ## Executa testes com coverage
	@echo "🧪 Executando testes com coverage..."
	@go test -cover ./...

clean: ## Remove arquivos de build
	@echo "🧹 Limpando arquivos de build..."
	@rm -f $(BINARY_NAME)
	@go clean

install: ## Instala dependências
	@echo "📦 Instalando dependências..."
	@go mod tidy
	@go mod download

example-simple: ## Executa exemplo de email simples
	@echo "📧 Executando exemplo de email simples..."
	@go run examples/simple-email/main.go

example-attachment: ## Executa exemplo de email com anexo
	@echo "📎 Executando exemplo de email com anexo..."
	@go run examples/with-attachment/main.go

diagnostic-unsigned: ## Diagnóstico para erro "unsigned"
	@echo "🔍 Executando diagnóstico para erro 'unsigned'..."
	@go run examples/diagnostic-unsigned/main.go

lint: ## Executa linter (requer golangci-lint)
	@echo "🔍 Executando linter..."
	@golangci-lint run

fmt: ## Formata o código
	@echo "✨ Formatando código..."
	@go fmt ./...

mod-update: ## Atualiza dependências
	@echo "📦 Atualizando dependências..."
	@go get -u ./...
	@go mod tidy

docker-build: ## Constrói imagem Docker
	@echo "🐳 Construindo imagem Docker..."
	@docker build -t $(BINARY_NAME) .

git-init: ## Inicializa repositório Git
	@echo "📁 Inicializando repositório Git..."
	@git init
	@git add .
	@git commit -m "Initial commit: Mailchimp Transactional Sender"

git-status: ## Mostra status do Git
	@git status

release: build test ## Prepara release
	@echo "🎉 Preparando release..."
	@echo "✅ Build e testes concluídos com sucesso!"

dev: ## Modo desenvolvimento (watch)
	@echo "🔄 Modo desenvolvimento - pressione Ctrl+C para parar"
	@while true; do \
		go run $(MAIN_PATH); \
		echo "💤 Aguardando 5 segundos..."; \
		sleep 5; \
	done