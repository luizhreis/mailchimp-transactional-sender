// Exemplo de email simples
// Para executar: go run examples/simple-email/main.go
package main

import (
	"log"
	"os"

	"github.com/luizhreis/mailchimp-transactional-sender/pkg/mandrill"
)

func main() {
	// Configurar cliente
	client := mandrill.NewClient(os.Getenv("MANDRILL_API_KEY"))

	// Email simples
	email := mandrill.EmailRequest{
		To:          "destinatario@example.com", // Altere aqui
		Subject:     "Email Simples de Teste",
		FromEmail:   "seu-email@gmail.com", // Altere aqui
		FromName:    "Seu Nome",            // Altere aqui
		HTMLContent: "<h1>Olá!</h1><p>Este é um email de teste simples.</p>",
		TextContent: "Olá! Este é um email de teste simples.",
	}

	// Enviar
	err := client.SendEmail(email)
	if err != nil {
		log.Fatalf("Erro ao enviar email: %v", err)
	}

	log.Println("✅ Email simples enviado com sucesso!")
}
