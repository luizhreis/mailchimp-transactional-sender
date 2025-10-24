#!/bin/bash

# Script para configurar o Mailchimp Transactional (Mandrill)
echo "=== Configuração do Mailchimp Transactional (Mandrill) ==="
echo

echo "🚀 Este script vai configurar as variáveis de ambiente para o Mandrill."
echo "📧 Mailchimp Transactional é ideal para emails individuais com anexos."
echo

read -p "Digite sua Mandrill API Key: " API_KEY

if [ -z "$API_KEY" ]; then
    echo "❌ Erro: API Key não pode estar vazia!"
    exit 1
fi

export MANDRILL_API_KEY="$API_KEY"

echo
echo "=== Configuração completa! ==="
echo "✅ MANDRILL_API_KEY: $MANDRILL_API_KEY"
echo

echo "Para tornar permanente, adicione ao seu ~/.zshrc:"
echo "export MANDRILL_API_KEY=\"$MANDRILL_API_KEY\""
echo

echo "📝 Próximos passos:"
echo "1. Edite o arquivo mandrill_sender.go"
echo "2. Altere as linhas com destinatário, remetente e nome"
echo "3. Execute: go run mandrill_sender.go"
echo

echo "💡 Dica: O arquivo exemplo.txt será anexado automaticamente se existir."

# Verificar se o arquivo exemplo existe
if [ -f "exemplo.txt" ]; then
    echo "📎 Arquivo exemplo.txt encontrado - será usado como anexo."
else
    echo "📎 Criando arquivo exemplo.txt..."
    cat > exemplo.txt << 'EOF'
# Arquivo de Exemplo para Anexo

Este é um arquivo de teste criado automaticamente.

Conteúdo:
- Data: 22 de outubro de 2025
- Projeto: Mailchimp Transactional Email Sender
- Linguagem: Go

Este arquivo será anexado ao email como demonstração.

Você pode substituir este arquivo por qualquer documento que desejar anexar.
EOF
    echo "✅ Arquivo exemplo.txt criado com sucesso!"
fi

echo
echo "🚀 Agora você pode executar: go run mandrill_sender.go"