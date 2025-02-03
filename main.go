package main

import (
	"fmt"

	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

func main() {
	filePath := "example_file.txt"

	// Check if the file exists
	if gfile.Exists(filePath) {
		glog.Info("File exists: ", filePath)
	} else {
		glog.Info("File does not exist, creating: ", filePath)
		initialContent := "This file was created using the GoFrame gfile package.\n"
		if err := gfile.PutContents(filePath, initialContent); err != nil {
			glog.Error("Error creating file: ", err)
			return
		}
	}

	// Append content to the file
	appendContent := "Appending a line: Performing file operations in Golang.\n"
	if err := gfile.PutContentsAppend(filePath, appendContent); err != nil {
		glog.Error("Error appending to file: ", err)
		return
	}

	// Read file content
	content, err := gfile.GetContents(filePath)
	if err != nil {
		glog.Error("Error reading file: ", err)
		return
	}
	fmt.Println("File content:")
	fmt.Println(content)

	// Get file size in bytes
	size, err := gfile.Size(filePath)
	if err != nil {
		glog.Error("Error getting file size: ", err)
		return
	}
	fmt.Printf("File size: %d bytes\n", size)

	// Retrieve directory and base name from the file path
	directory := gfile.Dir(filePath)
	baseName := gfile.Basename(filePath)
	fmt.Println("File directory:", directory)
	fmt.Println("File name:", baseName)
}
