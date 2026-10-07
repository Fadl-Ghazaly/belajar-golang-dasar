package main

import "fmt"

// Definisi Struct
type Server struct {
	Hostname   string
	IPAddress  string
	Port       int
	IsActive   bool
	MaxStorage int
}

// Struct Method (Menempel pada Struct Server)
func (s Server) GetFullAddress() string {
	return fmt.Sprintf("%s:%d", s.IPAddress, s.Port)
}

func (s Server) DisplayInfo() {
	fmt.Println("----------------------------------------")
	fmt.Println("Hostname  :", s.Hostname)
	fmt.Println("IP Addr   :", s.GetFullAddress())
	fmt.Println("Status    :", s.IsActive)
	fmt.Println("Storage   :", s.MaxStorage, "GB")
}

func main() {
	fmt.Println("=== MANAJEMEN STRUCT SERVER ===")

	// Inisialisasi Struct 1: Eksplisit Field Name
	server1 := Server{
		Hostname:   "web-prod-01",
		IPAddress:  "192.168.1.10",
		Port:       8080,
		IsActive:   true,
		MaxStorage: 500,
	}

	// Inisialisasi Struct 2: Literal Nilai
	server2 := Server{"db-prod-01", "192.168.1.20", 3306, true, 2000}

	// Panggil Struct Method
	server1.DisplayInfo()
	server2.DisplayInfo()
}