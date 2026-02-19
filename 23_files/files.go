package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	paths := []string{
		"example.txt",
		filepath.Join("23_files", "example.txt"),
	}

	var f *os.File
	var err error
	for _, path := range paths {
		f, err = os.Open(path)
		if err == nil {
			break
		}
	}

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	fmt.Println("fileName:", fileInfo.Name())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Println("Size:", fileInfo.Size()) // in bytes
	fmt.Println("file permission:", fileInfo.Mode()) // file permission in octal format (linux file permission)
	fmt.Println("modified time:", fileInfo.ModTime()) // modified time of the file


	example1, err := os.Open("23_files/example1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer example1.Close()
	example1Info, err := example1.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	buffer := make([]byte, example1Info.Size())

	// read the content of the file and store it in the buffer
	_, err = example1.Read(buffer) // here we pass the buffer to the Read method and it will read the content of the file and store it in the buffer
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("file content:", buffer)
}
