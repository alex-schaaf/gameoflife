package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var nx int = 100
var ny int = 100
var iterations = 1000
var seed int64 = 42

func main() {
	now := time.Now()
	folder := now.Format("2006-01-02_15-04-05")
	err := os.Mkdir(folder, 0755)
	handle(err)

	fmt.Println("This is the Game of Life..")
	fmt.Println("Seeding a world of size", nx, "by", ny)

	world := seedWorld(nx, ny, seed)

	for i := 0; i < iterations; i++ {
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
			}
		}

		fname := fmt.Sprint(folder, "/", i, ".ppm")
		writeImage(fname, newWorld)
	}
}

// Lazy generic error handling.
func handle(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Creates the 2D array world given nx, ny cells along each axis
// and randomly generating the starting cell values using RNG.
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

// Write given 2D array to PPM image at given filepath.
func writeImage(fp string, arr [][]int) {
	f, err := os.Create(fp)
	handle(err)
	defer f.Close()

	w := bufio.NewWriter(f)

	fmt.Fprintf(w, "P1\n")
	fmt.Fprintf(w, "#\n")
	fmt.Fprintf(w, "%v ", len(arr))
	fmt.Fprintf(w, "%v\n", len(arr[0]))

	for i := range arr {
		for j := range arr[i] {
			if j < len(arr[i])-1 {
				fmt.Fprintf(w, "%v ", arr[i][j])
			} else {
				fmt.Fprintf(w, "%v\n", arr[i][j])
			}
		}
	}
	w.Flush()
}
