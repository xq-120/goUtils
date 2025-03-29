package goutils

import (
	"fmt"
	"os"
	"strings"
)

// FormatBytes 将字节数转换为可读格式 (KB, MB, GB)
func FormatBytes(bytes int64) string {
	const (
		kb = 1024
		mb = kb * 1024
		gb = mb * 1024
	)

	switch {
	case bytes >= gb:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(gb))
	case bytes >= mb:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(mb))
	case bytes >= kb:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(kb))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}

func GetFileExtension(filename string) string {
	if len(strings.Split(filename, ".")) == 2 && filename[0] == '.' { //隐藏文件
		return ""
	}
	index := strings.LastIndex(filename, ".")
	if index == -1 {
		return ""
	}
	ext := filename[index+1:]
	return ext
}

func IsHiddenFile(filename string) bool {
	/**
	无法判断"chflags hidden 文件夹"的情况。
	chflags hidden 文件夹
	chflags nohidden 文件夹
	*/
	if len(filename) > 0 && filename[0] == '.' {
		return true
	}
	return false
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	// 不能使用os.IsExist(err) 判断文件是否存在
	// os.Stat(path) 如果返回的是 权限错误，os.IsExist(err) 仍然返回 false，这会误判为“文件不存在”。
	// os.IsExist(err) 主要用于创建文件或目录时检查是否已存在，而不是检查文件是否已经存在。
	// 但是 os.IsNotExist(err) 是可以判断文件不存在的。
	return !os.IsNotExist(err)
}
