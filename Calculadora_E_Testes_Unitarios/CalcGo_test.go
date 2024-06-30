package main

import "testing"

func TestSomaCorrect(t *testing.T) { //ShouldSumCorrect
	//arrange
	teste := soma(3, 2, 1)
	//act
	resultado := 6
	//assert
	if teste != resultado { //assert
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSomaIncorrect(t *testing.T) { //ShouldSumIncorrect
	teste := soma(3, 2, 1) //arrange
	resultado := 7         //act

	if teste != resultado { //assert
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMultCorrect(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMultIncorrect(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 2560
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSubCorrect(t *testing.T) {
	teste := subtrai(10, 5)
	resultado := -5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSubIncorrect(t *testing.T) {
	teste := subtrai(10, 10)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestDivCorrect(t *testing.T) {
	teste := divide(10)
	resultado := 10
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestDivIncorrect(t *testing.T) {
	teste := divide(10)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
