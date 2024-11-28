package main

import (
	"fmt"
)

func inputData(x, y int) [][]float64 {
	data := make([][]float64, x)

	fmt.Println("berat ikan:")
	for i := 0; i < x; i++ {
		data[i] = make([]float64, y)
		fmt.Printf("Wadah ke-%d:\n", i+1)
		for j := 0; j < y; j++ {
			fmt.Printf("  Berat ikan ke-%d: ", j+1)
			fmt.Scanln(&data[i][j])
		}
	}

	return data
}

func hitungTotalPerWadah(data [][]float64) []float64 {
	totals := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		for _, berat := range data[i] {
			totals[i] += berat
		}
	}
	return totals
}

func hitungRataRataPerWadah(data [][]float64) []float64 {
	rataRata := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		total := 0.0
		for _, berat := range data[i] {
			total += berat
		}
		rataRata[i] = total / float64(len(data[i]))
	}
	return rataRata
}

func cariNilaiEkstrem(totals []float64) (float64, float64) {
	min := totals[0]
	max := totals[0]

	for _, total := range totals {
		if total < min {
			min = total
		}
		if total > max {
			max = total
		}
	}

	return min, max
}

func cetakHasil(totals, rataRata []float64, min, max float64) {
	fmt.Println("\nHasil:")
	fmt.Println("Total berat :")
	for i, total := range totals {
		fmt.Printf("  Wadah ke-%d: %.2f\n", i+1, total)
	}
	fmt.Println("\nRata-rata berat :")
	for i, rata := range rataRata {
		fmt.Printf("  Wadah ke-%d: %.2f\n", i+1, rata)
	}
	fmt.Printf("\nBerat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}

func main() {
	var x, y int

	fmt.Print("jumlah wadah (x): ")
	fmt.Scanln(&x)
	fmt.Print("jumlah ikan per wadah (y): ")
	fmt.Scanln(&y)

	if x <= 0 || y <= 0 {
		fmt.Println("Jumlah wadah dan ikan per wadah lebih besar dari 0.")
		return
	}

	data := inputData(x, y)
	totals := hitungTotalPerWadah(data)
	rataRata := hitungRataRataPerWadah(data)
	min, max := cariNilaiEkstrem(totals)

	cetakHasil(totals, rataRata, min, max)
}
