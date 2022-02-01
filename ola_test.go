package main

import "testing"

func TestOla(t *testing.T) {
  verificarMensagemCorreta := func(t *testing.T, resultado, esperado string) {
    t.Helper()
    if resultado != esperado {
      t.Errorf("O esperado era '%s' mas o resultado foi '%s'!", esperado, resultado)
    }
  }
  t.Run("Diz olá para as pessoas", func(t *testing.T) {
    resultado := Ola("Dheisom")
    esperado := "Olá, Dheisom!"
    verificarMensagemCorreta(t, resultado, esperado)
  })
  t.Run("Diz olá para o mundo se o nome for vazio", func(t *testing.T) {
    resultado := Ola("")
    esperado := "Olá, Mundo!"
    verificarMensagemCorreta(t, resultado, esperado)
  })
}
