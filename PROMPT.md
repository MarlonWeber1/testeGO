# Registro de Uso de Inteligência Artificial

## Dados Gerais
- **Ferramenta de IA utilizada:** Gemini 3.5 Flash (via Antigravity Coding Assistant)
- **Ambiente utilizado:** Windows Powershell, Go 1.20+ (ou versão instalada localmente)

---

## Etapa 1 — Compreensão do problema

### Prompt Utilizado
> Quero implementar em Go uma função CountWords(text string) map[string]int. A função deve contar a frequência das palavras em um texto, converter palavras para minúsculas, remover pontuação simples e ignorar palavras com menos de 3 caracteres. Explique a especificação e os principais cuidados de implementação.

### Resumo da Resposta da IA
A especificação solicita uma função `CountWords` que recebe um texto (string) e devolve um mapa (`map[string]int`) contendo a frequência de cada palavra que atenda a certos critérios de normalização e filtragem.
Os principais pontos destacados pela IA foram:
1. **Tokenização**: O texto precisa ser quebrado em palavras individuais. A separação por espaços em branco (`strings.Fields`) é ideal para isso.
2. **Normalização de Casing**: Toda palavra deve ser convertida para minúsculas usando `strings.ToLower` para evitar que "Casa" e "casa" sejam contadas separadamente.
3. **Limpeza de Pontuação**: Deve-se remover pontuações comuns (como `,`, `.`, `!`, `?`, `;`, `:`). O cuidado aqui é remover pontuações sem corromper as palavras ou remover caracteres acentuados.
4. **Filtragem por Tamanho**: Palavras com menos de 3 caracteres devem ser descartadas.
5. **Suporte a Unicode/UTF-8 (Acentuação)**: Em Go, strings são sequências de bytes. Caracteres acentuados (como "á" ou "é") ocupam mais de um byte em UTF-8. Para medir o tamanho da palavra corretamente, deve-se usar o comprimento em runas (`utf8.RuneCountInString(word)`) em vez de `len(word)`.

### Cuidados Identificados
- **Acentos e UTF-8**: O uso de `len()` em strings com acentos retornará um tamanho em bytes maior que o número real de caracteres (runas). Se usarmos `len("é")`, o resultado será 2, mas o tamanho em caracteres é 1. Portanto, para verificar se a palavra tem menos de 3 caracteres, devemos contar runas.
- **Pontuação no final de palavras**: Em "casa!", o sinal "!" faz parte da palavra após a separação inicial por espaços. Devemos limpar os caracteres de pontuação do início e fim (ou de toda a palavra), mantendo as letras.
- **Pontuações internas ou isoladas**: Pontuações sozinhas podem virar palavras vazias após a limpeza, que devem ser descartadas.

### Decisões Tomadas
- Usar `strings.Fields` para realizar a quebra de palavras por qualquer espaço em branco.
- Tratar caracteres usando funções do pacote `unicode` para preservar letras com acentos (como `unicode.IsLetter` ou `unicode.IsSpace`), removendo pontuações e símbolos simples.
- Para verificar o limite de tamanho menor que 3, usaremos `utf8.RuneCountInString` para garantir a contagem correta de runas/caracteres.

---

## Etapa 2 — Implementação da função

### Prompt Utilizado
> Implemente em Go a função CountWords(text string) map[string]int. A função deve converter palavras para minúsculas, remover pontuação simples, ignorar palavras com menos de 3 caracteres e retornar um map[string]int com a frequência das palavras.

### Código Gerado
```go
package testeGO

import (
	"strings"
	"unicode/utf8"
)

func CountWords(text string) map[string]int {
	frequencies := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		lowerWord := strings.ToLower(word)

		var cleanedBuilder strings.Builder
		for _, r := range lowerWord {
			if r == ',' || r == '.' || r == '!' || r == '?' || r == ';' || r == ':' {
				continue
			}
			cleanedBuilder.WriteRune(r)
		}
		cleaned := cleanedBuilder.String()

		if utf8.RuneCountInString(cleaned) >= 3 {
			frequencies[cleaned]++
		}
	}

	return frequencies
}
```

### Verificações Realizadas
- **O código compila?** Sim, compilou com sucesso após ajuste nas importações (removido `"unicode"` que não estava sendo usado).
- **A função tem a assinatura esperada?** Sim, `func CountWords(text string) map[string]int`.
- **A normalização está correta?** Sim, usamos `strings.ToLower` para converter todos os caracteres para minúsculas antes de contar.
- **Acentos são preservados?** Sim, o mapeamento de pontuação foca apenas em `,`, `.`, `!`, `?`, `;`, `:`, mantendo as letras acentuadas intactas (por exemplo, "árvore", "útil"). E o comprimento é checado por runas, garantindo que acentos (que usam múltiplos bytes) sejam contados como 1 caractere único.

---

## Etapa 3 — Criação do teste automático

### Prompt Utilizado
> Crie um teste automático em Go, usando o pacote padrão testing, para validar a função CountWords. O teste deve usar a seguinte entrada:
> 
> Casa, casa! A casa é azul.
> Árvore; árvore? verde.
> Go go Go. IA é útil, mas IA erra.
> 
> O resultado esperado deve ser:
> casa: 3
> árvore: 2
> azul: 1
> verde: 1
> útil: 1
> mas: 1
> erra: 1
> 
> O teste deve comparar o mapa produzido pela função com o mapa esperado e falhar caso haja qualquer diferença.
> O teste deve ser executado com:
> go test

### Teste Gerado
```go
package testeGO

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	text := `Casa, casa! A casa é azul.
Árvore; árvore? verde.
Go go Go. IA é útil, mas IA erra.`

	expected := map[string]int{
		"casa":   3,
		"árvore": 2,
		"azul":   1,
		"verde":  1,
		"útil":   1,
		"mas":    1,
		"erra":   1,
	}

	result := CountWords(text)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CountWords() failed.\nExpected:\n%v\nGot:\n%v", expected, result)
	}
}
```

### Resultado da Execução
O teste foi executado com o comando `go test -v` no terminal do Windows Powershell e passou sem problemas:
```
=== RUN   TestCountWords
--- PASS: TestCountWords (0.00s)
PASS
ok  	testeGO	0.663s
```

### Erros Encontrados & Correções Realizadas
- Nenhum erro foi encontrado durante a compilação ou execução do teste básico. O mapa produzido coincidiu perfeitamente com o esperado.

---

## Etapa 4 — Revisão crítica do teste

### Prompt Utilizado
> Revise o teste automático abaixo. Verifique se ele testa corretamente a conversão para minúsculas, a remoção de pontuação, o descarte de palavras com menos de 3 caracteres, a contagem de repetições e a comparação completa entre o mapa esperado e o mapa produzido.
> [Código do Teste de wordcount_test.go]

### Resumo da Resposta da IA
A IA revisou o teste e confirmou as seguintes validações:
1. **Conversão para minúsculas**: Testado com sucesso, pois "Casa", "casa" e "casa!" são consolidados no mesmo termo "casa".
2. **Remoção de pontuação**: Testado com sucesso, pois termos como "casa!" perdem a pontuação e viram "casa", combinando-se com as demais ocorrências.
3. **Descarte de palavras curtas**: Testado com sucesso. Palavras como "A", "é", "Go", "go", "IA" foram corretamente excluídas por possuírem menos de 3 caracteres. (Nota: "mas" possui 3 caracteres, logo foi incluída).
4. **Contagem de repetições**: Correto. As palavras "casa" e "árvore" tiveram contagem 3 e 2 respectivamente.
5. **Comparação de mapas**: O uso de `reflect.DeepEqual` é apropriado e rigoroso. Ele falhará se houver chaves extras, chaves ausentes ou valores discrepantes.

**Sugestões da IA:**
- Utilizar a técnica de **Table-Driven Tests** (Testes Baseados em Tabelas) para estruturar melhor o teste, permitindo testar múltiplos casos de forma limpa e isolada usando `t.Run`.
- Adicionar casos de teste com strings vazias, apenas pontuações, caracteres especiais adicionais e espaços extras.

### Decisões sobre as sugestões da IA
- **Sugestões Aceitas**: 
  - Estruturar os testes usando Table-Driven Tests.
  - Adicionar casos de teste para tratar cenários extremos (ex: string vazia, apenas palavras curtas, combinações variadas de casing e pontuação).
- **Sugestões Rejeitadas**:
  - Nenhuma sugestão foi rejeitada, todas as melhorias sugeridas são boas práticas de engenharia de software em Go.
- **Conclusão sobre a cobertura**: O teste mínimo atual valida corretamente a especificação, mas estendê-lo garantirá robustez contra regressões.


