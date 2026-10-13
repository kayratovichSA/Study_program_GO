package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ParseNumbers принимает целые числа, разделённые пробелами или запятыми.
func ParseNumbers(s string) ([]int, error) {
	fields := strings.Fields(strings.ReplaceAll(s, ",", " "))
	if len(fields) == 0 || len(fields) > 100 {
		return nil, fmt.Errorf("введите от 1 до 100 целых чисел")
	}
	result := make([]int, len(fields))
	for i, f := range fields {
		n, e := strconv.Atoi(f)
		if e != nil || n < -1000000 || n > 1000000 {
			return nil, fmt.Errorf("число %q: нужен integer от −1000000 до 1000000", f)
		}
		result[i] = n
	}
	return result, nil
}

func Calculate(a, b float64, op string) (float64, error) {
	var r float64
	switch op {
	case "+":
		r = a + b
	case "−":
		r = a - b
	case "×":
		r = a * b
	case "÷":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль запрещено")
		}
		r = a / b
	default:
		return 0, fmt.Errorf("выберите операцию")
	}
	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0, fmt.Errorf("результат выходит за диапазон float64")
	}
	return r, nil
}

func Sum(values ...int) int {
	s := 0
	for _, v := range values {
		s += v
	}
	return s
}
func MinMax(values []int) (int, int) {
	lo, hi := values[0], values[0]
	for _, v := range values {
		if v < lo {
			lo = v
		}
		if v > hi {
			hi = v
		}
	}
	return lo, hi
}
func Factorial(n int) uint64 {
	if n <= 1 {
		return 1
	}
	return uint64(n) * Factorial(n-1)
}
func Transform(values []int, f func(int) int) []int {
	r := make([]int, len(values))
	for i, v := range values {
		r[i] = f(v)
	}
	return r
}
func Counter() func() int      { n := 0; return func() int { n++; return n } }
func IncrementValue(n int) int { n++; return n }
func IncrementPointer(n *int)  { *n++ }
func Swap(a, b *int)           { *a, *b = *b, *a }

type Age = int
type Address struct{ City string }
type Student struct {
	Name    string
	Age     Age
	Address Address
}

func (s Student) Description() string {
	return fmt.Sprintf("%s, %d лет, город: %s", s.Name, s.Age, s.Address.City)
}
func (s *Student) Birthday() { s.Age++ }
