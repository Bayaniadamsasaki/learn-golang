// File I/O: membaca dan menulis file dengan berbagai metode
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("=== BASIC FILE OPERATIONS ===")
	
	filename := "example.txt"
	content := "Hello, World!\nThis is a test file.\nGolang file I/O example."
	
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	fmt.Printf("File '%s' created successfully\n", filename)
	
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("File content:\n%s\n", string(data))
	
	fmt.Println("\n=== FILE INFO ===")
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}
	
	fmt.Printf("File name: %s\n", info.Name())
	fmt.Printf("File size: %d bytes\n", info.Size())
	fmt.Printf("File mode: %s\n", info.Mode())
	fmt.Printf("Modified time: %s\n", info.ModTime())
	fmt.Printf("Is directory: %t\n", info.IsDir())
	
	fmt.Println("\n=== BUFFERED FILE OPERATIONS ===")
	file, err := os.Create("buffered.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()
	
	writer := bufio.NewWriter(file)
	for i := 1; i <= 5; i++ {
		_, err := writer.WriteString(fmt.Sprintf("Line %d\n", i))
		if err != nil {
			fmt.Printf("Error writing to buffer: %v\n", err)
			return
		}
	}
	writer.Flush()
	
	file.Close()
	
	file, err = os.Open("buffered.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("Line %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
	
	fmt.Println("\n=== DIRECTORY OPERATIONS ===")
	dirName := "test_directory"
	
	err = os.Mkdir(dirName, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	fmt.Printf("Directory '%s' created\n", dirName)
	
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}
	
	fmt.Println("Directory contents:")
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("  [DIR]  %s\n", entry.Name())
		} else {
			fmt.Printf("  [FILE] %s\n", entry.Name())
		}
	}
	
	fmt.Println("\n=== FILE PATH OPERATIONS ===")
	fullPath := filepath.Join("test_directory", "subdir", "file.txt")
	fmt.Printf("Full path: %s\n", fullPath)
	fmt.Printf("Directory: %s\n", filepath.Dir(fullPath))
	fmt.Printf("Filename: %s\n", filepath.Base(fullPath))
	fmt.Printf("Extension: %s\n", filepath.Ext(fullPath))
	
	abs, err := filepath.Abs(filename)
	if err == nil {
		fmt.Printf("Absolute path: %s\n", abs)
	}
	
	fmt.Println("\n=== COPY FILE ===")
	sourceFile := "example.txt"
	destFile := "copy_example.txt"
	
	source, err := os.Open(sourceFile)
	if err != nil {
		fmt.Printf("Error opening source file: %v\n", err)
		return
	}
	defer source.Close()
	
	destination, err := os.Create(destFile)
	if err != nil {
		fmt.Printf("Error creating destination file: %v\n", err)
		return
	}
	defer destination.Close()
	
	bytesWritten, err := io.Copy(destination, source)
	if err != nil {
		fmt.Printf("Error copying file: %v\n", err)
		return
	}
	fmt.Printf("Copied %d bytes from %s to %s\n", bytesWritten, sourceFile, destFile)
	
	fmt.Println("\n=== CLEANUP ===")
	files := []string{filename, "buffered.txt", destFile}
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			fmt.Printf("Error removing %s: %v\n", file, err)
		} else {
			fmt.Printf("Removed file: %s\n", file)
		}
	}
	
	err = os.Remove(dirName)
	if err != nil {
		fmt.Printf("Error removing directory %s: %v\n", dirName, err)
	} else {
		fmt.Printf("Removed directory: %s\n", dirName)
	}
}
