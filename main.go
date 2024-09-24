package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/nfgrep/necromancer/entities"
	"github.com/nfgrep/necromancer/gfx"
	"github.com/nfgrep/necromancer/linalg"
	"github.com/nfgrep/necromancer/player"
	"github.com/nfgrep/necromancer/world"
)

// maps world.WorldMap values to styles
var styleMap = map[string]tcell.Style{
	"a": tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlue),
	"w": tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorTeal),
}

// TODO: make a 'scene' package?
func drawScene(screen tcell.Screen, player *player.Player, worldMap world.Map, walls [][]*world.Wall, style tcell.Style) {
	intersections := player.CastViewRays(worldMap, screen, rayStyle)

	textureSlices := make([][]tcell.Style, len(intersections))
	for i, intersect := range intersections {
		wall := walls[int(intersect.Y)][int(intersect.X)]
		if wall != nil {
			textureSlices[i] = findTextureSlice(wall, intersect)
		}
	}

	dists := player.CalculateViewDistances(intersections)
	maxHeight := 30
	horizonYPos := 20
	gfx.DrawBarsForDists(screen, dists, player.ViewLen, maxHeight, horizonYPos, worldMap.Width(), style, textureSlices)
}

func findTextureSlice(wall *world.Wall, intersect *linalg.Vec2) []tcell.Style {
	wallVector := wall.End.Sub(wall.Start)
	intersectVector := intersect.Sub(wall.Start)

	percentAlongWall := intersectVector.Magnitude() / wallVector.Magnitude()

	texture := wall.GetTexture()

	texWidth := len(texture[0])

	// Subtract 1 because we're using 0-indexing
	textureSlice := texture[int(percentAlongWall*float64(texWidth-1))]
	return textureSlice
}

var sceneStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlue)
var rayStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)

type NecroYml struct {
	// Maps is a hashmap of maps each with a key/name, and each map is a 2d array of strings
	Maps map[string]world.Map `yaml:"maps"`
	// Types is a hashmap of types for which each unit of the map is of that type
	Types map[string]string `yaml:"types"`
}

func main() {

	parsedEntities, err := entities.ParseEntitiesFromFile("necro.yml")
	if err != nil {
		log.Fatalf("unable to parse entities: %v", err)
	}

	// Look through entities and find all the wall entities
	// A map from terminal symbol to wall entity
	wallEntities := make(map[string]entities.WallEntity)
	for _, entity := range parsedEntities {
		if wallEntity, ok := entity.(*entities.WallEntity); ok {
			wallEntities[wallEntity.TerminalSymbol] = *wallEntity
		}
	}

	// TODO: only load once
	maps, err := world.MapsFromConfig("necro.yml")
	if err != nil {
		log.Fatalf("unable to parse world: %v", err)
	}

	worldMap := maps["spawn"]

	walls := world.WallsFromMap(worldMap, wallEntities)

	// fmt.Println(walls[11][1])

	// return

	// walls := world.GenerateWalls(worldMap, []string{"spawn"})
	// fmt.Println(walls)

	// return

	// fmt.Println(worldMap)
	// TODO: make these styles global consts or sometething?
	//boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
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
	defaultStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defaultStyle)
	//s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	// Pass this to our debug stuff so we can call DrawDebug from anywhere without passing a screen aroung
	gfx.SetDebugScreen(s)

	var playerPos = linalg.Vec2{X: 2, Y: 2}
	//var playerFwd = linalg.Vec2{X: 0, Y: 0}
	var playerRot = 0.0
	var rayCount = 12
	var fov = 0.72

	var p = &player.Player{
		Position: &playerPos,
		// Fwd:      &playerFwd,
		Rotation: &playerRot,
		ViewRays: player.GenerateRays(rayCount, fov, &playerPos, &playerRot),
		// rayStyle: tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed),
		ViewLen: 70,
	}
	// Event loop
	for {
		s.Clear()
		//for i := 0; i < 200; i++ {
		//	drawBar(s, i+px, 0, 20, boxStyle)
		//}
		//drawRays(s, player, world.WorldMap, rayStyle)
		//drawPlayer(s, player.Player, playerStyle)
		//drawMap(s, world.WorldMap, boxStyle)
		worldMap.Draw(s, styleMap)
		p.Draw(s, playerStyle)
		drawScene(s, p, worldMap, walls, sceneStyle)

		s.Show()

		handleInput(s, p, worldMap)
	}
}

func handleInput(s tcell.Screen, p *player.Player, worldMap world.Map) {
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
			p.MoveFwd(1, worldMap)
		} else if ev.Rune() == 'a' {
			p.Rotate(-0.1)
		} else if ev.Rune() == 's' {
			p.MoveBack(1, worldMap)
		} else if ev.Rune() == 'd' {
			p.Rotate(0.1)
		} else if ev.Rune() == 'n' {
			p.MoveLeft(1, worldMap)
		} else if ev.Rune() == 'm' {
			p.MoveRight(1, worldMap)
		}
	}
}

// TODO: fucking god damnit this is messy

//func intAbs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}

// Some fancy version of bresenhams
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// http://members.chello.at/~easyfilter/Bresenham.pdf

// Returns the distance, and the value of the map it intersected.
// If no intersection, map value == 0
// func castRay(screen tcell.Screen, WorldMap [][]int, x0, y0, x1, y1 int, style tcell.Style) *linalg.Vec2 {
// 	dx := intAbs(x1 - x0)
// 	dy := -intAbs(y1 - y0)
// 	sx := -1
// 	if x0 < x1 {
// 		sx = 1
// 	}
// 	sy := -1
// 	if y0 < y1 {
// 		sy = 1
// 	}
// 	er := dx + dy

// 	x := x0
// 	y := y0

// 	//rayLen := math.Sqrt(math.Pow(float64(x1-x0), 2) + math.Pow(float64(y1-y0), 2))
// 	//dist := rayLen
// 	intersectionPoint := &linalg.Vec2{}
// 	for {
// 		setContentEqualWidth(screen, x, y, ' ', nil, style)

// 		// Get dists for drawing 3D scene
// 		if world.WorldMap[y][x] != 0 {
// 			//dist = math.Sqrt(math.Pow(float64(x-x0), 2) + math.Pow(float64(y-y0), 2))
// 			intersectionPoint = &linalg.Vec2{float64(x), float64(y)}
// 			break
// 		}

// 		if x == x1 && y == y1 {
// 			break
// 		}

// 		e2 := 2 * er
// 		if e2 >= dy {
// 			if x == x1 {
// 				break
// 			}
// 			er = er + dy
// 			x = x + sx
// 		}

// 		if e2 <= dx {
// 			if y == y1 {
// 				break
// 			}
// 			er = er + dx
// 			y = y + sy
// 		}
// 	}
// 	return intersectionPoint
// }
