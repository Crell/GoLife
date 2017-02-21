package GoLife

import (
	"testing"
)

type cellTest struct {
	start     cellState
	neighbors []cellState
	expected  cellState
}

var cellTests = []cellTest{
	// Empty, 1 neighbor.
	{start: cellEmpty, neighbors: []cellState{cellEmpty}, expected: cellEmpty},
	{start: cellEmpty, neighbors: []cellState{cellState("1")}, expected: cellEmpty},
	// Empty, 2 neighbors.
	{start: cellEmpty, neighbors: []cellState{cellEmpty, cellState("1")}, expected: cellEmpty},
	{start: cellEmpty, neighbors: []cellState{cellState("1"), cellState("1")}, expected: cellEmpty},
	// Empty, 3 neighbors.
	{start: cellEmpty, neighbors: []cellState{cellState("1"), cellState("1"), cellState("1")}, expected: cellState("1")},
	{start: cellEmpty, neighbors: []cellState{cellState("1"), cellState("1"), cellFood}, expected: cellState("1")},
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
		{cellState("1"), true},
		{cellState("0"), true},
	}

	for _, tt := range tests {
		c := &cell{state: tt.state}
		if c.isPlayer() != tt.expected {
			t.Errorf("State %s: Got %s, expected %s", tt.state, c.isPlayer(), tt.expected)
		}
	}
}
