package main

import "github.com/gdamore/tcell/v2"

type Texture [][]tcell.Style

// Colours for each parts of the wall texture
var y = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorYellow)
var b = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlue)
var r = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)
var p = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPink)
var checkerTexture = Texture{
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
	{y, b, y, b, y, b, y, b, y, b, y, r},
	{b, y, b, y, b, y, b, y, b, y, b, y},
}

var g = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGreen)
var wallTexture = Texture{
	{y, y, y, y, y, y, r, r, r, r, r, r},
	{y, y, y, y, y, y, r, r, r, r, r, r},
	{y, y, y, y, y, y, r, r, r, r, r, r},
	{y, y, y, y, y, y, r, r, r, r, r, r},
	{y, y, y, p, y, y, r, r, r, r, r, r},
	{y, y, p, p, y, y, r, r, r, r, r, r},
	{y, p, p, p, y, y, r, r, r, r, r, r},
	{y, p, p, p, y, y, r, r, r, r, r, r},
	{y, y, p, p, y, y, r, r, r, r, r, r},
	{y, y, y, p, y, y, r, r, r, r, r, r},
	{y, y, y, y, y, y, r, r, r, r, r, r},
	{y, y, y, y, y, y, r, r, r, r, r, r},
	{y, y, y, y, y, y, r, r, r, r, r, r},
	{y, y, y, y, y, y, r, r, r, r, r, r},
	{y, y, y, y, y, y, r, r, r, r, r, r},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
	{g, g, g, g, g, g, b, b, b, b, b, b},
}

var smallTexture = Texture{
	{p, p, r, r, g, g, g, g},
	{p, p, r, r, g, g, g, g},
	{r, r, r, r, g, g, g, g},
	{r, r, r, r, g, g, g, g},
	{b, b, b, b, y, y, y, y},
	{b, b, b, b, y, y, y, y},
	{b, b, b, b, y, y, y, y},
	{b, b, b, b, y, y, y, y},
}

// WorldMap values to textures
//var textures = map[int]Texture{
//	1: wallTexture,
//}

// TODO: make this more automagic/generated?
// Maps x, y in world coords to x, y in wallTex coords
//var wallTexMap = map[Point]Point{
//	{x: 1, y: 1}: {}
//}

// Returns a map from x, y in world coords to the index of the column of the texture to use
//func generateWallTexMap(worldMap [][]int) map[Point]int {
//	textureIndices := make(map[int]int) // Maps map values (i.e. "1") to the indices for the texture for the map value
//	texMap := make(map[Point]int)
//	for y, row := range worldMap {
//		for x, mapVal := range row {
//			texForVal := textures[mapVal]
//			idx := textureIndices[mapVal] // Indexing an empty map returns 0, not nil
//			texWidth := len(texForVal[0]) // Assumes equal width all the way down the texture
//			texMap[Point{float64(x), float64(y)}] = idx % texWidth
//		}
//	}
//	return texMap
//}

// Takes a verticle slice of a texture, and returns the magnified or minified texture slice of len height
//func minifyTextureSlice(s tcell.Screen, inTex []tcell.Style, height int) []tcell.Style {
//	outTex := []tcell.Style{}
//
//	quotient := float64(len(inTex)) / float64(height)
//	// Only do minification. If the size of the tex is smaller than our output
//	if quotient < 1.0 {
//		return inTex
//	}
//
//	window_start := 0.0
//
//	for i := 1; i < height-1; i++ {
//		window_end := float64(i) * quotient
//
//		prefix_idx := int(window_start)   // 1.4 -> [ ,X, ...]
//		suffix_idx := int(window_end) + 1 // 2.8 -> [ , , ,X, ...]
//
//		whole_start_idx := prefix_idx + 1
//		whole_end_idx := suffix_idx - 1
//
//		prefix_colour := inTex[prefix_idx]
//		suffix_colour := inTex[suffix_idx]
//		whole_colours := inTex[whole_start_idx:whole_end_idx]
//
//		prefix_colour_weight := 1.0 - (window_start - math.Trunc(window_start)) // 1.0 - (1.4 - 1.0) = 0.6
//		suffix_colour_weight := window_end - math.Trunc(window_end)
//		whole_colour_weight := 1.0
//
//		//mixColours(prefix_colour, prefix_colour_weight, whole_colours, whole_colour_weight, suffix_colour, suffix_colour_weight)
//		//fmt.Println(prefix_colour, prefix_colour_weight, whole_colours, whole_colour_weight, suffix_colour, suffix_colour_weight)
//		drawDebugText(s, rayStyle, fmt.Sprint(prefix_colour, prefix_colour_weight, whole_colours, whole_colour_weight, suffix_colour, suffix_colour_weight))
//
//		window_start = window_end
//	}
//
//	return outTex
//}
