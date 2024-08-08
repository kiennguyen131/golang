package basic

import (
	"testing"
)

func Test(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	actual := AddOne(1)
	if actual != output {
		t.Errorf("AddOne(%d), input %d, actual = %d", input, output, actual)
	}

	// assert.Equal(t, AddOne(2), 4, "AddOne (2) should be 3")
}
