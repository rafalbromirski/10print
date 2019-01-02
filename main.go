package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/namsral/flag"
)

var (
	cols = flag.Int("cols", 8, "Number of columns")
	rows = flag.Int("rows", 16, "Number of rows")
)

// Grid is responsible for generating random "grid"
//
// It implements:
//  > 10 PRINT CHR$(205.5+RND(1)); : GOTO 10
type Grid struct {
	r    *rand.Rand
	cols int
	rows int
}

// NewGrid constructs NewGrid
func NewGrid(cols, rows int) *Grid {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	return &Grid{
		r:    r,
		cols: cols,
		rows: rows,
	}
}

// RandomChar randomly picks "\" or "/""
func (g *Grid) RandomChar() string {
	if g.r.Float64() <= 0.5 {
		return fmt.Sprint("\u2571")
	}

	return fmt.Sprint("\u2572")
}

// Generate creates random pattern
func (g *Grid) Generate() string {
	var arr []string

	for i := 0; i < g.rows; i++ {
		var s string

		for j := 0; j < g.cols; j++ {
			s += g.RandomChar()
		}

		arr = append(arr, s)
	}

	return strings.Join(arr, "\n")
}

func main() {
	flag.Parse()

	g := NewGrid(*cols, *rows)
	msg := g.Generate()

	log.Printf("\n%s\n", msg)
}
