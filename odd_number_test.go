package main

import (
	"testing"
)

func TestOddNumberCase1(t *testing.T) {
	text := []int{7}
	expected := 7
	result := FindOddNumber(text)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestOddNumberCase2(t *testing.T) {
	text := []int{0}
	expected := 0
	result := FindOddNumber(text)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestOddNumberCase3(t *testing.T) {
	text := []int{1, 1, 2}
	expected := 2
	result := FindOddNumber(text)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestOddNumberCase4(t *testing.T) {
	text := []int{0, 1, 0, 1, 0}
	expected := 0
	result := FindOddNumber(text)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestOddNumberCase5(t *testing.T) {
	text := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}
	expected := 4
	result := FindOddNumber(text)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
