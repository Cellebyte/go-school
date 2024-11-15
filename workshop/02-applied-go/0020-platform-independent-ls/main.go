package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to my awesome 'ls' program written in Golang")

	// TODO 1:  get current directory
	var dirPath string
	dirPath, err := os.Getwd()
	if err != nil {
		fmt.Println("cannot find current dir directory")
		os.Exit(1)
	}
	// TODO 2:  open directory and get file pointer
	var dir *os.File
	dir, err = os.Open(dirPath)
	if err != nil {
		fmt.Println("cannot open current directory")
		os.Exit(1)
	}
	// TODO 3:  get files from directory
	var files []os.FileInfo
	files, err = dir.Readdir(0)
	if err != nil {
		fmt.Println("cannot read current directory")
		os.Exit(1)
	}

	// TODO 4:  print file contents
	for _, file := range files {
		fmt.Printf("%v %d %v %s\n", file.Mode(), file.Size(), file.ModTime(), file.Name())
	}
}
