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
	// _ is buffer length
	_, err = example1.Read(buffer) // here we pass the buffer to the Read method and it will read the content of the file and store it in the buffer
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("file content:", string(buffer))
	// we can also iterate over the buffer and print each byte and its corresponding character
	for i:=0; i<len(buffer);i++ {
		fmt.Println("byte:", buffer[i], "character:", string(buffer[i]))
	}


	// we can use the os.ReadFile function to read the content of the file in one line of code
	// it is not recommended to use os.ReadFile for large files as it will read the entire file into memory and can cause out of memory error
	content, err := os.ReadFile("23_files/example1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("file content using os.ReadFile:", string(content))

	dir, err := os.Open(".")
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer dir.Close()
	fileInfo2, err := dir.ReadDir(40) // read the content of the directory and store it in a slice of os.DirEntry
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range fileInfo2 {
		fmt.Println("file name:", file.Name())
		fmt.Println("is directory:", file.IsDir())
	}
	
}
