package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/gdamore/tcell/v2"
)

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

// This is the current height of the map
var debugTextY int = 30
var debugTextWidth int = 200

func drawDebugText(s tcell.Screen, style tcell.Style, text string) {
	drawText(s, 0, debugTextY, debugTextWidth, debugTextY+5, style, text)
}

func setContentEqualWidth(screen tcell.Screen, x int, y int, primary rune, combining []rune, style tcell.Style) {
	screen.SetContent((x*2)+1, y, ' ', nil, style)
	screen.SetContent(x*2, y, ' ', nil, style)
}

//type Object interface {
//worldToLocal(x, y int) (int, int)
//}

//func worldToObjectSpace(x, y int, object) (int, int) {
//	// Get a wall from the world coords
//
//}

// Draws a vertivcal bar centered about y
func drawBar(s tcell.Screen, screenX, screenY, height int, texSlice []tcell.Style) {
	//compressedTex := minifyTextureSlice(texSlice, height)
	//ytop := screenY - (height / 2)
	//ybot := screenY + (height / 2)
	//for y := ytop; y <= ybot; y++ {
	//	//wallX, wallY := worldToObjectSpace(screenX, y) // y will probably be the same for world and object space
	//	//u, v := projectorFunction(wall, wallX, wallY)``
	//	setContentEqualWidth(s, screenX, y, ' ', nil, tex[(y+ybot)])
	//}
	ytop := screenY - (height / 2)
	for i, style := range texSlice { // Assuming len(texSlice) == height
		y := i + ytop
		setContentEqualWidth(s, screenX, y, ' ', nil, style)
	}
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Vec2 struct {
	x float64
	y float64
}

type Vec3 struct {
	x float64
	y float64
	z float64
}

// TODO: fucking god damnit this is messy

// Some fancy version of bresenhams
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// http://members.chello.at/~easyfilter/Bresenham.pdf

// Returns the distance, and the value of the map it intersected.
// If no intersection, map value == 0
func castRay(screen tcell.Screen, worldMap [][]int, x0, y0, x1, y1 int, style tcell.Style) *Vec2 {
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

	//rayLen := math.Sqrt(math.Pow(float64(x1-x0), 2) + math.Pow(float64(y1-y0), 2))
	//dist := rayLen
	intersectionPoint := &Vec2{}
	for {
		setContentEqualWidth(screen, x, y, ' ', nil, style)

		// Get dists for drawing 3D scene
		if worldMap[y][x] != 0 {
			//dist = math.Sqrt(math.Pow(float64(x-x0), 2) + math.Pow(float64(y-y0), 2))
			intersectionPoint = &Vec2{float64(x), float64(y)}
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
	return intersectionPoint
}

func drawPlayer(screen tcell.Screen, player *Player, style tcell.Style) {
	// Draw the player
	setContentEqualWidth(screen, int(player.x), int(player.y), ' ', nil, style)

	// Some debug
	drawDebugText(screen, style, fmt.Sprintf("player rot: %v", player.rot))
}

var maxWallHeight int = 50
var horizonYPos int = 15

func drawScene(screen tcell.Screen, player *Player, worldMap [][]int, style tcell.Style) {
	// Get distances
	//dists := []int{}
	for i, ray := range player.rays {
		rx1 := player.x + math.Cos(ray.rot+player.rot)*float64(player.viewLen)
		ry1 := player.y + math.Sin(ray.rot+player.rot)*float64(player.viewLen)
		intersect := castRay(screen, worldMap, int(player.x), int(player.y), int(rx1), int(ry1), player.rayStyle)
		rayDist := math.Sqrt(math.Pow(float64(player.x-intersect.x), 2) + math.Pow(float64(player.y-intersect.y), 2))

		// -- Draw bar
		// Continue if we didn't intersect anything
		if int(rayDist) == player.viewLen {
			continue
		}
		// Projected onto the flat view/camera plane
		// Minimizes fish-eye effect a bit
		correctedRayDist := rayDist * math.Cos(ray.rot)
		barHeight := maxWallHeight - int(correctedRayDist)
		if barHeight == 0 {
			continue
		}

		tex := wallTexture
		texX := int(intersect.x) % len(tex[0])
		texSlice := getTexSlice(wallTexture, texX)
		filteredTexSlice := filterTexSlice(texSlice, barHeight)
		// To push the scene view to the right of the map
		screenXOffset := len(worldMap[0])
		//barStyle := styleMap[worldMap[int(intersect.y)][int(intersect.x)]]
		drawBar(screen, i+screenXOffset, horizonYPos, barHeight, filteredTexSlice)
		// -- end Draw bar

		//dists = append(dists, int(rayDist))
		//drawText(screen, 2, i+30, 70, i+35, style, fmt.Sprintf("ray: %v, ray.rot: %v, rx1: %v, ry1: %v, rayDist: %v", i, ray.rot, rx1, ry1, rayDist))
	}

}

func getTexSlice(tex Texture, texX int) []tcell.Style {
	// Generate the vertical slice of texture for this bar
	//barTex := textures[worldMap[int(intersect.y)][int(intersect.x)]]
	texSlice := []tcell.Style{}
	for _, horiz := range tex {
		//textureColumn := wallTextureMap[*intersect]
		texSlice = append(texSlice, horiz[texX])
	}
	return texSlice
}

func filterTexSlice(texSlice []tcell.Style, height int) []tcell.Style {
	// Project texVslice onto the height of a wall
	projectedTexVSlice := []tcell.Style{}
	texelsPerPixel := float64(len(texSlice)) / float64(height)
	j := texelsPerPixel
	for j < float64(len(texSlice)) {
		projectedTexVSlice = append(projectedTexVSlice, texSlice[(int(j))])
		j += texelsPerPixel
	}
	return projectedTexVSlice
}

// Maps a x, y world coord to the index of a vertical slice of a texture
//var wallTextureMap map[Point]int

//func Init() {
//	wallTextureMap = generateWallTexMap(worldMap)
//}

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
			player.rotate(-0.1)
		} else if ev.Rune() == 's' {
			player.moveBack(1, worldMap)
		} else if ev.Rune() == 'd' {
			player.rotate(0.1)
		} else if ev.Rune() == 'n' {
			player.moveLeft(1, worldMap)
		} else if ev.Rune() == 'm' {
			player.moveRight(1, worldMap)
		}
	}
}
