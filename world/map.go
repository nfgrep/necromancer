package world

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nfgrep/necromancer/gfx"
)

type Map [][]int

// TODO: I'm assuming that that passing the map by value is fine, since it's just a pointer to an array of arrays of ints.
func (m Map) at(x, y float64) int {
	return m[int(y)][int(x)]
}

func (m Map) Width() int {
	return len(m[0]) // NOTE: this assumes map is a rectangle, all rows being equal length
}

func (m Map) Height() int {
	return len(m) // Because our map is y, x, not x, y
}

func (m Map) WallExistsAt(x, y float64) bool {
	return m.at(x, y) != 0
}

func (m Map) Draw(s tcell.Screen, styleMap map[int]tcell.Style) {
	// TODO: is this legit? what if becomes float?
	numRows := len(m)
	numCols := len(m[0])
	for x := 0; x < numCols; x++ {
		for y := 0; y < numRows; y++ {
			if m[y][x] != 0 {
				gfx.SetContentEqualWidth(s, x, y, ' ', nil, styleMap[m[y][x]])
			}
		}
	}
}

// ------------ Generating walls ------------

// We expect a total of 3 walls here
var testMap = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 1, 1, 1, 1, 0},
	{0, 1, 0, 0, 0, 0, 1, 0},
	{0, 1, 0, 0, 0, 0, 1, 0},
	{0, 1, 1, 1, 1, 1, 1, 0},
	{2, 2, 2, 2, 0, 0, 0, 0},
	{0, 0, 0, 2, 0, 0, 0, 0},
	{0, 0, 0, 2, 0, 0, 0, 0},
	{0, 0, 0, 2, 0, 0, 0, 0},
	{2, 2, 2, 2, 0, 3, 3, 3},
}

// func generateWalls(worldMap Map) ([]Wall, error) {
// 	done := false
// 	pos := linalg.Vec2{0, 0}
// 	//walls := []Wall{}
// 	for !done {
// 		fmt.Println("pos", pos)
// 		val := worldMap.at(pos.x, pos.y)
// 		if val != 0 { // We found a wall
// 			p0 := pos // This is the "start" of the wall
// 			// First we check if neighbor in any direction
// 			firstNeighbor, err := findNeighbor(p0, val, worldMap)
// 			if err != nil {
// 				return nil, fmt.Errorf("error finding first neighbor: %v", err)
// 			}

// 			// Direction from p0 to the first neighbor
// 			direction := firstNeighbor.sub(p0).normalized()

// 			// Now we find the rest of the wall
// 			// TODO: we need to keep track of how much of the map we've already seen
// 			// we should come up with an actual algorithm for this
// 			nextSegment, err := findNextWallSegment(firstNeighbor, *direction, val, worldMap)
// 			if err != nil {
// 				return nil, fmt.Errorf("error finding next wall segment: %v", err)
// 			}

// 			fmt.Println("----------")
// 			fmt.Println("p0:", p0)
// 			fmt.Println("direction:", direction)
// 			fmt.Println("firstNeighbor:", firstNeighbor)
// 			fmt.Println("nextSegment:", nextSegment)

// 			done = true
// 		}
// 	}
// 	return nil, fmt.Errorf("not implemented")
// }

// // TODO: this is a little chaotic, and converting float to int feels weird here
// // Returns the position of the first neighbor found, iterates in a clockwise fashion starting top left
// func findNeighbor(point linalg.Vec2, val int, worldMap Map) (linalg.Vec2, error) {
// 	// Top left
// 	if point.x > 0 && point.y > 0 {
// 		if worldMap.at(point.x-1, point.y-1) == val {
// 			return linalg.Vec2{point.x - 1, point.y - 1}, nil
// 		}
// 	}
// 	// Top
// 	if point.y > 0 {
// 		if worldMap.at(point.x, point.y-1) == val {
// 			return linalg.Vec2{point.x, point.y - 1}, nil
// 		}
// 	}
// 	// Top right
// 	if int(point.x) < worldMap.width()-1 && point.y > 0 {
// 		if worldMap.at(point.x+1, point.y-1) == val {
// 			return linalg.Vec2{point.x + 1, point.y - 1}, nil
// 		}
// 	}
// 	// Right
// 	if int(point.x) < worldMap.width()-1 {
// 		if worldMap.at(point.x+1, point.y) == val {
// 			return linalg.Vec2{point.x + 1, point.y}, nil
// 		}
// 	}
// 	// Bottom right
// 	if int(point.x) < worldMap.width()-1 && int(point.y) < worldMap.height()-1 {
// 		if worldMap.at(point.x+1, point.y+1) == val {
// 			return linalg.Vec2{point.x + 1, point.y + 1}, nil
// 		}
// 	}
// 	// Bottom
// 	if int(point.y) < worldMap.height()-1 {
// 		if worldMap.at(point.x, point.y+1) == val {
// 			return linalg.Vec2{point.x, point.y + 1}, nil
// 		}
// 	}
// 	// Bottom left
// 	if point.x > 0 && int(point.y) < worldMap.height()-1 {
// 		if worldMap.at(point.x-1, point.y+1) == val {
// 			return linalg.Vec2{point.x - 1, point.y + 1}, nil
// 		}
// 	}
// 	// Left
// 	if point.x > 0 {
// 		if worldMap.at(point.x-1, point.y) == val {
// 			return linalg.Vec2{point.x - 1, point.y}, nil
// 		}
// 	}

// 	return linalg.Vec2{}, fmt.Errorf("no neighbor found")
// }

// // TODO: make sure this doesn't segfault if we're at the edge of the map
// // Returns the next point in the wall, trying in direction first, then anti-direction
// // direction should be normalized i.e. x and y values are in the range -1, 1 inclusive
// // direction should also probably be orthogonal, i.e. not diagonal
// func findNextWallSegment(point, direction linalg.Vec2, val int, worldMap Map) (linalg.Vec2, error) {
// 	if worldMap.at(point.x+direction.x, point.y+direction.y) == val {
// 		return linalg.Vec2{point.x + direction.x, point.y + direction.y}, nil
// 	} else if worldMap.at(point.x-direction.x, point.y-direction.y) == val {
// 		return linalg.Vec2{point.x - direction.x, point.y - direction.y}, nil
// 	}
// 	return linalg.Vec2{0, 0}, fmt.Errorf("no next wall segment found")
// }
