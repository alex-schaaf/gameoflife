package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"math/rand"
	"os"
	"time"
)

func initialize() tcell.Screen {
	screen, e := tcell.NewScreen()
	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := screen.Init(); e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorWhite).
		Foreground(tcell.ColorBlack)
	screen.SetStyle(defStyle)
	return screen
}

func seedWorld(nx int, ny int, seed int64) [][]int {
	world := make([][]int, nx)
	rows := make([]int, nx*ny)
	for i := 0; i < nx; i++ {
		world[i] = rows[i*ny : (i+1)*ny]
	}

	rng := rand.New(rand.NewSource(seed))

	for x := 0; x < nx; x++ {
		for y := 0; y < ny; y++ {
			world[x][y] = rng.Intn(2)
		}
	}

	return world
}

// Check the neighbors of cell in world at given x, y coordinates
// and return the sum of the number of neighbors.
func checkNeighbors(x int, y int, world [][]int) int {
	var neighbors = 0

	neighbors += world[x-1][y]
	neighbors += world[x+1][y]
	neighbors += world[x][y+1]
	neighbors += world[x][y-1]

	neighbors += world[x-1][y+1]
	neighbors += world[x-1][y-1]
	neighbors += world[x+1][y+1]
	neighbors += world[x+1][y-1]
	return neighbors
}

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
	screen := initialize()
	//screen.SetContent(0,0, rune('h'), []rune(""), tcell.StyleDefault)
	screen.Show()

	var nx = 111
	var ny = 22

	world := seedWorld(nx, ny, 42)

	for i := 0; i < 100; i++ {
		newWorld := world
		for x := 1; x < nx-1; x++ {
			for y := 1; y < ny-1; y++ {
				neighbors := checkNeighbors(x, y, world)
				alive := world[x][y]

				if alive == 1 && neighbors < 2 {
					newWorld[x][y] = 0
				} else if alive == 1 && neighbors <= 3 {
					newWorld[x][y] = 1 // redundant
				} else if alive == 1 && neighbors > 3 {
					newWorld[x][y] = 0
				} else if alive == 0 && neighbors == 3 {
					newWorld[x][y] = 1
				}

				if world[x][y] == 1 {
					screen.SetContent(x, y, tcell.RuneBlock, []rune(""), tcell.StyleDefault)
				} else {
					screen.SetContent(x, y, rune(' '), []rune(""), tcell.StyleDefault)
				}

			}
		}
		screen.Show()
		time.Sleep(100 * time.Millisecond)
	}

	// key handling
	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}
