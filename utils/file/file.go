package file

import (
	"fmt"
	"os"
)

// 读取文件到[]byte中
func FileName2Bytes(filename string) ([]byte, error) {
	// File
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// FileInfo:
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// []byte
	data := make([]byte, stats.Size())
	count, err := file.Read(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("read file %s len: %d \n", filename, count)
	return data, nil
}
func File2Bytes(file *os.File) ([]byte, error) {
	// File
	// FileInfo:
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, stats.Size())
	file.Read(data)
	return data, nil
}
