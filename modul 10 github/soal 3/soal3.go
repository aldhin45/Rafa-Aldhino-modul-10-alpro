package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func hitungRerata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var arrBerat arrBalita
	var bMin, bMax float64

	fmt.Print("banyak data berat balita: ")
	fmt.Scanln(&n)

	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data balita")
		return
	}

	fmt.Println("berat balita (kg):")
	for i := 0; i < n; i++ {
		fmt.Printf("berat balita ke-%d: ", i+1)
		fmt.Scanln(&arrBerat[i])
	}

	hitungMinMax(arrBerat, n, &bMin, &bMax)
	rerata := hitungRerata(arrBerat, n)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata)
}
