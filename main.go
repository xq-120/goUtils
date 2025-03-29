package main

import (
	"fmt"

	"github.com/xq-120/goUtils/pkg/file"
)

func main() {
	filepath := "example.txt"
	ext := file.GetFileExtension(filepath)
	fmt.Println("File extension:", ext)
}
