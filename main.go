package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/gdamore/tcell/v2"
)

func Init() {

}

// TODO: single ray dist, i.e. "viewLen" or something on Player
var player = &Player{
	x: 2, y: 2,
	rot: math.Pi / 2,
	rays: []*Ray{
		{rot: -0.55},
		{rot: -0.5},
		{rot: -0.45},
		{rot: -0.4},
		{rot: -0.35},
		{rot: -0.3},
		{rot: -0.25},
		{rot: -0.2},
		{rot: -0.15},
		{rot: -0.1},
		{rot: -0.05},
		{rot: 0.0},
		{rot: 0.05},
		{rot: 0.1},
		{rot: 0.15},
		{rot: 0.2},
		{rot: 0.25},
		{rot: 0.3},
		{rot: 0.35},
		{rot: 0.4},
		{rot: 0.45},
		{rot: 0.5},
		{rot: 0.55},
	},
	viewLen: 70,
}

var rayStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)

var styleMap = map[int]tcell.Style{
	1: tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlue),
	2: tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorTeal),
}

var worldMap = [][]int{
	{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 1, 1, 1, 1, 1, 1, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 1, 1, 1, 1, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2},
	{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range text {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func setContentEqualWidth(screen tcell.Screen, x int, y int, primary rune, combining []rune, style tcell.Style) {
	screen.SetContent((x*2)+1, y, ' ', nil, style)
	screen.SetContent(x*2, y, ' ', nil, style)
}

func drawMap(s tcell.Screen, inMap [][]int, style tcell.Style) {
	// TODO: is this legit? what if becomes float?
	numRows := len(inMap)
	numCols := len(inMap[0])
	for x := 0; x < numCols; x++ {
		for y := 0; y < numRows; y++ {
			if worldMap[y][x] != 0 {
				setContentEqualWidth(s, x, y, ' ', nil, styleMap[worldMap[y][x]])
			}
		}
	}
}

// Draws a vertivcal bar centered about y
func drawBar(s tcell.Screen, x, y, height int, style tcell.Style) {
	ytop := y - (height / 2)
	ybot := y + (height / 2)
	for y := ytop; y <= ybot; y++ {
		setContentEqualWidth(s, x, y, ' ', nil, style)
	}
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// TODO: fucking god damnit this is messy

// Some fancy version of bresenhams
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// http://members.chello.at/~easyfilter/Bresenham.pdf

// Returns the distance, and the value of the map it intersected.
// If no intersection, map value == 0
func castRay(screen tcell.Screen, worldMap [][]int, x0, y0, x1, y1 int, style tcell.Style) (float64, int) {
	dx := intAbs(x1 - x0)
	dy := -intAbs(y1 - y0)
	sx := -1
	if x0 < x1 {
		sx = 1
	}
	sy := -1
	if y0 < y1 {
		sy = 1
	}
	er := dx + dy

	x := x0
	y := y0

	rayLen := math.Sqrt(math.Pow(float64(x1-x0), 2) + math.Pow(float64(y1-y0), 2))
	dist := rayLen
	mapVal := 0
	for {
		setContentEqualWidth(screen, x, y, ' ', nil, style)

		// Get dists for drawing 3D scene
		if worldMap[y][x] != 0 {
			dist = math.Sqrt(math.Pow(float64(x-x0), 2) + math.Pow(float64(y-y0), 2))
			mapVal = worldMap[y][x]
			break
		}

		if x == x1 && y == y1 {
			break
		}

		e2 := 2 * er
		if e2 >= dy {
			if x == x1 {
				break
			}
			er = er + dy
			x = x + sx
		}

		if e2 <= dx {
			if y == y1 {
				break
			}
			er = er + dx
			y = y + sy
		}
	}
	return dist, mapVal
}

func drawPlayer(screen tcell.Screen, player *Player, style tcell.Style) {
	// Draw the player
	setContentEqualWidth(screen, int(player.x), int(player.y), ' ', nil, style)

	// Some debug
	drawText(screen, 2, 5, 50, 5, style, fmt.Sprintf("rot: %v", player.rot))
}

func drawScene(screen tcell.Screen, player *Player, worldMap [][]int, style tcell.Style) {
	// Get distances
	//dists := []int{}
	for i, ray := range player.rays {
		rx1 := player.x + math.Cos(ray.rot+player.rot)*float64(player.viewLen)
		ry1 := player.y + math.Sin(ray.rot+player.rot)*float64(player.viewLen)
		rayDist, mapVal := castRay(screen, worldMap, int(player.x), int(player.y), int(rx1), int(ry1), rayStyle)

		// Draw bar
		if int(rayDist) == player.viewLen {
			continue
		}
		// Projected onto the flat view/camera plane
		projectedRayDist := rayDist * math.Cos(ray.rot)
		drawBar(screen, i+len(worldMap[0]), 10, 40-int(projectedRayDist), styleMap[mapVal])

		//dists = append(dists, int(rayDist))
		drawText(screen, 2, i+30, 70, i+35, style, fmt.Sprintf("ray: %v, ray.rot: %v, rx1: %v, ry1: %v, rayDist: %v", i, ray.rot, rx1, ry1, rayDist))
	}

	// Draw bar for each distance
	//offset := len(worldMap[0]) // So that we render off to the right of the map, *2 because we're using double width
	//for i, dist := range dists {
	//	if dist == player.viewLen {
	//		continue
	//	}
	//	drawBar(screen, i+offset, 10, 40-dist, style)
	//}
}

func main() {
	// TODO: make these styles global consts or sometething?
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
	sceneStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlue)
	playerStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGreen)
	//rayStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	//s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	// Draw initial boxes
	//drawBox(s, 1, 1, 42, 7, boxStyle, "Click and drag to draw a box")
	//drawBox(s, 5, 9, 32, 14, boxStyle, "Press C to reset")

	// Event loop
	for {
		s.Clear()
		//for i := 0; i < 200; i++ {
		//	drawBar(s, i+px, 0, 20, boxStyle)
		//}
		//drawRays(s, player, worldMap, rayStyle)
		drawPlayer(s, player, playerStyle)
		drawScene(s, player, worldMap, sceneStyle)
		drawMap(s, worldMap, boxStyle)
		s.Show()

		handleInput(s)
	}
}

func handleInput(s tcell.Screen) {
	// Poll event
	ev := s.PollEvent()

	// Process event
	switch ev := ev.(type) {
	case *tcell.EventResize:
		s.Sync()
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
			s.Fini()
			os.Exit(0)
		} else if ev.Key() == tcell.KeyCtrlL {
			s.Sync()
		} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
			s.Clear()
		} else if ev.Rune() == 'w' { // Movement
			player.moveFwd(1, worldMap)
		} else if ev.Rune() == 'a' {
			player.moveLeft(1, worldMap)
		} else if ev.Rune() == 's' {
			player.moveBack(1, worldMap)
		} else if ev.Rune() == 'd' {
			player.moveRight(1, worldMap)
		} else if ev.Rune() == 'n' { // Rotation TODO: make arrow keys
			player.rotate(-0.1)
		} else if ev.Rune() == 'm' {
			player.rotate(0.1)
		}
	}
}
