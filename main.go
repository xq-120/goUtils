package main

import (
	"fmt"

	goUtils "github.com/xq-120/goUtils/pkg/file" // 如果你使用的是模块化的方式
)

func main() {
	filepath := "example.txt"
	ext := goUtils.GetFileExtension(filepath)
	fmt.Println("File extension:", ext)
}
