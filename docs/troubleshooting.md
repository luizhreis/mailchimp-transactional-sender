# 🚨 SOLUÇÃO PARA: recipient-domain-mismatch

## ❌ Problema identificado:
Reject Reason: `recipient-domain-mismatch`

Este erro ocorre quando o Mailchimp Transactional tem configurações de **IP dedicado** ou **restrições de domínio** ativadas que estão causando conflito entre o domínio do remetente (`lreis.com.br`) e o domínio do destinatário (`gmail.com`).

## ✅ SOLUÇÕES (em ordem de prioridade):

### **Solução 1: Desativar Domain Matching (MAIS PROVÁVEL)**

1. **Acesse:** https://mandrillapp.com/settings/sending-domains

2. **Procure por:** "Dedicated IP Domain Matching" ou "Domain Restrictions"

3. **DESATIVE** qualquer opção que force matching de domínios

4. **Salve** as configurações

### **Solução 2: Configurar Tracking Domains**

1. **Acesse:** https://mandrillapp.com/settings/tracking-domains

2. **Adicione:** `lreis.com.br` como tracking domain

3. **Configure** o CNAME conforme instruções:
   ```dns
   CNAME email.lreis.com.br mandrillapp.com
   ```

### **Solução 3: Teste com mesmo domínio**

Temporariamente, teste enviando de `luiz@lreis.com.br` para outro email do mesmo domínio:

```go
destinatario := "outro-email@lreis.com.br"  // Mesmo domínio
```

### **Solução 4: Verificar configurações da conta**

1. **Acesse:** https://mandrillapp.com/settings/account

2. **Verifique:** se há alguma configuração de "Whitelist" ou "Domain Restrictions"

3. **Desative** restrições temporariamente para teste

### **Solução 5: Usar subdomínio**

Configure um subdomínio específico para emails:

1. **Configure DNS:**
   ```dns
   CNAME mail.lreis.com.br mandrillapp.com
   ```

2. **Use no código:**
   ```go
   remetente := "luiz@mail.lreis.com.br"
   ```

## 🔍 DIAGNÓSTICO ADICIONAL:

### Verificar configurações atuais:

1. **Acesse:** https://mandrillapp.com/settings/

2. **Verifique estas seções:**
   - Sending Domains (deve mostrar `lreis.com.br` verificado)
   - Tracking Domains
   - Dedicated IPs (se aplicável)
   - Account Settings

### Teste com diferentes combinações:

```go
// Teste 1: Gmail para Gmail
remetente := "seu-gmail@gmail.com"
destinatario := "outro-gmail@gmail.com"

// Teste 2: Seu domínio para seu domínio
remetente := "luiz@lreis.com.br"
destinatario := "contato@lreis.com.br"

// Teste 3: Seu domínio para Gmail (problema atual)
remetente := "luiz@lreis.com.br"
destinatario := "luizhreis.gris@gmail.com"
```

## 📞 SUPORTE MAILCHIMP:

Se nenhuma solução funcionar:

1. **Acesse:** https://mailchimp.com/contact/
2. **Mencione:** "recipient-domain-mismatch error"
3. **Informe:** que seu domínio está verificado mas ainda há erro

## 🚀 TESTE RÁPIDO:

Para confirmar se é problema de configuração, teste temporariamente:

```bash
# Edite mandrill_sender.go linha ~327
remetente := "luizhreis.gris@gmail.com"  # Mesmo domínio do destinatário
```

Se funcionar, confirma que é problema de domain matching.

---

**🎯 AÇÃO IMEDIATA: Acesse as configurações do Mandrill e desative qualquer "Domain Matching" ou "Domain Restrictions"!**