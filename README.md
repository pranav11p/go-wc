# go-wc

A go-wc cli tool (like wc in Unix) written in GoLang. This tool is for counting the number of Words, Lines, Bytes and Character in the file. or a standard input.

### Features

- Counts words, lines, bytes, and characters
- Supports reading from files or standard input
- Detailed output with options to show individual counts and total sums

### Getting Started

To run go-wc, you will need

1. [GoLang](https://go.dev/)
2. Any of the below supported platform environment:
   - [Linux based environment](https://en.wikipedia.org/wiki/Comparison_of_Linux_distributions)
   - [OSX (Darwin) based environment](https://en.wikipedia.org/wiki/MacOS)

Follow the steps to get started with go-wc

1. Clone the repository using Git:

   ```
   git clone https://github.com/pranavpatel3012/go-wc.git
   ```

2. Go to the project directory

   ```
   cd go-wc
   ```

3. Build the binary:

   ```
   go build
   ```

### Usage

```
go-wc [flags] [filename]

Flags:
  -h, --help   help for go-wc
  -c    Count the number of bytes
  -l    Count the number of lines
  -m    Count the number of characters
  -w    Count the number of words
```

If no flags are specified, the default output of go-wc is the number of words, lines, bytes, and characters for the file. You can use the -m and -c flags to display only characters and bytes, and -w and -l to focus on words or lines respectively.

#### Example:

1. Get the number of words of a file test.txt

   ```
   ./go-wc -w test.txt
   ```

2. Get the number of lines of a file test.txt

   ```
   ./go-wc -l test.txt
   ```
