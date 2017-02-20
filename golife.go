package GoLife

import (
	"fmt"
	"strconv"
)

type cell struct {
	state cellState
	mirrorCell *cell
	neighbors []*cell
}

type cellState string

func (c cell) String() string {
	return fmt.Sprintf("%q", c.state)
}

const (
	cellEmpty cellState = cellState("E")
	cellFood = cellState("F")
	cellRock = cellState("R")
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


	if currentState == cellEmpty {

	}

	c.state = currentState

}

func (c *cell) isPlayer() bool {
	_, err := strconv.ParseInt(string(c.state), 10, 8)
	return err == nil
}
