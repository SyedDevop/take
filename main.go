package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SyedDevop/take/path"
)

const VERSION = "G Path 1.0.0:V"

const USAGE = `Usage: g_Path [options] <path>
Options:
  -f        Extract and display the file name from the given path (e.g., fod/bar/bass.go -> bass.go)
  -d        Extract and display the directory path from the given path (e.g., fod/bar/bass.go -> fod/bar)
  -h        Show help and usage information
  -v        Show help
Example:
  g_Path -f fod/bar/bass.go
  g_Path -d fod/bar/bass.go
`

var (
	f = flag.Bool("f", false, "Extract and display the file name from the given path")
	d = flag.Bool("d", false, "Extract and display the directory path from the given path")
	h = flag.Bool("h", false, "Show help and usage information")
	v = flag.Bool("v", false, "Show version")
)

func PrintUsage() {
	fmt.Print(USAGE)
}

func init() {
	flag.Parse()

	if *v {
		fmt.Println(VERSION)
		os.Exit(0)
	}
	if *h {
		fmt.Println(VERSION)
		PrintUsage()
		os.Exit(0)
	}
}

func main() {
	p := flag.Arg(0)

	if !*f && !*d {
		return
	}

	if p == "" {
		return
	}

	if *f {
		fmt.Print(path.Base(p))
	} else if *d {
		fmt.Print(path.Dir(p))
	}
}
