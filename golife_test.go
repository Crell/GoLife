package GoLife

import (
	"testing"
)

type cellTest struct {
	start     cellState
	neighbors []cellState
	expected  cellState
}

const (
	cellPlayer = cellState("1")
	cellEnemy = cellState("2")
)

var cellTests = []cellTest{
	// Empty, 1 neighbor.
	{start: cellEmpty, neighbors: []cellState{cellEmpty}, expected: cellEmpty},
	{start: cellEmpty, neighbors: []cellState{cellPlayer}, expected: cellEmpty},
	// Empty, 2 neighbors.
	{start: cellEmpty, neighbors: []cellState{cellEmpty, cellPlayer}, expected: cellEmpty},
	{start: cellEmpty, neighbors: []cellState{cellPlayer, cellPlayer}, expected: cellEmpty},
	// Empty, 3 neighbors.
	{start: cellEmpty, neighbors: []cellState{cellPlayer, cellPlayer, cellPlayer}, expected: cellPlayer},
	{start: cellEmpty, neighbors: []cellState{cellPlayer, cellPlayer, cellFood}, expected: cellPlayer},
	// Empty, 4 neighbors.
	{start: cellEmpty, neighbors: []cellState{cellPlayer, cellPlayer, cellPlayer, cellEmpty}, expected: cellPlayer}, // Born from 3 neighbors.
	{start: cellEmpty, neighbors: []cellState{cellPlayer, cellEnemy, cellPlayer, cellRock}, expected: cellEmpty},       // Hostile neighbor prevents birth.
	// Living, 1 neighbor.
	{start: cellPlayer, neighbors: []cellState{cellPlayer}, expected: cellEmpty},
	// Living, 2 neighbors.
	{start: cellPlayer, neighbors: []cellState{cellPlayer, cellPlayer}, expected: cellPlayer},
	{start: cellPlayer, neighbors: []cellState{cellPlayer, cellFood}, expected: cellPlayer},

}

func TestCellUpdateValue(t *testing.T) {

	for iteration, tt := range cellTests {
		neighbors := make([]*cell, len(tt.neighbors))
		for i, state := range tt.neighbors {
			neighbors[i] = &cell{state: state}
		}
		c := &cell{
			state:      cellEmpty,
			mirrorCell: &cell{state: tt.start},
			neighbors:  neighbors,
		}

		c.updateValue()

		if tt.expected != c.state {
			t.Errorf("Iteration %d: Expected %s, got %s", iteration, tt.expected, c.state)
		}
	}
}

func TestPlayerDetection(t *testing.T) {

	tests := []struct {
		state    cellState
		expected bool
	}{
		{cellEmpty, false},
		{cellFood, false},
		{cellRock, false},
		{cellPlayer, true},
		{cellEnemy, true},
	}

	for _, tt := range tests {
		c := &cell{state: tt.state}
		if c.isPlayer() != tt.expected {
			t.Errorf("State %s: Got %s, expected %s", tt.state, c.isPlayer(), tt.expected)
		}
	}
}
