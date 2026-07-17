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
