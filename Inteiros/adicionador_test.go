package inteiros

import (
  "testing"
  "fmt"
)

func TestAdicionador(t *testing.T) {
  soma := Adiciona(23, 32)
  esperado := 55
  if soma != esperado {
    t.Errorf("O resultado foi '%d' mas o esperado era '%d'", soma, esperado)
  }
}

func ExampleAdiciona() {
  soma := Adiciona(4, 9)
  fmt.Println(soma)
  // Output: 13
}
