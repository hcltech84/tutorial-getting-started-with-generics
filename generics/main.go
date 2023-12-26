package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumNumber[K comparable, V Number](m map[K]V) V{
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := map[string]int64{
		"one": 10,
		"two": 25,
	}
	floats := map[string]float64{
		"one": 10.06,
		"two": 2.01,
	}

	fmt.Println(SumNumber[string, int64](ints))

	// You can omit the type constraint argument in calling code if the Go compiler can infer it
	// So, instead of writting
	// fmt.Println(SumNumber[string, float64](floats))
	// You can write
	fmt.Println(SumNumber(floats))
}