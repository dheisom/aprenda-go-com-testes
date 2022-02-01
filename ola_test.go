package main

import "testing"

func TestOla(t *testing.T) {
  resultado := Ola("Dheisom")
  esperado := "Olá, Dheisom!"
  if resultado != esperado {
    t.Errorf("'%s' esperado, recebido '%s'", esperado, resultado)
  }
}
