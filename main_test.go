package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(20, 30)
	expected := 50
	if result != expected {
		t.Errorf("Add(20, 30) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(50, 30)
	expected := 20
	if result != expected {
		t.Errorf("Subtract(50, 30) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(20, 30)
	expected := 600
	if result != expected {
		t.Errorf("Multiply(20, 30) = %d; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {

	result, err := Divide(10, 2)
	expected := 5
	if err != nil || result != expected {
		t.Errorf("Divide(10, 2) = %d, %v; want %d, <nil>", result, err, expected)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Expected error for division by zero, got nil")
	}
}
