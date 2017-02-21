package GoLife

import (
	"fmt"
	"strconv"
)

type cell struct {
	state      cellState
	mirrorCell *cell
	neighbors  []*cell
}

type cellState string

func (c *cell) String() string {
	return fmt.Sprintf("%s", string(c.state))
}

func (cs cellState) String() string {
	return string(cs)
}

const (
	cellEmpty cellState = cellState("E")
	cellFood            = cellState("F")
	cellRock            = cellState("R")
)

func (c *cell) updateValue() {
	// The current state is actually the state of the mirror cell, as that
	// determines whether we may die or be born.
	currentState := c.mirrorCell.state

	// Rocks and food never change.
	if currentState == cellRock || currentState == cellFood {
		c.state = currentState
		return
	}

	foodCount := 0
	liveNeighbors := 0
	maxSpecies := 0
	var candidateSpecies cellState
	var speciesCounts = make(map[cellState]int)

	for _, neighbor := range c.neighbors {
		if neighbor.state == cellFood {
			foodCount++
		} else if neighbor.isPlayer() {
			if _, ok := speciesCounts[neighbor.state]; !ok {
				speciesCounts[neighbor.state] = 0
			}
			speciesCounts[neighbor.state]++
			liveNeighbors++
			if speciesCounts[neighbor.state] > maxSpecies {
				maxSpecies = speciesCounts[neighbor.state]
				candidateSpecies = neighbor.state
			}
		}
	}

	if currentState == cellEmpty && liveNeighbors >= 1 && liveNeighbors <= 4 && maxSpecies+foodCount >= 3 {
		c.state = candidateSpecies
	} else if c.isPlayer() && (liveNeighbors >= 4 || maxSpecies+foodCount < 2) {
		c.state = cellEmpty
	} else {
		c.state = currentState
	}
}

func (c *cell) isPlayer() bool {
	_, err := strconv.ParseInt(string(c.state), 10, 8)
	return err == nil
}
