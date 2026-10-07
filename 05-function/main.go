package main

import "fmt"

// 1. Function biasa tanpa parameter dan return value
func printHeader() {
	fmt.Println("=== SISTEM KALKULATOR DATA CENTER ===")
}

// 2. Function dengan parameter
func logActivity(activity string) {
	fmt.Println("[LOG]:", activity)
}

// 3. Function dengan return value
func calculateServerPower(serverCount int, wattPerServer float64) float64 {
	return float64(serverCount) * wattPerServer
}

// 4. Function dengan Multiple Return Values
func calculateStorageRequirements(users int, gbPerUser float64) (float64, float64) {
	totalGB := float64(users) * gbPerUser
	totalTB := totalGB / 1024.0
	return totalGB, totalTB
}

// 5. Function dengan Named Return Values
func getNetworkBandwidth(usageGbps float64) (bandwidthRequirement float64, isSafe bool) {
	bandwidthRequirement = usageGbps * 1.5 // Cadangan 50%
	isSafe = bandwidthRequirement <= 100.0 // Maksimal kapasitas 100 Gbps
	return
}

func main() {
	printHeader()
	logActivity("Mulai perhitungan kalkulasi sumber daya")

	// Panggil function dengan return value
	totalWatt := calculateServerPower(10, 450.5)
	fmt.Printf("Total Daya Listrik Server : %.2f Watt\n", totalWatt)

	// Panggil function dengan Multiple Return Values
	gb, tb := calculateStorageRequirements(500, 50.0)
	fmt.Printf("Kapasitas Storage         : %.0f GB (%.2f TB)\n", gb, tb)

	// Abaikan salah satu return value menggunakan blank identifier (_)
	_, tbOnly := calculateStorageRequirements(1000, 100.0)
	fmt.Printf("Total Storage (TB Only)   : %.2f TB\n", tbOnly)

	// Panggil function dengan Named Return Values
	req, safe := getNetworkBandwidth(45.0)
	fmt.Printf("Kebutuhan Bandwidth       : %.2f Gbps (Aman? %t)\n", req, safe)
}