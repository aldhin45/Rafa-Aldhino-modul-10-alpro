package main

import (
	"fmt"
)

func inputBerat(n int) []float64 {
	berat := make([]float64, n)
	fmt.Println("berat anak kelinci:")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scanln(&berat[i])
	}
	return berat
}

func cariNilaiEkstrem(data []float64) (float64, float64) {
	min := data[0]
	max := data[0]

	for _, value := range data {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return min, max
}

func cetakHasil(min, max float64) {
	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}

func main() {
	var n int

	fmt.Print("jumlah anak kelinci: ")
	fmt.Scanln(&n)

	if n <= 0 {
		fmt.Println("Jumlah anak kelinci lebih besar dari 0.")
		return
	}

	berat := inputBerat(n)
	min, max := cariNilaiEkstrem(berat)

	cetakHasil(min, max)
}
