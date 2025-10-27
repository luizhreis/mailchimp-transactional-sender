package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/luizhreis/mailchimp-transactional-sender/pkg/mandrill"
)

func main() {
	fmt.Println("🔍 DIAGNÓSTICO: Erro 'unsigned' do Mandrill")
	fmt.Println("=" + strings.Repeat("=", 50))

	// Verificar API Key
	apiKey := os.Getenv("MANDRILL_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ MANDRILL_API_KEY não configurada")
	}
	fmt.Printf("✅ API Key configurada: %s...\n", apiKey[:10])

	// Configurar cliente
	client := mandrill.NewClient(apiKey)

	// Email de teste com domínio comum
	email := mandrill.EmailRequest{
		To:          "teste@gmail.com", // Use um email real aqui
		Subject:     "Teste Diagnóstico - Unsigned",
		FromEmail:   "seu-email@gmail.com", // MUDE PARA SEU EMAIL REAL
		FromName:    "Teste Mandrill",
		HTMLContent: "<h1>Teste</h1><p>Este é um teste para diagnosticar erro 'unsigned'.</p>",
		TextContent: "Teste para diagnosticar erro 'unsigned'.",
	}

	fmt.Println("\n📧 CONFIGURAÇÃO DO EMAIL:")
	fmt.Printf("   De: %s (%s)\n", email.FromEmail, email.FromName)
	fmt.Printf("   Para: %s\n", email.To)
	fmt.Printf("   Assunto: %s\n", email.Subject)

	fmt.Println("\n🚨 VERIFICAÇÕES IMPORTANTES:")

	// Verificar se está usando email de exemplo
	if strings.Contains(email.FromEmail, "seu-email@") || strings.Contains(email.FromEmail, "@example.com") {
		fmt.Println("❌ PROBLEMA: Você está usando um email de exemplo!")
		fmt.Println("   Mude 'seu-email@gmail.com' para seu email real.")
		fmt.Println("   Exemplo: 'joao.silva@gmail.com'")
		return
	}

	if strings.Contains(email.To, "teste@") || strings.Contains(email.To, "@example.com") {
		fmt.Println("❌ PROBLEMA: Você está usando um email de destino de exemplo!")
		fmt.Println("   Mude 'teste@gmail.com' para um email real que você possa verificar.")
		return
	}

	fmt.Println("✅ Emails parecem válidos")

	fmt.Println("\n🔄 ENVIANDO EMAIL...")

	// Tentar enviar
	err := client.SendEmail(email)
	if err != nil {
		fmt.Printf("\n❌ ERRO ENCONTRADO:\n%v\n", err)

		if strings.Contains(err.Error(), "unsigned") {
			fmt.Println("\n🎯 DIAGNÓSTICO ESPECÍFICO PARA 'UNSIGNED':")
			fmt.Println("=" + strings.Repeat("=", 50))
			fmt.Println("Este erro significa que o Mandrill não reconhece seu domínio de envio.")
			fmt.Println("")
			fmt.Println("📋 SOLUÇÕES EM ORDEM DE PRIORIDADE:")
			fmt.Println("")
			fmt.Println("1️⃣ SOLUÇÃO IMEDIATA - Use Gmail/Yahoo:")
			fmt.Printf("   - Mude FromEmail para um email @gmail.com/@yahoo.com que você possui\n")
			fmt.Printf("   - Exemplo: se você tem joao@gmail.com, use este email\n")
			fmt.Println("")
			fmt.Println("2️⃣ VERIFICAR DOMÍNIO NO MANDRILL:")
			fmt.Println("   - Acesse: https://mandrillapp.com/settings/sending-domains")
			fmt.Println("   - Clique em 'Add Domain'")
			fmt.Printf("   - Adicione o domínio: %s\n", getDomainFromEmail(email.FromEmail))
			fmt.Println("   - Siga as instruções de verificação DNS")
			fmt.Println("")
			fmt.Println("3️⃣ VERIFICAR CONFIGURAÇÕES:")
			fmt.Println("   - Acesse: https://mandrillapp.com/settings")
			fmt.Println("   - Verifique se não há restrições ativas")
			fmt.Println("   - Confirme que sua conta está ativa")
			fmt.Println("")
			fmt.Println("4️⃣ TESTAR COM DOMÍNIO VERIFICADO:")
			fmt.Println("   - Após verificar o domínio, teste novamente")
			fmt.Println("   - Aguarde até 24h para propagação DNS")
		}
		return
	}

	fmt.Println("✅ Email enviado com sucesso!")
	fmt.Println("🎉 Nenhum erro 'unsigned' encontrado.")
}

func getDomainFromEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) > 1 {
		return parts[1]
	}
	return "seu-dominio.com"
}
