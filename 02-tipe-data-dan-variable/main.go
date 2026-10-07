package main

import "fmt"

func main() {
	// 1. Constant (Nilai tetap)
	const storeName = "TechStore Lampung"
	const taxRate = 0.11 // PPN 11%

	// 2. Variable dengan Type Inference (:=)
	productName := "Monitor Gaming 144Hz"
	basePrice := 2500000.0
	initialStock := 15

	// 3. Variable dengan deklarasi eksplisit
	var itemsSold int = 3
	var isAvailable bool = initialStock > 0

	// 4. Operasi Matematika & Augmented Assignment
	remainingStock := initialStock - itemsSold
	totalSales := basePrice * float64(itemsSold)
	totalTax := totalSales * taxRate
	grandTotal := totalSales + totalTax

	// 5. Output Informasi
	fmt.Println("=== SISTEM KASIR", storeName, "===")
	fmt.Println("Nama Produk    :", productName)
	fmt.Println("Status Produk  : Tersedia? ->", isAvailable)
	fmt.Println("Sisa Stok      :", remainingStock, "unit")
	fmt.Println("Total Jual     : Rp", totalSales)
	fmt.Println("PPN (11%)      : Rp", totalTax)
	fmt.Println("Grand Total    : Rp", grandTotal)
}