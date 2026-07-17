package testeGO

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected map[string]int
	}{
		{
			name: "Caso Minimo Obrigatorio",
			text: `Casa, casa! A casa é azul.
Árvore; árvore? verde.
Go go Go. IA é útil, mas IA erra.`,
			expected: map[string]int{
				"casa":   3,
				"árvore": 2,
				"azul":   1,
				"verde":  1,
				"útil":   1,
				"mas":    1,
				"erra":   1,
			},
		},
		{
			name:     "Texto Vazio",
			text:     "",
			expected: map[string]int{},
		},
		{
			name:     "Apenas Palavras Curtas",
			text:     "A é do da de no na ir go ia ok me te se",
			expected: map[string]int{},
		},
		{
			name: "Combinacao Casing e Pontuacao",
			text: "Gol! gol, GOL... gol? Gol.",
			expected: map[string]int{
				"gol": 5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWords(tt.text)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CountWords() failed for %s.\nExpected:\n%v\nGot:\n%v", tt.name, tt.expected, result)
			}
		})
	}
}
