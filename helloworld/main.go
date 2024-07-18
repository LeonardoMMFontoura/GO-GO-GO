package main

import "fmt"

const prefixoEmPortugues = "Olá, "

func Ola(nome string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefixoEmPortugues + nome
}

func main() {
	fmt.Println(Ola("Leonardo"))
}
