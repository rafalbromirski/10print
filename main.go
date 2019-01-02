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
	rows = flag.Int("rows", 16, "Number of rows")
	cols = flag.Int("cols", 8, "Number of columns")
)

// Grid is responsible for generating random "grid"
//
// It implements:
//  > 10 PRINT CHR$(205.5+RND(1)); : GOTO 10
type Grid struct {
	r    *rand.Rand
	rows int
	cols int
}

// NewGrid constructs NewGrid
func NewGrid(rows, cols int) *Grid {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	return &Grid{r, rows, cols}
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

	for i := 0; i < g.cols; i++ {
		var s string

		for j := 0; j < g.rows; j++ {
			s += g.RandomChar()
		}

		arr = append(arr, s)
	}

	return strings.Join(arr, "\n")
}

func main() {
	flag.Parse()

	g := NewGrid(*rows, *cols)
	msg := g.Generate()

	log.Printf("\n%s\n", msg)
}
