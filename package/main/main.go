package main

import (
	"fmt"

	"github.com/Vlad1slavIP74/GO_PACKAGE/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
