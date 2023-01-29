package twod

// Location represents the two-dimensional integer coordinates of a point in
// the place.
type Location [2]int

// Direction consists of two values, representing the values added to each
// coordinate when moving to that direction. It can be thought of as a vector
// in the plane.
type Direction [2]int

// Move takes an initial location and a direction, and returns the new location
// after moving in that direction.
func Move(l Location, d Direction) Location {
	return Location{l[0] + d[0], l[1] + d[1]}
}

// AddDirs adds two directions by individually adding the direction's horizontal
// and vertical component. It works the same way as adding two vectors.
func AddDirs(d1, d2 Direction) Direction {
	return Direction{d1[0] + d2[0], d1[1] + d2[1]}
}

var (
	// Origin represents the center of the plane, or (0,0).
	ORIGIN Location = Location{0, 0}
)

var (
	// Moves one unit to the left of the x-axis.
	LEFT Direction = Direction{0, -1}
	// Moves one unit to the right of the x-axis.
	RIGHT Direction = Direction{0, 1}
	// Moves one unit upwards the y-axis. Likely for parsing reasons, in all
	// Advent-of-Code problems the y-axis is reversed.
	UP Direction = Direction{-1, 0}
	// Moves one unit downwards the y-axis. Likely for parsing reasons, in all
	// Advent-of-Code problems the y-axis is reversed.
	DOWN Direction = Direction{1, 0}
	// Doesn't move.
	NULL Direction = Direction{0, 0}
)

/* Directions, expressed as cardinal directions. */

var (
	// Moves one unit upwards the y-axis. Likely for parsing reasons, in all
	// Advent-of-Code problems the y-axis is reversed.
	N Direction = UP
	// Moves one unit upwards the y-axis. Likely for parsing reasons, in all
	// Advent-of-Code problems the y-axis is reversed.
	S Direction = DOWN
	// Moves one unit to the left of the x-axis.
	W Direction = LEFT
	// Moves one unit to the right of the x-axis.
	E Direction = RIGHT
	// Moves one unit upwards the y-axis and and one unit to the right of the
	// y-axis.
	NE Direction = Direction{-1, 1}
	// Moves one unit upwards the y-axis and and one unit to the left of the
	// y-axis.
	NW Direction = Direction{-1, -1}
	// Moves one unit downwards the y-axis and and one unit to the right of the
	// y-axis.
	SE Direction = Direction{1, 1}
	// Moves one unit downwards the y-axis and and one unit to the right of the
	// y-axis.
	SW Direction = Direction{1, -1}
)
