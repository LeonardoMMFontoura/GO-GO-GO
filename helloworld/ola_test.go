package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Leo")
	esperado := "Ola, Leo"
	if resultado != esperado {
		t.Errorf("Resultado '%s', eperado '%s'", resultado, esperado)
	}
}
