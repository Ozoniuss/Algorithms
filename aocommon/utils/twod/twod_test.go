package twod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	type test struct {
		loc      Location
		dir      Direction
		expected Location
	}
	tests := []test{
		{
			loc:      Location{1, 1},
			dir:      Direction{-1, 1},
			expected: Location{0, 2},
		},
	}
	predefinedDirs := []Direction{
		LEFT, RIGHT, UP, DOWN, N, S, E, W, NE, NW, SE, SW,
	}
	for _, predefinedDir := range predefinedDirs {
		tests = append(tests, test{
			loc:      ORIGIN,
			dir:      predefinedDir,
			expected: Location{predefinedDir[0], predefinedDir[1]},
		})
	}
	for _, t := range tests {
		assert.Equal(t.expected, Add(t.loc, t.dir))
	}
}
