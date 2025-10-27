// Exemplo de email com anexo
// Para executar: go run examples/with-attachment/main.go
package main

import (
	"log"
	"os"

	"github.com/luizhreis/mailchimp-transactional-sender/pkg/mandrill"
)

func main() {
	// Configurar cliente
	client := mandrill.NewClient(os.Getenv("MANDRILL_API_KEY"))

	// Carregar anexo
	attachment, err := mandrill.LoadAttachment("exemplo.txt")
	if err != nil {
		log.Fatalf("Erro ao carregar anexo: %v", err)
	}

	// Email com anexo
	email := mandrill.EmailRequest{
		To:          "destinatario@example.com", // Altere aqui
		Subject:     "Email com Anexo",
		FromEmail:   "seu-email@gmail.com", // Altere aqui
		FromName:    "Seu Nome",            // Altere aqui
		HTMLContent: "<h1>Documento em Anexo</h1><p>Segue o documento solicitado em anexo.</p>",
		TextContent: "Documento em Anexo\n\nSegue o documento solicitado em anexo.",
		Attachment:  attachment,
	}

	// Enviar
	err = client.SendEmail(email)
	if err != nil {
		log.Fatalf("Erro ao enviar email: %v", err)
	}

	log.Printf("✅ Email com anexo enviado com sucesso! Anexo: %s", attachment.Filename)
}
