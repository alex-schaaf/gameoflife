package main

import "fmt"

const width = 30

func main() {
	initialSlice := [width]uint8{}
	initialSlice[width/2] = 1

	var p, q, r uint8

	slices := [][width]uint8{}
	slices = append(slices, initialSlice)

	for slices[len(slices)-1][0] == 0 {
		currentSlice := slices[len(slices)-1]
		fmt.Println(currentSlice)
		newSlice := [width]uint8{}
		for i := 0; i < width; i++ {

			if i > 0 {
				p = currentSlice[i-1]
			}

			q = currentSlice[i]

			if i < width-1 {
				r = currentSlice[i+1]
			}

			result := p ^ (q | r)

			newSlice[i] = result
		}
		slices = append(slices, newSlice)
	}

}
