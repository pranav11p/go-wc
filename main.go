package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to wc tool written in Go")

	args := os.Args[1:]
	var res int64
	var err error

	if args[0] == "-c" {
		res, err = countBytes(args[1])
	} else {
		fmt.Println("Invalid arguments")
		return
	}
	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println(res)
	}
}

func countBytes(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return -1, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return -1, fmt.Errorf("error getting the file stat: %v", err)
	}

	totalBytes := fileStat.Size()

	return totalBytes, nil
}
