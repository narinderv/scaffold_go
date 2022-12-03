package main

import (
	"fmt"
	"testing"
)

func Test_greeting(t *testing.T) {
	fmt.Println("Running Test for greeting()...")

	input := "narinderv"
	expectedOutput := "Hello narinderv"

	if actualOutput := greeting(input); actualOutput != expectedOutput {
		t.Errorf("greeting() = %s, expected = %s", actualOutput, expectedOutput)
	}
}

func Test_fiboGenerator(t *testing.T) {
	fmt.Println("Running Test for printFiboSeries()...")
	maxCount := 5
	expectedOutput := []int{1, 1, 2, 3, 5}

	actualOutput, err := printFiboSeries(maxCount)

	if err != nil {
		t.Errorf("printFiboSeries() returns error %s", err)
	} else {
		for i, val := range actualOutput {
			if val != expectedOutput[i] {
				t.Errorf("printFiboSeries() Element %d = %d, expected = %d", i, val, expectedOutput[i])
			}
		}
	}

}
