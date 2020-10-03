package main

import "fmt"

const width = 30

var p, q, r uint8

func main() {
	initialSate := [width]uint8{}
	initialSate[width/2] = 1

	states := [][width]uint8{}
	states = append(states, initialSate)

	for states[len(states)-1][0] == 0 {
		currentState := states[len(states)-1]
		fmt.Println(currentState)
		newState := evolve(currentState)
		states = append(states, newState)
	}

}

// Evolve a current state of life using a rule and return the evolved state.
func evolve(currentState [width]uint8) [width]uint8 {
	newState := [width]uint8{}
	for i := 0; i < width; i++ {

		if i > 0 {
			p = currentState[i-1]
		}

		q = currentState[i]

		if i < width-1 {
			r = currentState[i+1]
		}

		result := rule30(p, q, r)

		newState[i] = result
	}
	return newState
}

// rule30: (p, q, r) ↦ p XOR (q OR r)
func rule30(p uint8, q uint8, r uint8) uint8 {
	return p ^ (q | r)
}
