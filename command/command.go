package command

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Count struct {
	lines int
	bytes int
}

type flagsStruct struct {
	l bool // Count lines
	c bool // Count bytes
}

const LineEndBytes int = 2

func ExecuteCommand() {

	flags := getFlags()

	inputStream := getInputStream()
	defer inputStream.Close()

	resCounts := Count{}

	resCounts.countBytesLinesWord(inputStream, flags)
	resCounts.printCounts(flags)
}

func getFlags() flagsStruct {
	flags := flagsStruct{}

	flag.BoolVar(&flags.c, "c", false, "Count the number of bytes")
	flag.BoolVar(&flags.l, "l", false, "Count the number of lines")

	flag.Parse()
	return flags
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

		if flags.l {
			resCounts.lines++
		}
		if flags.c {
			resCounts.bytes += len(bytesArray) + LineEndBytes
		}
	}
}

func (resCount *Count) printCounts(flags flagsStruct) {
	resArray := []int{}
	if flags.l {
		resArray = append(resArray, resCount.lines)
	}
	if flags.c {
		resArray = append(resArray, resCount.bytes)
	}

	for idx, val := range resArray {
		if idx > 0 {
			fmt.Printf("  ")
		}
		fmt.Printf("%d", val)
	}
	fmt.Printf("\n")
}
