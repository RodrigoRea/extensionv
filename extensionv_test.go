package extensionv 

import (
	"testing"
)

func TestDescreva(t *testing.T) {

	var n float64 = 7199002801000500.89

	texto, _ := Descreva(n)
	if texto == "" {
		t.Errorf("Não consegui escrever o número")
	} 
}