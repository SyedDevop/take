package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const USAGE = `Usage: take [options] <path>
Options:
  -f        Extract and display the file name from the given path (e.g., fod/bar/bass.go -> bass.go)
  -d        Extract and display the directory path from the given path (e.g., fod/bar/bass.go -> fod/bar)
  -h        Show help and usage information
Example:
  take -f fod/bar/bass.go
  take -d fod/bar/bass.go
`

var (
	f = flag.Bool("f", false, "Extract and display the file name from the given path")
	d = flag.Bool("d", false, "Extract and display the directory path from the given path")
	h = flag.Bool("h", false, "Show help and usage information")
)

func PrintUsage() {
	fmt.Print(USAGE)
}

func init() {
	flag.Parse()
	if *h {
		PrintUsage()
		os.Exit(0)
	}
	if !*f && !*d {
		PrintUsage()
		os.Exit(0)
	}
}

func main() {
	path := flag.Arg(0)
	if path == "" {
		fmt.Println("Error: A path argument is required")
		PrintUsage()
		os.Exit(1)
	}

	if *f {
		fmt.Println(filepath.Base(path))
	} else if *d {

		endIndex := len(path) - 1

		// Return empty string if path is empty.
		if endIndex < 0 {
			fmt.Print("")
			return
		}

		for {
			switch path[endIndex] {
			case '.':
				fmt.Println("d", filepath.Dir(path))
				return
			case '/':
				fmt.Println(path)
				return
			}
			endIndex--
		}
	}
}
