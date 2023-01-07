package world

import (
	"github.com/nfgrep/necromancer/gfx"
	"github.com/nfgrep/necromancer/linalg"
)

// So the idea is basiclaly to to world -> wall -> texture
// We get the wall from the number in the map, we then take that walls position/dimensions to get the local coords from the world coords
// We then use those local coords, and map them to tex coords, which we then use to index into the texture

type Wall struct {
	texture    gfx.Texture
	position   linalg.Vec2 // in world space
	dimensions linalg.Vec2 // x = width, y = height

	//p0      Vec3 // minima, in world-space (probably upper left?)
	//p1      Vec3 // maxima, in world-space (probbaly lower right?)
	//normal  Vec3
}

func (w *Wall) worldToLocal(point linalg.Vec2) linalg.Vec2 {
	x := point.X - w.position.X
	y := point.X - w.position.X
	return linalg.Vec2{X: x, Y: y}
}

//var wallFlavours = map[int]Wall{
//	1: {texture: checkerTexture, height: 10},
//	2: {texture: wallTexture, height: 12},
//}
