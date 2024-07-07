package path_test

import (
	"testing"

	"github.com/SyedDevop/take/path"
	"github.com/stretchr/testify/assert"
)

func TestDirPath(t *testing.T) {
	pathTable := [][2]string{
		// {"", ""},
		{"./m", "m"},
		{"main.go", ""},
		{"./main.go", ""},
		{"./home/main.go", "home"},
		{".///home/main.go/", "home"},
		{"/home/main.go///", "home"},
		{"/home/go/main.go/", "home/go"},
	}

	for _, p := range pathTable {
		assert.Equal(t, p[1], path.Dir(p[0]))
	}
}

func TestFilePath(t *testing.T) {
	pathTable := [][2]string{
		// {"", ""},
		{"./m", ""},
		{"main.go", "main.go"},
		{"./main.go", "main.go"},
		{"./home/main.go", "main.go"},
		{".///home/main.go/", "main.go"},
		{"/home/main.go///", "main.go"},
		{"/home/go/main.go/", "main.go"},
	}

	for _, p := range pathTable {
		assert.Equal(t, p[1], path.Base(p[0]))
	}
}

var pathTable = []string{
	// {"", ""},
	"./m",
	"main.go",
	"./main.go",
	"./home/main.go",
	".///home/main.go/",
	"/home/main.go///",
	"/home/go/main.go/",
}

// func BenchmarkDir(b *testing.B) {
// 	for _, v := range pathTable {
// 		b.Run("Path :="+v, func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				path.Dir(v)
// 			}
// 		})
// 	}
// }

func BenchmarkFile(b *testing.B) {
	for _, v := range pathTable {
		b.Run("Path :="+v, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				path.Base(v)
			}
		})
	}
}
