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
	for _, r := range []rune(text) {
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

var worldMap = [][]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
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
			if worldMap[y][x] == 1 {
				setContentEqualWidth(s, x, y, ' ', nil, style)
			}
		}
	}
}

// where ytop <= ybot
func drawBar(s tcell.Screen, x, ytop, ybot int, style tcell.Style) {
	for y := ytop; y <= ybot; y++ {
		s.SetContent(x, y, ' ', nil, style)
	}
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Some fancy version of bresenhams
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// http://members.chello.at/~easyfilter/Bresenham.pdf
func plotLine(screen tcell.Screen, x0, y0, x1, y1 int, style tcell.Style) {
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

	for {
		//screen.SetContent(x0, y0, ' ', nil, style)
		setContentEqualWidth(screen, x0, y0, ' ', nil, style)

		if x0 == x1 && y0 == y1 {
			break
		}

		e2 := 2 * er
		if e2 >= dy {
			if x0 == x1 {
				break
			}
			er = er + dy
			x0 = x0 + sx
		}

		if e2 <= dx {
			if y0 == y1 {
				break
			}
			er = er + dx
			y0 = y0 + sy
		}
	}
}

type Player struct {
	x   int
	y   int
	rot float64
}

// Moves player if possible, i.e. if there's no wall in the way
func (p *Player) move(dx, dy int, worldMap [][]int) {
	newX := player.x + dx
	newY := player.y + dy
	if worldMap[newY][newX] == 0 {
		player.x += dx
		player.y += dy
	}
}

func (p *Player) rotate(rad float64) {
	p.rot += rad
	// Wrap around. 0 rad == 2PI
	if p.rot >= (2 * math.Pi) {
		p.rot -= (2 * math.Pi)
	}
	if p.rot <= 0 {
		p.rot += (2 * math.Pi)
	}
}

var player = &Player{x: 2, y: 2, rot: 0.0}

func drawPlayer(screen tcell.Screen, player *Player, style tcell.Style) {
	setContentEqualWidth(screen, player.x, player.y, ' ', nil, style)
	// Get slope from rot in rad
	m := math.Tan(player.rot)

	// Get rise + run from slope in decimal form
	// https://www.mathsisfun.com/converting-decimals-fractions.html
	dy := int(m * 1000000000000) // Lets hope 100 is big enough to turn this into an int without precision issues
	dx := 1000000000000

	divisor := gcd(dx, dy)
	dy /= divisor
	dx /= divisor

	// If we're between pi/2 and 3pi/2, i.e. in the left half
	if m > (math.Pi/2) && m < ((3*math.Pi)/2) {
		dx = -dx
	}

	// Ohh boy, this is gonna be messy
	x1 := player.x + dx
	y1 := player.y + dy

	drawText(screen, 2, 5, 50, 5, style, fmt.Sprintf("rot: %v", player.rot))
	drawText(screen, 2, 15, 50, 15, style, fmt.Sprintf("x1: %v, y1: %v", x1, y1))
	plotLine(screen, player.x, player.y, x1, y1, style)
}

// GCDRemainder calculates GCD iteratively using remainder.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
	playerStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGreen)

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
		drawPlayer(s, player, playerStyle)
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
			player.move(0, -1, worldMap)
		} else if ev.Rune() == 'a' {
			player.move(-1, 0, worldMap)
		} else if ev.Rune() == 's' {
			player.move(0, 1, worldMap)
		} else if ev.Rune() == 'd' {
			player.move(1, 0, worldMap)
		} else if ev.Rune() == 'n' { // Rotation TODO: make arrow keys
			player.rotate(0.1)
		} else if ev.Rune() == 'm' {
			player.rotate(-0.1)
		}
		//case *tcell.EventMouse:
		//	x, y := ev.Position()
		//	button := ev.Buttons()
		//	// Only process button events, not wheel events
		//	button &= tcell.ButtonMask(0xff)

		//	if button != tcell.ButtonNone && ox < 0 {
		//		ox, oy = x, y
		//	}
		//	switch ev.Buttons() {
		//	case tcell.ButtonNone:
		//		if ox >= 0 {
		//			label := fmt.Sprintf("%d,%d to %d,%d", ox, oy, x, y)
		//			drawBox(s, ox, oy, x, y, boxStyle, label)
		//			ox, oy = -1, -1
		//		}
		//	}
	}
}
