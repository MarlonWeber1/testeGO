# testeGO

Este repositório contém a atividade de construção e validação de testes automáticos em Go para uma função de contagem de frequência de palavras, desenvolvida em conjunto com o assistente de Inteligência Artificial.

## Descrição da Atividade
O objetivo desta atividade é implementar uma função em Go que realiza a contagem da frequência de palavras em um texto, aplicando regras de normalização e filtragem, e criar testes automatizados robustos para validar sua corretude.

O foco da atividade consiste em exercitar a análise crítica sobre a especificação do problema, o código produzido e a eficácia das suítes de teste geradas pela IA.

## Como Executar os Testes
Para executar a suíte de testes automáticos, certifique-se de possuir o Go instalado em sua máquina e execute o seguinte comando no diretório raiz do projeto:

```bash
go test -v
```

## Explicação da Função `CountWords`
A assinatura da função é:
```go
func CountWords(text string) map[string]int
```

A função processa o texto de entrada seguindo estas etapas:
1. **Tokenização**: Divide o texto original em palavras individuais utilizando espaços em branco como delimitadores (através de `strings.Fields`).
2. **Normalização de Caixa**: Converte cada palavra para minúsculas usando `strings.ToLower`.
3. **Limpeza de Pontuação**: Remove caracteres de pontuação simples (vírgula `,`, ponto final `.`, ponto de exclamação `!`, ponto de interrogação `?`, ponto e vírgula `;`, dois-pontos `:`).
4. **Filtro de Tamanho**: Descarta palavras que tenham menos de 3 caracteres. Para isso, utiliza `utf8.RuneCountInString` garantindo suporte completo a caracteres acentuados em UTF-8 (como "é" ou "á" que contam como 1 caractere, mas ocupam múltiplos bytes).
5. **Contagem**: Contabiliza as palavras válidas restantes e retorna um mapa contendo as frequências.

## Casos de Teste Implementados
A suíte de testes em `wordcount_test.go` foi estruturada utilizando a técnica de *Table-Driven Tests* e contém os seguintes cenários:
1. **Caso Mínimo Obrigatório**: Valida a especificação exata fornecida na atividade:
   - Texto de entrada: `"Casa, casa! A casa é azul. Árvore; árvore? verde. Go go Go. IA é útil, mas IA erra."`
   - Valida se as palavras `"a"`, `"é"`, `"go"` e `"ia"` são descartadas e as demais frequências computadas de forma exata.
2. **Texto Vazio**: Valida que uma string vazia retorna um mapa de resultados vazio, prevenindo erros de execução.
3. **Apenas Palavras Curtas**: Valida se uma string contendo apenas conectivos e termos com menos de 3 caracteres resulta em um mapa de resultados vazio.
4. **Combinação de Casing e Pontuação**: Testa a robustez da limpeza de pontuação consecutiva e normalização com o termo `"Gol! gol, GOL... gol? Gol."`, garantindo que todas se convertam no termo `"gol"` com contagem igual a 5.

## Resultado do `go test`
Ao executar os testes locais, o resultado obtido foi:

```
=== RUN   TestCountWords
=== RUN   TestCountWords/Caso_Minimo_Obrigatorio
=== RUN   TestCountWords/Texto_Vazio
=== RUN   TestCountWords/Apenas_Palavras_Curtas
=== RUN   TestCountWords/Combinacao_Casing_e_Pontuacao
--- PASS: TestCountWords (0.00s)
    --- PASS: TestCountWords/Caso_Minimo_Obrigatorio (0.00s)
    --- PASS: TestCountWords/Texto_Vazio (0.00s)
    --- PASS: TestCountWords/Apenas_Palavras_Curtas (0.00s)
    --- PASS: TestCountWords/Combinacao_Casing_e_Pontuacao (0.00s)
PASS
ok  	testeGO	0.599s
```
