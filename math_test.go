package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é invaálido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSomaErrada(t *testing.T) {
	total := SomaErrada(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é invaálido: Resultado %d. Esperado: %d", total, 30)
	}
}