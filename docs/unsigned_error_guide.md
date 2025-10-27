# 🚨 Erro "unsigned" - Guia de Solução Rápida

## O que significa o erro "unsigned"?

O erro `unsigned` no Mandrill significa que o **domínio do email remetente não está verificado** ou não é reconhecido como confiável pelo Mandrill.

## ⚡ Soluções Rápidas (em ordem de facilidade)

### 1️⃣ SOLUÇÃO IMEDIATA (5 minutos)
**Use um email @gmail.com ou @yahoo.com que você possui:**

```go
email := mandrill.EmailRequest{
    FromEmail: "seuemail@gmail.com", // Seu email real do Gmail
    FromName:  "Seu Nome",
    // ... resto da configuração
}
```

✅ **Por que funciona:** Gmail e Yahoo são domínios pré-aprovados no Mandrill.

### 2️⃣ VERIFICAR DOMÍNIO PRÓPRIO (30 minutos)

1. **Acesse o painel do Mandrill:**
   - https://mandrillapp.com/settings/sending-domains

2. **Adicione seu domínio:**
   - Clique em "Add Domain"
   - Digite seu domínio (ex: `meusite.com`)

3. **Configure DNS:**
   - Copie os registros DNS fornecidos
   - Adicione no seu provedor de DNS
   - Aguarde propagação (até 24h)

### 3️⃣ VERIFICAR CONTA MANDRILL (10 minutos)

1. **Status da conta:**
   - https://mandrillapp.com/settings
   - Confirme que está ativa

2. **Limites de envio:**
   - Verifique se não atingiu limites
   - Confirme se não há restrições

## 🔧 Ferramentas de Diagnóstico

### Execute o diagnóstico automático:
```bash
make diagnostic-unsigned
```

### Ou diretamente:
```bash
go run examples/diagnostic-unsigned/main.go
```

## 📝 Exemplo Completo Funcionando

```go
package main

import (
    "log"
    "os"
    "github.com/luizhreis/mailchimp-transactional-sender/pkg/mandrill"
)

func main() {
    client := mandrill.NewClient(os.Getenv("MANDRILL_API_KEY"))
    
    email := mandrill.EmailRequest{
        To:          "destinatario@gmail.com",    // Email real do destinatário
        Subject:     "Teste sem erro unsigned",
        FromEmail:   "seuemail@gmail.com",        // SEU email @gmail.com
        FromName:    "Seu Nome Real",
        HTMLContent: "<h1>Teste</h1><p>Email funcionando!</p>",
        TextContent: "Email funcionando!",
    }
    
    err := client.SendEmail(email)
    if err != nil {
        log.Fatalf("Erro: %v", err)
    }
    
    log.Println("✅ Email enviado com sucesso!")
}
```

## ❓ FAQ

**P: Posso usar qualquer email @gmail.com?**
R: Não! Deve ser um email que você realmente possui e tem acesso.

**P: Quanto tempo leva para verificar um domínio?**
R: Entre 15 minutos a 24 horas, dependendo da propagação DNS.

**P: O erro ainda persiste após usar @gmail.com?**
R: Verifique se sua API Key está correta e se a conta Mandrill está ativa.

**P: Posso usar domínios gratuitos como @hotmail.com?**
R: Gmail e Yahoo funcionam melhor. Hotmail pode ter restrições.

## 🆘 Ainda com problemas?

Execute o diagnóstico completo:
```bash
make diagnostic-unsigned
```

O diagnóstico irá:
- ✅ Verificar sua configuração
- ✅ Testar o envio de email
- ✅ Dar orientações específicas para seu caso