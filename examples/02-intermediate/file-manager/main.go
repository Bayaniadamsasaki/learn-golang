package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	currentDir, _ := os.Getwd()
	
	for {
		showMenu(currentDir)
		choice := getChoice()
		
		switch choice {
		case 1:
			listFiles(currentDir)
		case 2:
			createDirectory(currentDir)
		case 3:
			copyFile(currentDir)
		case 4:
			moveFile(currentDir)
		case 5:
			deleteFile(currentDir)
		case 6:
			showFileInfo(currentDir)
		case 7:
			searchFiles(currentDir)
		case 8:
			currentDir = changeDirectory(currentDir)
		case 9:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option!")
		}
		
		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln()
	}
}

func showMenu(currentDir string) {
	fmt.Println("\n=== FILE MANAGER ===")
	fmt.Printf("Current directory: %s\n", currentDir)
	fmt.Println("1. List Files")
	fmt.Println("2. Create Directory")
	fmt.Println("3. Copy File")
	fmt.Println("4. Move File")
	fmt.Println("5. Delete File")
	fmt.Println("6. File Info")
	fmt.Println("7. Search Files")
	fmt.Println("8. Change Directory")
	fmt.Println("9. Exit")
}

func getChoice() int {
	var choice int
	fmt.Print("Choose option (1-9): ")
	fmt.Scanln(&choice)
	return choice
}

func listFiles(dir string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}
	
	fmt.Printf("\nContents of %s:\n", dir)
	fmt.Println("Type\tSize\t\tName")
	fmt.Println("----\t----\t\t----")
	
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		
		fileType := "FILE"
		if entry.IsDir() {
			fileType = "DIR"
		}
		
		fmt.Printf("%s\t%d bytes\t%s\n", fileType, info.Size(), entry.Name())
	}
}

func createDirectory(currentDir string) {
	var dirName string
	fmt.Print("Enter directory name: ")
	fmt.Scanln(&dirName)
	
	fullPath := filepath.Join(currentDir, dirName)
	err := os.MkdirAll(fullPath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	
	fmt.Printf("Directory '%s' created successfully\n", dirName)
}

func copyFile(currentDir string) {
	var source, dest string
	fmt.Print("Enter source file: ")
	fmt.Scanln(&source)
	fmt.Print("Enter destination: ")
	fmt.Scanln(&dest)
	
	sourcePath := filepath.Join(currentDir, source)
	destPath := filepath.Join(currentDir, dest)
	
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		fmt.Printf("Error opening source file: %v\n", err)
		return
	}
	defer sourceFile.Close()
	
	destFile, err := os.Create(destPath)
	if err != nil {
		fmt.Printf("Error creating destination file: %v\n", err)
		return
	}
	defer destFile.Close()
	
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Printf("Error copying file: %v\n", err)
		return
	}
	
	fmt.Printf("File copied from '%s' to '%s'\n", source, dest)
}

func moveFile(currentDir string) {
	var source, dest string
	fmt.Print("Enter source file: ")
	fmt.Scanln(&source)
	fmt.Print("Enter destination: ")
	fmt.Scanln(&dest)
	
	sourcePath := filepath.Join(currentDir, source)
	destPath := filepath.Join(currentDir, dest)
	
	err := os.Rename(sourcePath, destPath)
	if err != nil {
		fmt.Printf("Error moving file: %v\n", err)
		return
	}
	
	fmt.Printf("File moved from '%s' to '%s'\n", source, dest)
}

func deleteFile(currentDir string) {
	var filename string
	fmt.Print("Enter file/directory to delete: ")
	fmt.Scanln(&filename)
	
	fullPath := filepath.Join(currentDir, filename)
	
	info, err := os.Stat(fullPath)
	if err != nil {
		fmt.Printf("File not found: %v\n", err)
		return
	}
	
	if info.IsDir() {
		err = os.RemoveAll(fullPath)
	} else {
		err = os.Remove(fullPath)
	}
	
	if err != nil {
		fmt.Printf("Error deleting: %v\n", err)
		return
	}
	
	fmt.Printf("'%s' deleted successfully\n", filename)
}

func showFileInfo(currentDir string) {
	var filename string
	fmt.Print("Enter filename: ")
	fmt.Scanln(&filename)
	
	fullPath := filepath.Join(currentDir, filename)
	info, err := os.Stat(fullPath)
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}
	
	fmt.Printf("\nFile Information:\n")
	fmt.Printf("Name: %s\n", info.Name())
	fmt.Printf("Size: %d bytes\n", info.Size())
	fmt.Printf("Mode: %s\n", info.Mode())
	fmt.Printf("Modified: %s\n", info.ModTime().Format(time.RFC3339))
	fmt.Printf("Is Directory: %t\n", info.IsDir())
}

func searchFiles(currentDir string) {
	var pattern string
	fmt.Print("Enter search pattern: ")
	fmt.Scanln(&pattern)
	
	var matches []string
	
	err := filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		
		if strings.Contains(strings.ToLower(info.Name()), strings.ToLower(pattern)) {
			relPath, _ := filepath.Rel(currentDir, path)
			matches = append(matches, relPath)
		}
		
		return nil
	})
	
	if err != nil {
		fmt.Printf("Error searching: %v\n", err)
		return
	}
	
	if len(matches) == 0 {
		fmt.Println("No files found matching the pattern")
		return
	}
	
	fmt.Printf("\nFound %d matches:\n", len(matches))
	for _, match := range matches {
		fmt.Printf("- %s\n", match)
	}
}

func changeDirectory(currentDir string) string {
	var newDir string
	fmt.Print("Enter new directory (or .. for parent): ")
	fmt.Scanln(&newDir)
	
	var targetDir string
	if newDir == ".." {
		targetDir = filepath.Dir(currentDir)
	} else if filepath.IsAbs(newDir) {
		targetDir = newDir
	} else {
		targetDir = filepath.Join(currentDir, newDir)
	}
	
	_, err := os.Stat(targetDir)
	if err != nil {
		fmt.Printf("Directory not found: %v\n", err)
		return currentDir
	}
	
	fmt.Printf("Changed to directory: %s\n", targetDir)
	return targetDir
}
