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
			t.Error("Expected %s, got %s", tt.expected, c.state)
		}

	}


}
