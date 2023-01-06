package gfx

import (
	"math"

	"github.com/nfgrep/necromancer/linalg"
)

type Ray struct {
	Origin   *linalg.Vec2
	Rotation *float64
}

func (r *Ray) Cast(dist float64, intersects func(x, y float64) bool, rayIdx int) *linalg.Vec2 {
	rayTip := r.Tip(dist)
	x0 := r.Origin.X
	x1 := rayTip.X
	y0 := r.Origin.Y
	y1 := rayTip.Y
	//DrawDebugText(rayIdx, fmt.Sprintf("ray: %v, x0: %v, y0: %v, x1: %v, y1: %v", rayIdx, x0, y0, x1, y1))
	return castRay(x0, y0, x1, y1, intersects, rayIdx)
}

func (r *Ray) Tip(length float64) *linalg.Vec2 {
	return linalg.Vec2FromAngle(*r.Rotation).Mul(length).Add(*r.Origin)
	// return &linalg.Vec2{
	// 	X: r.Origin.X + (r.Direction.X * length),
	// 	Y: r.Origin.Y + (r.Direction.Y * length),
	// }
}

func castRay(x0, y0, x1, y1 float64, intersects func(x, y float64) bool, rayIdx int) *linalg.Vec2 {
	if math.Abs(y1-y0) < math.Abs(x1-x0) {
		if x0 > x1 {
			return castRayLow(x1, y1, x0, y0, intersects, rayIdx)
		} else {
			return castRayLow(x0, y0, x1, y1, intersects, rayIdx)
		}
	} else {
		if y0 > y1 {
			return castRayHigh(x1, y1, x0, y0, intersects, rayIdx)
		} else {
			return castRayHigh(x0, y0, x1, y1, intersects, rayIdx)
		}
	}
}

func castRayLow(x0, y0, x1, y1 float64, intersects func(x, y float64) bool, rayIdx int) *linalg.Vec2 {
	dx := x1 - x0
	dy := y1 - y0
	yi := 1.0
	if dy < 0 {
		yi = -1.0
		dy = -dy
	}
	D := (2 * dy) - dx
	y := y0
	x := x0
	//DrawDebugText(rayIdx, fmt.Sprintf("x: %v, y: %v", x, y))
	for x <= x1 {
		if intersects(x, y) {
			return &linalg.Vec2{X: x, Y: y}
		}
		if D > 0 {
			y = y + yi
			D = D + (2 * (dy - dx))
		} else {
			D = D + 2*dy
		}
		x += 1
	}
	return &linalg.Vec2{X: x, Y: y}
}

func castRayHigh(x0, y0, x1, y1 float64, intersects func(x, y float64) bool, rayIdx int) *linalg.Vec2 {
	dx := x1 - x0
	dy := y1 - y0
	xi := 1.0
	if dx < 0 {
		xi = -1.0
		dx = -dx
	}
	D := (2 * dx) - dy
	x := x0
	y := y0
	//DrawDebugText(rayIdx, fmt.Sprintf("x: %v, y: %v", x, y))
	for y <= y1 {
		if intersects(x, y) {
			return &linalg.Vec2{X: x, Y: y}
		}
		if D > 0 {
			x = x + xi
			D = D + (2 * (dx - dy))
		} else {
			D = D + 2*dx
		}
		y += 1
	}
	return &linalg.Vec2{X: x, Y: y}
}
