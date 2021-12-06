package main

import (
	"fmt"
	"math"
)

func main() {
	const samples = 50
	const tau = 2 * math.Pi
	const rad = tau / samples
	for i := 0; i < samples; i++ {
		samp := math.Sin(rad * float64(i))
		fmt.Printf("%.10f\n", samp)
	}
}
