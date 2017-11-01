package main

// Test, Benchmark and Example
import "testing"

func TestAddDivisors_3and5(t *testing.T) {

	until := 5
	expected := 8

	result := addDivisors_3and5(until)

	if result != expected {
		t.Fatalf("%d Should be equals to %d", result, expected)
	}
}
