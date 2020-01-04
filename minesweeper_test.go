package minesweeper1d

import (
	"testing"
)

func TestCountMines(t *testing.T) {
	BasicCountMinesTest(t, []int{0, 0, 0, 0, 0, 0}, 0)
	BasicCountMinesTest(t, []int{1, 1, 0, 0, 0, 0}, 1)
	BasicCountMinesTest(t, []int{1, 2, 2, 1, 0, 1}, 4)
	BasicCountMinesTest(t, []int{1, 0, 0, 0, 0, 0}, -1)
	BasicCountMinesTest(t, []int{0, 0, 0, 1, 0, 0}, -1)
	BasicCountMinesTest(t, []int{0, 0, 0, 1, 2, 0}, -1)
	BasicCountMinesTest(t, []int{0, 0, 0, 0, 2, 0}, -1)
}

func BasicCountMinesTest(t *testing.T, input []int, expected int) {
	if output := CountMines(input); output != expected {
		t.Errorf("Test Failed with input %v \nExpecting : \t\t %v\nBut got : \t\t %v", input, expected, output)
	}
}

func TestCountMinesString(t *testing.T) {
	BasicCountMinesStringTest(t, "0, 0, 0, 0, 0, 0", 0, false)
	BasicCountMinesStringTest(t, "1, 1, 0, 0, 0, 0", 1, false)
	BasicCountMinesStringTest(t, "1, 2, 2, 1, 0, 1", 4, false)
	BasicCountMinesStringTest(t, "1, 0, 0, 0, 0, 0", -1, true)
	BasicCountMinesStringTest(t, "0, 0, 0, 1, 0, 0", -1, true)
	BasicCountMinesStringTest(t, "0, 0, 0, 1, 2, 0", -1, true)
	BasicCountMinesStringTest(t, "0, 0, 0, 0, 2, 0", -1, true)
	BasicCountMinesStringTest(t, "0, 0, 0, a, 2, 0", -1, true)
}

func BasicCountMinesStringTest(t *testing.T, input string, expected int, expectingErr bool) {
	output, err := CountMinesString(input)
	if output != expected {
		t.Errorf("Test Failed with input %v \nExpecting : \t\t %v\nBut got : \t\t %v", input, expected, output)
	}
	if expectingErr && err == nil {
		t.Errorf("Test Failed with input %v \nExpecting error but got no error\nOutput : \t\t%v", input, output)
	} else if !expectingErr && err != nil {
		t.Errorf("Test Failed with input %v \nExpecting no error but got error : \t\t%s", input, err.Error())
	}
}
