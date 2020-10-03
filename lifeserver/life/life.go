package main

import "fmt"

const width = 30

func main() {
	initialSate := [width]uint8{}
	initialSate[width/2] = 1

	var p, q, r uint8

	states := [][width]uint8{}
	states = append(states, initialSate)

	for states[len(states)-1][0] == 0 {
		currentState := states[len(states)-1]
		fmt.Println(currentState)
		newState := [width]uint8{}
		for i := 0; i < width; i++ {

			if i > 0 {
				p = currentState[i-1]
			}

			q = currentState[i]

			if i < width-1 {
				r = currentState[i+1]
			}

			result := p ^ (q | r)

			newState[i] = result
		}
		states = append(states, newState)
	}

}
