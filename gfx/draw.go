package gfx

import (
	"github.com/gdamore/tcell/v2"
)

var debugScreen tcell.Screen

func SetDebugScreen(s tcell.Screen) {
	debugScreen = s
}

// This is the current height of the map
var debugTextYOffset int = 30
var debugTextWidth int = 200
var debugStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

func DrawDebugText(offset int, text string) {
	drawText(debugScreen, 0, debugTextYOffset+offset, debugTextWidth, debugTextYOffset+offset, debugStyle, text)
}

// func DrawScene(player *player.Player, worldMap *world.Map, screen tcell.Screen, style tcell.Style) {
// 	dists := player.CastViewRays()
// 	// TODO: make these constants?
// 	maxHeight := 50
// 	horizonYPos := 20
// 	DrawBarsForDists(screen, dists, player.ViewLen, maxHeight, horizonYPos, worldMap.Width(), style)
// }

func DrawBarsForDists(screen tcell.Screen, dists []float64, viewLen float64, maxHeight, horizonYPos, screenXOffset int, style tcell.Style, textureSlices [][]tcell.Style) {
	for i, dist := range dists {
		// Dont draw anything if we didn't intersect anything
		if dist >= viewLen {
			continue
		}
		// TODO: is this a problem?
		// Projected onto the flat view/camera plane
		// Minimizes fish-eye effect a bit
		//correctedRayDist := rayDist * math.Cos(ray.rot)

		// Get the height of the bar to draw
		barHeight := maxHeight - int(dist)
		if barHeight == 0 {
			continue
		}

		// --- Whack texturing shit

		// worldRayRot := player.rot + ray.rot

		// var texXF float64
		// displacementRayX := math.Abs(intersect.X - player.pos.X)
		// displacementRayY := math.Abs(intersect.X - player.pos.Y)
		// if displacementRayX > displacementRayY {
		// 	texXF = intersect.Y
		// 	if !(worldRayRot < (3*math.Pi)/2 && worldRayRot > math.Pi/2) { // If we're in the right-half
		// 		texXF = float64(len(world.WorldMap[0])) - texXF
		// 	}
		// } else {
		// 	texXF = intersect.X
		// 	if worldRayRot > math.Pi { // If we're in the bottom half
		// 		texXF = float64(len(world.WorldMap[0])) - texXF
		// 	}
		// }

		//drawDebugText(screen, fmt.Sprintf("dispRayX"))
		//if worldRayRot > (math.Pi/2) && worldRayRot < (math.Pi+(math.Pi/2)) {
		//	texXF = intersect.y
		//} else {
		//	texXF = intersect.x
		//}

		// --- Textured bar stuff
		// tex := wallTexture
		// texX := int(texXF) % len(tex[0])
		// texSlice := getTexSlice(wallTexture, texX)
		// filteredTexSlice := filterTexSlice(texSlice, barHeight)

		//barStyle := styleMap[world.WorldMap[int(intersect.y)][int(intersect.x)]]
		//drawDebugText(screen, fmt.Sprintf("barHeight: %v", barHeight))
		textureSlice := textureSlices[i]
		DrawBarWithTexture(screen, i+screenXOffset, horizonYPos, barHeight, style, textureSlice)
		// DrawBarWithColor(screen, i+screenXOffset, horizonYPos, barHeight, style)
		// -- end Draw bar

		//dists = append(dists, int(rayDist))
		//drawText(screen, 2, i+30, 70, i+35, style, fmt.Sprintf("ray: %v, ray.rot: %v, rx1: %v, ry1: %v, rayDist: %v", i, ray.rot, rx1, ry1, rayDist))
	}
}

func DrawBarWithTexture(s tcell.Screen, screenPosX, screenPosY, height int, style tcell.Style, textureSlice []tcell.Style) {
	ytop := screenPosY - (height / 2)

	scaledSlice := scaleTextureSlice(textureSlice, height)

	for i := 0; i < height; i++ {
		y := i + ytop
		SetContentEqualWidth(s, screenPosX, y, ' ', nil, scaledSlice[i])
	}
}

func scaleTextureSlice(textureSlice []tcell.Style, height int) []tcell.Style {
	textureWidth := len(textureSlice)

	scaledSlice := make([]tcell.Style, height)

	for x := 0; x < height; x++ {
		texX := (x * textureWidth) / height
		scaledSlice[x] = textureSlice[texX]
	}

	return scaledSlice
}

func DrawBarWithColor(s tcell.Screen, screenPosX, screenPosY, height int, style tcell.Style) {
	ytop := screenPosY - (height / 2)
	for i := 0; i < height; i++ {
		y := i + ytop
		SetContentEqualWidth(s, screenPosX, y, ' ', nil, style)
	}
}

func SetContentEqualWidth(screen tcell.Screen, x int, y int, primary rune, combining []rune, style tcell.Style) {
	screen.SetContent((x*2)+1, y, primary, combining, style)
	screen.SetContent(x*2, y, primary, combining, style)
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
