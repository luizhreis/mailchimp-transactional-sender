# Mailchimp Transactional Email Sender

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Um sender de emails transacionais em Go usando a API do Mailchimp Transactional (Mandrill). Ideal para envio de emails individuais com anexos.

## 🚀 Características

- ✅ Emails transacionais individuais (não campanhas em massa)
- ✅ Suporte completo a anexos
- ✅ Validação automática de configurações
- ✅ Tratamento avançado de erros
- ✅ Diagnóstico de problemas comuns
- ✅ Configuração via variáveis de ambiente

## 📋 Pré-requisitos

- Go 1.21 ou superior
- Conta no [Mailchimp Transactional](https://mailchimp.com/transactional/)
- API Key do Mailchimp Transactional

## 🔧 Instalação

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/seu-usuario/mailchimp-transactional-sender.git
   cd mailchimp-transactional-sender
   ```

2. **Instale dependências:**
   ```bash
   go mod tidy
   ```

3. **Configure as variáveis de ambiente:**
   ```bash
   export MANDRILL_API_KEY="sua-api-key-aqui"
   ```
   
   Ou use o script automático:
   ```bash
   chmod +x setup.sh
   ./setup.sh
   ```

## 🚀 Uso Rápido

1. **Edite as configurações** em `cmd/main.go`:
   ```go
   destinatario := "destinatario@example.com"  // Seu email de teste
   remetente := "seu-email@gmail.com"          // Seu email verificado
   nomeRemetente := "Seu Nome"                 // Seu nome
   ```

2. **Execute:**
   ```bash
   go run cmd/main.go
   ```

## 📁 Estrutura do Projeto

```
.
├── cmd/
│   └── main.go              # Aplicação principal
├── pkg/
│   └── mandrill/
│       ├── client.go        # Cliente Mandrill
│       ├── types.go         # Tipos e estruturas
│       └── errors.go        # Tratamento de erros
├── examples/
│   ├── simple_email.go      # Exemplo básico
│   └── with_attachment.go   # Exemplo com anexo
├── docs/
│   ├── troubleshooting.md   # Resolução de problemas
│   └── api_errors.md        # Erros comuns da API
├── scripts/
│   └── setup.sh            # Script de configuração
├── .env.example            # Exemplo de variáveis de ambiente
├── .gitignore              # Arquivos ignorados pelo Git
├── go.mod                  # Dependências do Go
├── go.sum                  # Checksums das dependências
├── LICENSE                 # Licença do projeto
└── README.md               # Este arquivo
```

## 📖 Exemplos

### Email Simples
```go
package main

import (
    "log"
    "os"
    "github.com/seu-usuario/mandrill-sender/pkg/mandrill"
)

func main() {
    client := mandrill.NewClient(os.Getenv("MANDRILL_API_KEY"))
    
    email := mandrill.EmailRequest{
        To:          "usuario@example.com",
        Subject:     "Olá do Mandrill!",
        FromEmail:   "seu-email@gmail.com",
        FromName:    "Seu Nome",
        HTMLContent: "<h1>Olá!</h1><p>Este é um teste.</p>",
        TextContent: "Olá! Este é um teste.",
    }
    
    if err := client.SendEmail(email); err != nil {
        log.Fatal(err)
    }
}
```

### Email com Anexo
```go
attachment, err := mandrill.LoadAttachment("documento.pdf")
if err != nil {
    log.Fatal(err)
}

email := mandrill.EmailRequest{
    To:         "usuario@example.com",
    Subject:    "Documento em anexo",
    FromEmail:  "seu-email@gmail.com", 
    FromName:   "Seu Nome",
    HTMLContent: "<p>Segue documento em anexo.</p>",
    Attachment: attachment,
}
```

## 🚨 Problemas Comuns

### Email rejeitado com `unsigned` ⚠️ **MAIS COMUM**
Este é o erro mais frequente. O domínio do remetente não está verificado.

**🚀 Solução rápida:**
```bash
# Execute o diagnóstico automático
make diagnostic-unsigned
```

**🎯 Soluções:**
1. **Imediata:** Use um email @gmail.com/@yahoo.com que você possui
2. **Permanente:** Verifique seu domínio em https://mandrillapp.com/settings/sending-domains
3. **Diagnóstico:** Consulte [docs/unsigned_error_guide.md](docs/unsigned_error_guide.md)

### Email rejeitado com `recipient-domain-mismatch`  
- Desative "Domain Matching" nas configurações do Mandrill
- Configure tracking domains para seu domínio

### Outros erros
Consulte [docs/troubleshooting.md](docs/troubleshooting.md) para soluções detalhadas.

## 🧪 Testes

```bash
# Executar todos os testes
go test ./...

# Teste com coverage
go test -cover ./...

# Diagnóstico de problemas
make diagnostic-unsigned

# Executar exemplos
make example-simple
make example-attachment
```

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

## 🔗 Links Úteis

- [Mailchimp Transactional](https://mailchimp.com/transactional/)
- [Documentação da API](https://mailchimp.com/developer/transactional/api/)
- [Status da API](https://status.mailchimp.com/)

## ⭐ Suporte

Se este projeto te ajudou, considere dar uma estrela! ⭐

Para dúvidas ou problemas, abra uma [issue](https://github.com/seu-usuario/mandrill-sender/issues).