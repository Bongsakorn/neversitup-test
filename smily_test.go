package main

import "testing"

func TestSmilyCase1(t *testing.T) {
	text := []string{":)", ";(", ";}", ":-D"}
	expected := 2
	result := CountSmilyFace(text)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)

	}
}

func TestSmilyCase2(t *testing.T) {
	text := []string{";D", ":-(", ":-)", ";~)"}
	expected := 3
	result := CountSmilyFace(text)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)

	}
}

func TestSmilyCase3(t *testing.T) {
	text := []string{";]", ":[", ";*", ":$", ";-D"}
	expected := 1
	result := CountSmilyFace(text)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
