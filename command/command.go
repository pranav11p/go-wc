package command

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Count struct {
	bytes int
}

type flagsStruct struct {
	c bool
}

const LineEndBytes int = 2

func ExecuteCommand() {
	flags := flagsStruct{}
	flag.BoolVar(&flags.c, "c", false, "Count the number of bytes")
	flag.Parse()

	inputStream := getInputStream()
	defer inputStream.Close()

	resCounts := Count{}

	resCounts.countBytesLinesWord(inputStream, flags)
	resCounts.printCounts(flags)
}

func getInputStream() *os.File {
	filepath := flag.Arg(0)

	// If file is not provided, then take input from stdin
	if filepath == "" {
		return os.Stdin
	}

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file: err:", err)
		os.Exit(1)
	}

	return f
}

func (resCounts *Count) countBytesLinesWord(inputStream *os.File, flags flagsStruct) {
	// bufio takes some data from file into buffer/memory instead of whole file
	scanner := bufio.NewScanner(inputStream)

	for scanner.Scan() {

		bytesArray := scanner.Bytes()
		if flags.c {
			resCounts.bytes += len(bytesArray) + LineEndBytes
		}
	}
}

func (resCount *Count) printCounts(flags flagsStruct) {
	if flags.c {
		fmt.Printf("%d", resCount.bytes)
	}
	fmt.Printf("\n")
}
