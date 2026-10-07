package main

import (
	"errors"
	"fmt"
)

// Custom Error Struct
type ValidationError struct {
	Field   string
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("Validation Error pada '%s': %s", v.Field, v.Message)
}

// Function yang mengembalikan standar error
func divideServerLoad(totalLoad int, activeServers int) (int, error) {
	if activeServers == 0 {
		return 0, errors.New("jumlah server aktif tidak boleh nol")
	}
	return totalLoad / activeServers, nil}

// Function yang mendemonstrasikan Defer, Panic, dan Recover
func executeSystemInitialization(isConfigValid bool) {
	// Defer akan selalu dieksekusi di akhir fungsi, bahkan jika terjadi Panic
	defer func() {
		fmt.Println("[DEFER]: Pembersihan memori dan penutupan log selesai.")
		// Recover menangkap Panic agar program tidak crash total
		if r := recover(); r != nil {
			fmt.Println("[RECOVER]: Sistem berhasil pulih dari Panic! Pesan:", r)
		}
	}()

	fmt.Println("[PROCESS]: Memulai inisialisasi sistem...")
	if !isConfigValid {
		// Memicu Panic secara sengaja
		panic("CRITICAL: File konfigurasi tidak ditemukan/rusak!")
	}
	fmt.Println("[PROCESS]: Inisialisasi sistem sukses!")
}

func main() {
	fmt.Println("=== DEMO ERROR HANDLING ===")

	// 1. Penanganan Standar Error
	result, err := divideServerLoad(100, 0)
	if err != nil {
		fmt.Println("Gagal membagi beban:", err)
	} else {
		fmt.Println("Beban per server:", result)
	}

	fmt.Println("----------------------------------------")

	// 2. Custom Error Checking
	customErr := &ValidationError{Field: "IPAddress", Message: "Format IP tidak valid"}
	fmt.Println("Custom Error Output:", customErr.Error())

	fmt.Println("----------------------------------------")

	// 3. Eksekusi Defer, Panic, & Recover
	executeSystemInitialization(false) // Trigger panic

	fmt.Println("\n[INFO]: Program utama tetap berjalan lancar setelah ditangkap oleh Recover.")
}