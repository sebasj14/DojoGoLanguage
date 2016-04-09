package main

import "fmt"

func main() {
	x := []float64{
  		48,96,86,68,
  		57,82,63,70,
  		37,34,83,27,
  		19,97, 9,17,
	}

	var menor float64 = x[0]
	for i := 0; i < len(x); i++ {
		if x[i] < menor {
			menor = x[i]
		}
	}
	fmt.Println("El menor es ", menor)
	fmt.Println("El promedio es ", promedio(x))
}

func promedio(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}