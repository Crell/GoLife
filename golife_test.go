package GoLife

import (
	"testing"
)

type cellTest struct {
	start cellState
	neighbors []*cell
	expected cellState
}

var cellTests = []cellTest {
	// Empty, 1 neighbor
	{start: cellEmpty, neighbors:[]*cell{&cell{state: cellEmpty}}, expected: cellEmpty},
	{start: cellEmpty, neighbors:[]*cell{&cell{state: cellState("1")}}, expected: cellEmpty},
}

func TestCellUpdateValue(t *testing.T) {

	for _, tt := range cellTests {
		c := &cell{
			mirrorCell: &cell{state: tt.start},
			neighbors: tt.neighbors,
		}

		c.updateValue()

		if tt.expected != c.state {
			t.Errorf("Expected %s, got %s", tt.expected, c.state)
		}
	}
}

func TestPlayerDetection(t *testing.T) {

	tests := []struct{
		state cellState
		expected bool
	} {
		{cellEmpty, false},
		{cellFood, false},
		{cellRock, false},
		{cellState("1"), true},
		{cellState("0"), true},
	}

	for _, tt := range(tests) {
		c := &cell{state: tt.state}
		if c.isPlayer() != tt.expected {
			t.Errorf("State %s: Got %s, expected %s", tt.state, c.isPlayer(), tt.expected)
		}
	}
}