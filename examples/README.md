# Exemplos de Uso

Este diretório contém exemplos práticos de como usar o Mailchimp Transactional Sender.

## Estrutura

- `simple-email/` - Exemplo básico de envio de email
- `with-attachment/` - Exemplo de email com anexo

## Como Executar

### Email Simples
```bash
# Usando Make
make example-simple

# Ou diretamente com Go
go run examples/simple-email/main.go
```

### Email com Anexo
```bash
# Usando Make
make example-attachment

# Ou diretamente com Go
go run examples/with-attachment/main.go
```

## Configuração

Antes de executar qualquer exemplo, certifique-se de:

1. **Configurar a API Key do Mandrill:**
   ```bash
   export MANDRILL_API_KEY="sua-api-key-aqui"
   ```

2. **Editar os arquivos de exemplo** para usar seus próprios endereços de email:
   - Altere `destinatario@example.com` para o email do destinatário
   - Altere `seu-email@gmail.com` para seu email verificado no Mandrill
   - Altere `Seu Nome` para o nome do remetente

## Requisitos

- Go 1.21 ou superior
- API Key válida do Mailchimp Transactional (Mandrill)
- Domínio de envio verificado no Mandrill