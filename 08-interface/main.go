package main

import "fmt"

// 1. Definisi Interface (Kontrak Method)
type DeviceController interface {
	TurnOn() string
	TurnOff() string
}

// 2. Struct 1: Router
type Router struct {
	Model string
}

func (r Router) TurnOn() string {
	return fmt.Sprintf("Router %s: Booting RouterOS...", r.Model)
}

func (r Router) TurnOff() string {
	return fmt.Sprintf("Router %s: Shutting down interface...", r.Model)
}

// 3. Struct 2: Switch
type SwitchDevice struct {
	PortCount int
}

func (s SwitchDevice) TurnOn() string {
	return fmt.Sprintf("Switch %d-Port: Powering ON all ports...", s.PortCount)
}

func (s SwitchDevice) TurnOff() string {
	return fmt.Sprintf("Switch %d-Port: Powering OFF ports...", s.PortCount)
}

// Function yang menerima parameter Interface
func ExecuteDeviceAction(device DeviceController) {
	fmt.Println(device.TurnOn())
	fmt.Println(device.TurnOff())
}

// Function dengan Interface Kosong (any)
func LogAnyData(data any) {
	fmt.Printf("[LOG DYNAMIC DATA]: %v (Tipe: %T)\n", data, data)
}

func main() {
	fmt.Println("=== DEMO INTERFACE & INTERFACE KOSONG ===")

	ciscoRouter := Router{Model: "CCR1009"}
	tpLinkSwitch := SwitchDevice{PortCount: 24}

	// Mengirim struct yang mengimplementasikan interface
	ExecuteDeviceAction(ciscoRouter)
	ExecuteDeviceAction(tpLinkSwitch)

	fmt.Println("----------------------------------------")
	// Menggunakan interface kosong (any) untuk menampung tipe data apa saja
	LogAnyData("Sistem Berjalan")
	LogAnyData(2026)
	LogAnyData(true)
}