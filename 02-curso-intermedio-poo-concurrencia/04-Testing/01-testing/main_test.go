package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 5)

	if total != 10 {
		t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	}
}

func TestSum_(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, table := range tables {
		total := Sum(table.a, table.b)
		if total != table.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, table.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	result := GetMax(10, 15)

	if result != 15 {
		t.Errorf("GetMax was incorrect, got %d expected %d", result, 15)
	}
}

func TestGetMax_X(t *testing.T) {
	//act
	result := GetMax(15, 10)

	//asserts
	if result != 15 {
		t.Errorf("GetMax was incorrect, got %d expected %d", result, 15)
	}
}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		a int
		b int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.b {
			t.Errorf("Fibonacci was incorrect, got %d expected %d", fib, item.b)
		}
	}
}
