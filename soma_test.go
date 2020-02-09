package main

import "testing"
import "reflect"

func TestSoma(t *testing.T) {
	//t.Helper()

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		if esperado != resultado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})
}

func TestSomaTudo(t *testing.T) {

	recebido := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(recebido, esperado) {
		t.Errorf("recebido %v esperado %v", recebido, esperado)
	}
}
