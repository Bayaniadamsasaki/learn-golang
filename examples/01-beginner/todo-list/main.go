package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TodoList struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

func main() {
	todoList := loadTodoList()
	
	for {
		showMenu()
		choice := getChoice()
		
		switch choice {
		case 1:
			addTask(&todoList)
		case 2:
			showTasks(todoList)
		case 3:
			completeTask(&todoList)
		case 4:
			deleteTask(&todoList)
		case 5:
			saveTodoList(todoList)
			fmt.Println("Data disimpan. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
		
		fmt.Println()
	}
}

func showMenu() {
	fmt.Println("=== TODO LIST ===")
	fmt.Println("1. Tambah Tugas")
	fmt.Println("2. Lihat Tugas")
	fmt.Println("3. Selesaikan Tugas")
	fmt.Println("4. Hapus Tugas")
	fmt.Println("5. Keluar")
}

func getChoice() int {
	var choice int
	fmt.Print("Pilih menu (1-5): ")
	fmt.Scanln(&choice)
	return choice
}

func addTask(todoList *TodoList) {
	var title string
	fmt.Print("Masukkan tugas: ")
	fmt.Scanln(&title)
	
	task := Task{
		ID:    todoList.NextID,
		Title: title,
		Done:  false,
	}
	
	todoList.Tasks = append(todoList.Tasks, task)
	todoList.NextID++
	
	fmt.Println("Tugas berhasil ditambahkan!")
}

func showTasks(todoList TodoList) {
	if len(todoList.Tasks) == 0 {
		fmt.Println("Tidak ada tugas.")
		return
	}
	
	fmt.Println("\n=== DAFTAR TUGAS ===")
	for _, task := range todoList.Tasks {
		status := "[ ]"
		if task.Done {
			status = "[✓]"
		}
		fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
	}
}

func completeTask(todoList *TodoList) {
	showTasks(*todoList)
	if len(todoList.Tasks) == 0 {
		return
	}
	
	var id int
	fmt.Print("Masukkan ID tugas yang selesai: ")
	fmt.Scanln(&id)
	
	for i := range todoList.Tasks {
		if todoList.Tasks[i].ID == id {
			todoList.Tasks[i].Done = true
			fmt.Println("Tugas berhasil diselesaikan!")
			return
		}
	}
	
	fmt.Println("ID tugas tidak ditemukan!")
}

func deleteTask(todoList *TodoList) {
	showTasks(*todoList)
	if len(todoList.Tasks) == 0 {
		return
	}
	
	var id int
	fmt.Print("Masukkan ID tugas yang akan dihapus: ")
	fmt.Scanln(&id)
	
	for i, task := range todoList.Tasks {
		if task.ID == id {
			todoList.Tasks = append(todoList.Tasks[:i], todoList.Tasks[i+1:]...)
			fmt.Println("Tugas berhasil dihapus!")
			return
		}
	}
	
	fmt.Println("ID tugas tidak ditemukan!")
}

func loadTodoList() TodoList {
	file, err := os.ReadFile("todos.json")
	if err != nil {
		return TodoList{Tasks: []Task{}, NextID: 1}
	}
	
	var todoList TodoList
	err = json.Unmarshal(file, &todoList)
	if err != nil {
		return TodoList{Tasks: []Task{}, NextID: 1}
	}
	
	return todoList
}

func saveTodoList(todoList TodoList) {
	data, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		fmt.Println("Error menyimpan data:", err)
		return
	}
	
	err = os.WriteFile("todos.json", data, 0644)
	if err != nil {
		fmt.Println("Error menulis file:", err)
	}
}
