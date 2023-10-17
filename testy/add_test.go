package testy

import "testing"

func TestAdd(t *testing.T) {
	expected := 10
	actual := add(5, 5)
	if actual != expected {
		t.Errorf("Součet vyšel špatně, očekávaný výsledek (%v) není roven získanému výsledku (%v)\n", expected, actual)
	}
}

func TestSubstract(t *testing.T) {
	expected := 5
	actual := substract(10, 5)
	if actual != expected {
		t.Errorf("Rozdíl vyšel špatně, očekávaný výsledek (%v) není roven získanému výsledku (%v)\n", expected, actual)
	}
}

func TestDoMath(t *testing.T) {
	expected := 10
	actual := doMath(5, 5, add)
	if actual != expected {
		t.Errorf("Matikář vyšel špatně, očekávaný výsledek (%v) není roven získanému výsledku (%v)\n", expected, actual)
	}
}
