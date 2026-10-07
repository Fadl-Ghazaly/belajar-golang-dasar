package main

import "fmt"

func main() {
	// 1. IF EXPRESSION & SHORT STATEMENT
	// Memeriksa status suhu ruangan server
	serverTemp := 38 // dalam Derajat Celsius

	fmt.Println("=== 1. PENGECEKAN SUHU SERVER (IF STATEMENT) ===")
	if temp := serverTemp; temp > 40 {
		fmt.Println("CRITICAL: Suhu server sangat tinggi!", temp, "°C")
	} else if temp >= 30 {
		fmt.Println("WARNING: Suhu di atas normal:", temp, "°C (Pendingin Aktif)")
	} else {
		fmt.Println("NORMAL: Suhu ruangan stabil:", temp, "°C")
	}

	// 2. SWITCH EXPRESSION
	// Memeriksa kode status server
	statusCode := "ONLINE"

	fmt.Println("\n=== 2. STATUS SERVER (SWITCH STATEMENT) ===")
	switch statusCode {
	case "ONLINE":
		fmt.Println("Server beroperasi dengan lancar.")
	case "MAINTENANCE":
		fmt.Println("Server sedang dalam pemeliharaan berkala.")
	case "OFFLINE":
		fmt.Println("ALERT: Server mati/tidak merespons!")
	default:
		fmt.Println("Status server tidak dikenali.")
	}

	// 3. FOR LOOP STANDAR (DENGAN CONTINUE & BREAK)
	// Iterasi pengecekan port 1 sampai 5
	fmt.Println("\n=== 3. SCANNING PORT SERVER (FOR LOOP STANDAR) ===")
	for port := 1; port <= 5; port++ {
		if port == 2 {
			// Lewati port 2
			fmt.Println("Port 2: DILEWATI (Filtered)")
			continue
		}
		if port == 4 {
			// Hentikan scanning di port 4
			fmt.Println("Port 4: TERDETEKSI ANOMALI! Scanning dihentikan.")
			break
		}
		fmt.Println("Port", port, ": Terhubung & Aman")
	}

	// 4. FOR RANGE (ITERASI SLICE)
	// Memeriksa daftar layanan data center
	services := []string{"Web Server", "Database MySQL", "DNS Server", "FTP Server"}

	fmt.Println("\n=== 4. DAFTAR LAYANAN (FOR RANGE) ===")
	for index, service := range services {
		fmt.Printf("Layanan #%d: %s\n", index+1, service)
	}
}