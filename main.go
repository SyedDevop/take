package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SyedDevop/take/path"
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
	p := flag.Arg(0)
	if p == "" {
		fmt.Println("Error: A path argument is required")
		PrintUsage()
		os.Exit(1)
	}

	if *f {
		fmt.Print(path.Base(p))
	} else if *d {
		fmt.Print(path.Dir(p))
	}
}
