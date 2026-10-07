package main

import "fmt"

type DatabaseConfig struct {
	Host     string
	Port     int
	MaxConn  int
}

// Pass by Value (Tidak mengubah data asli)
func updatePortByValue(config DatabaseConfig, newPort int) {
	config.Port = newPort
}

// Pass by Reference menggunakan Pointer (*)
func updatePortByPointer(config *DatabaseConfig, newPort int) {
	config.Port = newPort
}

// Struct Method dengan Pointer Receiver
func (db *DatabaseConfig) IncreaseConnection(amount int) {
	db.MaxConn += amount
}

func main() {
	fmt.Println("=== DEMO POINTER GOLANG ===")

	// Deklarasi awal
	dbConfig := DatabaseConfig{
		Host:    "localhost",
		Port:    5432,
		MaxConn: 50,
	}

	fmt.Println("1. Kondisi Awal Config        :", dbConfig)

	// Uji Pass by Value
	updatePortByValue(dbConfig, 5433)
	fmt.Println("2. Setelah updatePortByValue  :", dbConfig, "(Port Tetap 5432)")

	// Uji Pass by Reference menggunakan Operator & (Mengambil Alamat Memori)
	updatePortByPointer(&dbConfig, 5433)
	fmt.Println("3. Setelah updatePortByPointer:", dbConfig, "(Port Berubah ke 5433)")

	// Pointer Receiver pada Struct Method
	dbConfig.IncreaseConnection(25)
	fmt.Println("4. Setelah IncreaseConnection :", dbConfig, "(MaxConn Berubah ke 75)")

	// Menggunakan Operator Asterisk (*) Dereference
	pointerToDB := &dbConfig
	fmt.Println("Alamat Memori Config          :", pointerToDB)
	fmt.Println("Nilai di Alamat Memori (*)    :", *pointerToDB)
}