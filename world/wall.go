package world

import (
	"strings"

	"github.com/nfgrep/necromancer/entities"
	"github.com/nfgrep/necromancer/linalg"
)

// A map from x, y world coords to a wall
type Walls [][]*Wall

func WallsFromMap(grid Map, walls map[string]entities.WallEntity) [][]*Wall {

	// These walls don't yet have start and end points
	wallsByTerminal := wallsByTerminalSymbol(walls)

	// fmt.Println(wallsByTerminal)

	// Fill in the start and end points for each wall
	// Also returns a 2D array from world coords to walls
	wallsOut := fillStartAndEnd(wallsByTerminal, grid)

	return wallsOut
}

// TODO: this does too much
func fillStartAndEnd(walls map[string]Wall, grid Map) [][]*Wall {
	rows, cols := len(grid), len(grid[0])

	wallsOut := make([][]*Wall, rows)
	for i := range wallsOut {
		wallsOut[i] = make([]*Wall, cols)
	}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			wall, ok := walls[grid[y][x]]
			if ok {
				symbol := wall.Symbol

				// Check horizontal wall
				if x+1 < cols && grid[y][x+1] == symbol {
					end := findWallEnd(grid, x+1, y, symbol, true)
					wall.Start = linalg.Vec2{X: float64(x), Y: float64(y)}
					wall.End = end
					for i := x; i <= int(end.X); i++ {
						wallsOut[y][i] = &wall
					}
				}

				// Check vertical wall
				if y+1 < rows && grid[y+1][x] == symbol {
					end := findWallEnd(grid, x, y+1, symbol, false)
					wall.Start = linalg.Vec2{X: float64(x), Y: float64(y)}
					wall.End = end
					for i := y; i <= int(end.Y); i++ {
						wallsOut[i][x] = &wall
					}
				}
			}
		}
	}

	return wallsOut
}

func wallsByTerminalSymbol(entities map[string]entities.WallEntity) map[string]Wall {
	walls := make(map[string]Wall)
	// for _, entity := range entities {
	// 	walls[entity.TerminalSymbol] = Wall{
	// 		Height:         entity.Height,
	// 		Symbol:         entity.Symbol,
	// 		TerminalSymbol: entity.TerminalSymbol,
	// 	}
	// }
	// return walls

	for _, entity := range entities {
		walls[entity.TerminalSymbol] = Wall{
			Height:         entity.Height,
			Symbol:         entity.Symbol,
			TerminalSymbol: entity.TerminalSymbol,
		}
	}
	return walls
}

func findWallEnd(grid [][]string, x, y int, symbol string, horizontal bool) linalg.Vec2 {
	if horizontal {
		for x < len(grid[0]) && (grid[y][x] == symbol || strings.ToLower(grid[y][x]) == symbol) {
			x++
		}
		return linalg.Vec2{X: float64(x - 1), Y: float64(y)}
	}
	for y < len(grid) && (grid[y][x] == symbol || strings.ToLower(grid[y][x]) == symbol) {
		y++
	}
	return linalg.Vec2{X: float64(x), Y: float64(y - 1)}
}

type Wall struct {
	Height int
	// Assume these are sorted, and that the first point is the start, and the last point is the end
	// Parts []linalg.Vec2

	Start          linalg.Vec2
	End            linalg.Vec2
	Symbol         string
	TerminalSymbol string
}
