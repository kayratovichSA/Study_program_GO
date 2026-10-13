package main

import (
	"reflect"
	"testing"
)

func TestLogic(t *testing.T) {
	if _, e := Calculate(1, 0, "÷"); e == nil {
		t.Fatal("division by zero")
	}
	for _, c := range []struct {
		op   string
		want float64
	}{{"+", 15}, {"−", 9}, {"×", 36}, {"÷", 4}} {
		r, e := Calculate(12, 3, c.op)
		if e != nil || r != c.want {
			t.Fatalf("calculator %s: %v %v", c.op, r, e)
		}
	}
	v, e := ParseNumbers("5, 2 9 -1 4")
	if e != nil || Sum(v...) != 19 {
		t.Fatal(v, e)
	}
	lo, hi := MinMax(v)
	if lo != -1 || hi != 9 {
		t.Fatal(lo, hi)
	}
	for _, s := range []string{"", "abc", "1000001"} {
		if _, e := ParseNumbers(s); e == nil {
			t.Fatal("accepted", s)
		}
	}
	if Factorial(0) != 1 || Factorial(20) != 2432902008176640000 {
		t.Fatal("factorial")
	}
	if !reflect.DeepEqual(Transform([]int{2, 3}, func(x int) int { return x * x }), []int{4, 9}) {
		t.Fatal("transform")
	}
	next := Counter()
	if next() != 1 || next() != 2 {
		t.Fatal("closure")
	}
	a, b := 3, 8
	Swap(&a, &b)
	if a != 8 || b != 3 {
		t.Fatal("swap")
	}
	IncrementPointer(&a)
	if a != 9 || IncrementValue(a) != 10 || a != 9 {
		t.Fatal("increment")
	}
	s := Student{"Анна", 20, Address{"Кызылорда"}}
	s.Birthday()
	if s.Age != 21 {
		t.Fatal("method")
	}
}
