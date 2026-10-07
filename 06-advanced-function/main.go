package main

import "fmt"

// 1. Variadic Function (Menerima jumlah parameter dinamis/slice)
func calculateTotalCost(costs ...float64) float64 {
	total := 0.0
	for _, cost := range costs {
		total += cost
	}
	return total
}

// 2. Function as Parameter & Type Declaration
type FilterFunc func(string) string

func processServerName(name string, filter FilterFunc) string {
	return filter(name)
}

func uppercaseFilter(name string) string {
	return "SERVER-" + name
}

func main() {
	fmt.Println("=== ADVANCED FUNCTION DEMO ===")

	// Penggunaan Variadic Function
	total1 := calculateTotalCost(150.0, 200.5, 300.0)
	fmt.Println("Total Biaya Operasional 1 : Rp", total1)

	// Kirim data Slice ke Variadic Function menggunakan operator ...
	costsSlice := []float64{100.0, 50.0, 75.0, 225.0}
	total2 := calculateTotalCost(costsSlice...)
	fmt.Println("Total Biaya Operasional 2 : Rp", total2)

	// Penggunaan Function as Parameter
	formattedName := processServerName("database-primary", uppercaseFilter)
	fmt.Println("Nama Server Terformat     :", formattedName)

	// Penggunaan Anonymous Function (Fungsi tanpa nama)
	isCriticalAlert := func(cpuUsage float64) bool {
		return cpuUsage > 85.0
	}
	fmt.Println("Peringatan CPU 90%        :", isCriticalAlert(90.0))

	// Penggunaan Closure (Fungsi yang mengakses variabel luar)
	counter := 0
	incrementCounter := func() {
		counter++ // Mengakses dan mengubah variabel counter di luar scope
	}

	incrementCounter()
	incrementCounter()
	fmt.Println("Total Log Terdeteksi      :", counter)
}