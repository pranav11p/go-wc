package command

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type Count struct {
	lines      int
	words      int
	characters int
	bytes      int
}

type flagsStruct struct {
	l bool // Count lines
	w bool // Count words
	m bool // Count characters
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

	flag.BoolVar(&flags.l, "l", false, "Count the number of lines")
	flag.BoolVar(&flags.w, "w", false, "Count the number of words")
	flag.BoolVar(&flags.c, "c", false, "Count the number of bytes")
	flag.BoolVar(&flags.m, "m", false, "Count the number of characters")

	flag.Parse()

	// If no flags are specified, the set flags -c -l -w as true
	if !(flags.l || flags.w || flags.c || flags.m) {
		flags.c = true
		flags.l = true
		flags.w = true
	}

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

		line := scanner.Text()
		bytesArray := scanner.Bytes()

		if flags.l {
			resCounts.lines++
		}
		if flags.c {
			// LineEndBytes added as the Scanner scans the lineSplit and it ignores the /n
			resCounts.bytes += len(bytesArray) + LineEndBytes
		}
		if flags.w {
			resCounts.words += len(strings.Fields(line))
		}
		if flags.m {
			resCounts.characters += utf8.RuneCountInString(line) + LineEndBytes
		}
	}
}

func (resCount *Count) printCounts(flags flagsStruct) {
	resArray := []int{}
	if flags.l {
		resArray = append(resArray, resCount.lines)
	}
	if flags.w {
		resArray = append(resArray, resCount.words)
	}
	if flags.m {
		resArray = append(resArray, resCount.characters)
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
