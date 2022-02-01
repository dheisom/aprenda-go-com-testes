package main

import "fmt"

const espanhol = "espanhol"
const frances = "francês"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func prefixoDeSaudacao(idioma string) (prefixo string) {
  switch idioma {
  case espanhol:
    prefixo = prefixoOlaEspanhol
  case frances:
    prefixo = prefixoOlaFrances
  default:
    prefixo = prefixoOlaPortugues
  }
  return
}

func Ola(nome string, idioma string) string {
  if nome == "" {
    nome = "Mundo"
  }
  return prefixoDeSaudacao(idioma) + nome + "!"
}

func main() {
  fmt.Println(Ola("Mundo", ""))
}
