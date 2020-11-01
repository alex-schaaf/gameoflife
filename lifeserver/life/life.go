package life

import "fmt"

var p, q, r int

// RunLife initializes and runs a 1-D Game of Life
func RunLife(width int) [][]int {
	initialSate := []int{}
	for i := 0; i < width; i++ {
		initialSate = append(initialSate, 0)
	}
	initialSate[width/2] = 1

	states := [][]int{}
	states = append(states, initialSate)

	for states[len(states)-1][0] == 0 {
		currentState := states[len(states)-1]
		fmt.Println(currentState)
		newState := evolve(currentState)
		states = append(states, newState)
	}
	return states
}

// Evolve a current state of life using a rule and return the evolved state.
func evolve(currentState []int) []int {
	newState := []int{}
	for i := 0; i < len(currentState); i++ {

		if i > 0 {
			p = currentState[i-1]
		}

		q = currentState[i]

		if i < len(currentState)-1 {
			r = currentState[i+1]
		}

		result := rule30(p, q, r)

		newState = append(newState, result)
	}
	return newState
}

// rule30: (p, q, r) ↦ p XOR (q OR r)
func rule30(p int, q int, r int) int {
	return p ^ (q | r)
}
