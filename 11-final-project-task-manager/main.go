package main

import (
	"errors"
	"fmt"
)

// Task Struct
type Task struct {
	ID        int
	Title     string
	Category  string
	IsDone    bool
}

// TaskManager Struct
type TaskManager struct {
	Tasks []Task
}

// Method dengan Pointer Receiver untuk menambah Task
func (tm *TaskManager) AddTask(title string, category string) {
	newTask := Task{
		ID:       len(tm.Tasks) + 1,
		Title:    title,
		Category: category,
		IsDone:   false,
	}
	tm.Tasks = append(tm.Tasks, newTask)
	fmt.Printf("✅ Tugas '%s' berhasil ditambahkan!\n", title)
}

// Method untuk menandai Task selesai
func (tm *TaskManager) CompleteTask(id int) error {
	for i := range tm.Tasks {
		if tm.Tasks[i].ID == id {
			if tm.Tasks[i].IsDone {
				return errors.New("tugas ini sudah ditandai selesai sebelumnya")
			}
			tm.Tasks[i].IsDone = true
			fmt.Printf("🎉 Tugas ID #%d berhasil diselesaikan!\n", id)
			return nil
		}
	}
	return fmt.Errorf("tugas dengan ID #%d tidak ditemukan", id)
}

// Method untuk menampilkan seluruh Task
func (tm TaskManager) ListTasks() {
	fmt.Println("\n==========================================")
	fmt.Println("        DAFTAR TUGAS DAFTAR KERJA         ")
	fmt.Println("==========================================")

	if len(tm.Tasks) == 0 {
		fmt.Println("Belum ada tugas tersimpan.")
		return
	}

	for _, task := range tm.Tasks {
		status := "❌ Belum Selesai"
		if task.IsDone {
			status = "✅ Selesai"
		}
		fmt.Printf("[%d] %s (%s) -> Status: %s\n", task.ID, task.Title, task.Category, status)
	}
	fmt.Println("------------------------------------------")
}

func main() {
	defer func() {
		fmt.Println("\n[INFO]: Aplikasi Task Manager ditutup secara aman.")
	}()

	manager := TaskManager{}

	// Menambahkan beberapa Tugas
	manager.AddTask("Instalasi Server Go", "Infrastructure")
	manager.AddTask("Konfigurasi Database MySQL", "Backend")
	manager.AddTask("Membuat Video Remake Golang", "Education")

	// Tampilkan Daftar Tugas Awal
	manager.ListTasks()

	// Selesaikan Salah Satu Tugas
	err := manager.CompleteTask(1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Coba Selesaikan Tugas yang Tidak Ada (Uji Error Handling)
	errNotExists := manager.CompleteTask(99)
	if errNotExists != nil {
		fmt.Println("Error Handled:", errNotExists)
	}

	// Tampilkan Kembali Daftar Tugas Akhir
	manager.ListTasks()
}