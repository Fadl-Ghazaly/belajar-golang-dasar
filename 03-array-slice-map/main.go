package main

import "fmt"

func main() {
	// 1. ARRAY (Kapasitas tetap 3)
	// Digunakan untuk daftar kategori lokasi rack server
	var serverRacks [3]string
	serverRacks[0] = "Rack-Alpha (Zone A)"
	serverRacks[1] = "Rack-Beta (Zone B)"
	serverRacks[2] = "Rack-Gamma (Zone C)"

	// 2. SLICE (Kapasitas dinamis)
	// Digunakan untuk daftar perangkat yang bisa bertambah
	serverDevices := []string{"Router Core", "Switch Managed 24-Port"}
	
	// Menambahkan elemen baru ke Slice menggunakan append()
	serverDevices = append(serverDevices, "UPS Online 3000VA", "Server Dell PowerEdge")

	// 3. MAP (Key-Value)
	// Digunakan untuk mencatat harga sewaktu-waktu berdasarkan nama perangkat
	devicePrices := map[string]int{
		"Router Core":            15000000,
		"Switch Managed 24-Port":  8500000,
		"UPS Online 3000VA":      12000000,
	}

	// Menambahkan data baru ke Map
	devicePrices["Server Dell PowerEdge"] = 45000000

	// Menghapus salah satu data dari Map
	delete(devicePrices, "Switch Managed 24-Port")

	// 4. Output Informasi
	fmt.Println("=== SISTEM INVENTARIS DATA CENTER ===")
	fmt.Println("Daftar Rack (Array)      :", serverRacks)
	fmt.Println("Total Rack (len)         :", len(serverRacks))
	fmt.Println("-------------------------------------")
	fmt.Println("Daftar Perangkat (Slice) :", serverDevices)
	fmt.Println("Jumlah Perangkat (len)   :", len(serverDevices))
	fmt.Println("Kapasitas Slice (cap)    :", cap(serverDevices))
	fmt.Println("-------------------------------------")
	fmt.Println("Daftar Harga (Map)       :", devicePrices)
	fmt.Println("Harga Server Dell        : Rp", devicePrices["Server Dell PowerEdge"])
}