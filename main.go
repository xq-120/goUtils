package main

import (
	"fmt"

	goUtils "github.com/xq-120/goUtils/pkg"
)

func main() {
	filepath := "example.txt"
	ext := goUtils.GetFileExtension(filepath)
	fmt.Println("File extension:", ext)
}
