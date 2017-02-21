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

func (c cell) String() string {
	return fmt.Sprintf("%q", c.state)
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

	// Rocks and food never change
	if currentState == cellRock || currentState == cellFood {
		c.state = currentState
		return
	}

	// Initialize the neighbor count lookup.
	neighborCounts := map[cellState]int{
		cellRock:     0,
		cellFood:     0,
		cellEmpty:    0,
		currentState: 0,
	}

	for _, neighbor := range c.neighbors {
		if neighbor.state == cellFood {

		}
		neighborCounts[neighbor.state]++
	}

	var speciesCounts map[cellState]int
	for species, count := range neighborCounts {
		speciesCounts[species] = count
	}

	if currentState == cellEmpty {

	}

	c.state = currentState

}

func (cs cellState) isPlayer() bool {
	_, err := strconv.ParseInt(string(cs), 10, 8)
	return err == nil
}
