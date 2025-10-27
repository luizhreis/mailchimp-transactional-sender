package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/luizhreis/mailchimp-transactional-sender/pkg/mandrill"
)

func main() {
	fmt.Println("🚀 Mailchimp Transactional Email Sender")
	fmt.Println("======================================")

	// Verificar API key
	apiKey := os.Getenv("MANDRILL_API_KEY")
	if apiKey == "" {
		fmt.Println("❌ Erro: Variável de ambiente MANDRILL_API_KEY não encontrada!")
		fmt.Println()
		fmt.Println("Como configurar:")
		fmt.Println("1. Execute: ./scripts/setup.sh")
		fmt.Println("2. Ou defina manualmente: export MANDRILL_API_KEY=\"sua-api-key\"")
		log.Fatal("Configure a API key e tente novamente.")
	}

	// Criar cliente
	client := mandrill.NewClient(apiKey)

	// ⚠️ IMPORTANTE: Configure seus emails aqui
	destinatario := "ana.borges@icrescer.org.br" // 🔥 ALTERE PARA O EMAIL DESEJADO
	remetente := "luizhreis.gris@gmail.com"         // 🔥 USE SEU EMAIL REAL AQUI
	nomeRemetente := "Ana"                // 🔥 ALTERE PARA SEU NOME

	// Validações básicas
	if strings.Contains(remetente, "@example.com") {
		fmt.Println("🚨 ERRO: Você precisa alterar o email remetente!")
		fmt.Println("❌ @example.com não é um email válido")
		fmt.Println()
		fmt.Println("💡 Use um dos seguintes:")
		fmt.Println("   • Seu email @gmail.com que você possui")
		fmt.Println("   • Seu email @yahoo.com ou @outlook.com")
		fmt.Println("   • Email do seu domínio verificado no Mandrill")
		fmt.Println()
		fmt.Println("📝 Edite este arquivo e use seu email real")
		return
	}

	if strings.Contains(destinatario, "@example.com") {
		fmt.Println("🚨 ERRO: Você precisa alterar o email destinatário!")
		fmt.Println("❌ @example.com não é um email válido")
		fmt.Println()
		fmt.Println("📝 Edite este arquivo e use um email real para teste")
		return
	}

	fmt.Printf("📧 De: %s <%s>\n", nomeRemetente, remetente)
	fmt.Printf("📧 Para: %s\n", destinatario)

	// Verificar potential domain mismatch
	remetenteDomain := strings.Split(remetente, "@")[1]
	destinatarioDomain := strings.Split(destinatario, "@")[1]

	if remetenteDomain != destinatarioDomain {
		fmt.Printf("⚠️ ATENÇÃO: Domínios diferentes detectados!\n")
		fmt.Printf("   Remetente: @%s\n", remetenteDomain)
		fmt.Printf("   Destinatário: @%s\n", destinatarioDomain)
		fmt.Println("💡 Se receber erro 'recipient-domain-mismatch':")
		fmt.Println("   1. Desative 'Domain Matching' no Mandrill")
		fmt.Println("   2. Ou teste com mesmo domínio temporariamente")
		fmt.Println()
	}

	// Dica sobre email remetente
	if strings.Contains(remetente, "@gmail.com") || strings.Contains(remetente, "@yahoo.com") || strings.Contains(remetente, "@outlook.com") {
		fmt.Println("✅ Email remetente parece ser de provedor confiável")
	} else {
		fmt.Println("✅ Usando domínio próprio - certifique-se que está verificado no Mandrill")
	}
	fmt.Println()

	// Carregar anexo (opcional)
	var attachment *mandrill.Attachment
	attachmentPath := "exemplo.txt"
	if _, err := os.Stat(attachmentPath); err == nil {
		var err error
		attachment, err = mandrill.LoadAttachment(attachmentPath)
		if err != nil {
			log.Printf("⚠️ Aviso: Não foi possível carregar o anexo: %v", err)
		} else {
			fmt.Printf("📎 Anexo encontrado: %s (%d bytes)\n", attachment.Filename, len(attachment.Content))
		}
	} else {
		fmt.Println("📎 Nenhum anexo encontrado (arquivo exemplo.txt não existe)")
	}

	// Criar solicitação de email
	emailReq := mandrill.EmailRequest{
		To:        destinatario,
		Subject:   "🎯 Email Individual - Mailchimp Transactional",
		FromEmail: remetente,
		FromName:  nomeRemetente,
		HTMLContent: fmt.Sprintf(`
			<div style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px;">
				<div style="background: linear-gradient(135deg, #ff6b35 0%%, #f7931e 100%%); color: white; padding: 30px; border-radius: 10px; text-align: center; margin-bottom: 30px;">
					<h1 style="margin: 0; font-size: 28px;">📧 Email Individual</h1>
					<p style="margin: 10px 0 0 0; opacity: 0.9;">Mailchimp Transactional (Mandrill)</p>
				</div>

				<div style="background: #f8f9fa; padding: 20px; border-radius: 8px; margin-bottom: 20px;">
					<h2 style="color: #495057; margin-top: 0;">Olá! 👋</h2>
					<p>Este email foi enviado <strong>individualmente</strong> para:</p>
					<div style="background: white; padding: 15px; border-radius: 5px; border-left: 4px solid #ff6b35;">
						<strong>📧 %s</strong>
					</div>
				</div>

				<div style="background: white; border: 1px solid #dee2e6; border-radius: 8px; overflow: hidden; margin-bottom: 20px;">
					<div style="background: #e9ecef; padding: 15px; border-bottom: 1px solid #dee2e6;">
						<h3 style="margin: 0; color: #495057;">🚀 Características do Mandrill:</h3>
					</div>
					<div style="padding: 20px;">
						<ul style="margin: 0; padding-left: 20px;">
							<li style="margin-bottom: 8px;">✅ Email transacional individual</li>
							<li style="margin-bottom: 8px;">✅ Anexos funcionam perfeitamente</li>
							<li style="margin-bottom: 8px;">✅ Produto específico do Mailchimp</li>
							<li style="margin-bottom: 8px;">✅ API robusta e confiável</li>
							<li style="margin-bottom: 8px;">✅ Não requer listas de contatos</li>
						</ul>
					</div>
				</div>

				<div style="background: #d4edda; border: 1px solid #c3e6cb; color: #155724; padding: 15px; border-radius: 8px; margin-bottom: 20px;">
					<strong>💡 Diferencial:</strong> Este email foi enviado usando Mailchimp Transactional, projetado especificamente para emails individuais!
				</div>

				<hr style="border: none; height: 1px; background: #dee2e6; margin: 30px 0;">
				
				<div style="text-align: center; color: #6c757d; font-size: 14px;">
					<p>Enviado via Mailchimp Transactional API por <strong>%s</strong></p>
					<p>%s</p>
				</div>
			</div>
		`, destinatario, nomeRemetente, "24 de outubro de 2025"),
		TextContent: fmt.Sprintf(`
Olá!

Este é um email individual enviado especificamente para: %s

Características do Mailchimp Transactional:
- Email transacional individual
- Anexos funcionam perfeitamente  
- Produto específico do Mailchimp
- API robusta e confiável

Enviado via Mailchimp Transactional API por %s
24 de outubro de 2025
		`, destinatario, nomeRemetente),
		Attachment: attachment,
	}

	// Enviar email
	fmt.Println("📤 Enviando email...")
	err := client.SendEmail(emailReq)
	if err != nil {
		log.Fatalf("❌ Erro ao enviar email: %v", err)
	}

	fmt.Println()
	fmt.Println("✅ Email individual enviado com sucesso!")
	fmt.Printf("📧 Destinatário: %s\n", destinatario)
	if attachment != nil {
		fmt.Printf("📎 Anexo incluído: %s\n", attachment.Filename)
	}
	fmt.Println()
	fmt.Println("💡 Dica: Verifique a caixa de entrada (e spam) do destinatário.")
}
