package main

import (
	"errors"
	"math"
)

type Algorithms struct{}

func (a *Algorithms) MinPositive(arr []int) (int, error) {
	min := math.MaxInt
	found := false
	for _, num := range arr {
		if num > 0 && num < min {
			min = num
			found = true
		}
	}
	if !found {
		return 0, errors.New("no positive numbers found")
	}
	return min, nil
}

func (a *Algorithms) SumNegative(arr []int) int {
	sum := 0
	for _, num := range arr {
		if num < 0 {
			sum += num
		}
	}
	return sum
}

func (a *Algorithms) Fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("n must be non-negative")
	}
	if n == 0 {
		return 0, nil
	}
	if n == 1 {
		return 1, nil
	}
	a1, a2 := 0, 1
	for i := 2; i <= n; i++ {
		a1, a2 = a2, a1+a2
	}
	return a2, nil
}

func (a *Algorithms) Current(voltage, resistance float64) (float64, error) {
	if resistance == 0 {
		return 0, errors.New("resistance cannot be zero")
	}
	return voltage / resistance, nil
}

func main() {
	alg := &Algorithms{}

	minPos, _ := alg.MinPositive([]int{-3, 4, 2, -1})
	sumNeg := alg.SumNegative([]int{-3, 4, -2, -1})
	fib, _ := alg.Fibonacci(10)
	current, _ := alg.Current(12, 4)

	println("Min Positive:", minPos)
	println("Sum Negative:", sumNeg)
	println("Fibonacci(10):", fib)
	println("Current (12V, 4Ω):", current)
}
