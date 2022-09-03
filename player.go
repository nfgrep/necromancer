package main

import "math"

type Ray struct {
	rot float64 // Relative to player rot
}

type Player struct {
	x    float64
	y    float64
	rot  float64
	rays []*Ray
	//	rayCount    int
	//	rayAngleGap float64
	viewLen int
}

// Moves player if possible, i.e. if there's no wall in the way
func move(dx, dy float64, worldMap [][]int) {
	newX := player.x + dx
	newY := player.y + dy
	newIntX := int(newX) // _Should_ just truncate the decimal off.
	newIntY := int(newY)
	if worldMap[newIntY][newIntX] == 0 {
		player.x = newX
		player.y = newY
	}
}

func (p *Player) moveFwd(dist float64, worldMap [][]int) {
	// Calc fwd vector [dx, dy]
	dx := math.Cos(p.rot) * dist
	dy := math.Sin(p.rot) * dist

	move(dx, dy, worldMap)
}

func (p *Player) moveLeft(dist float64, worldMap [][]int) {
	// TODO: I'm not totally sure how this works
	dx := math.Sin(p.rot) * dist
	dy := -math.Cos(p.rot) * dist

	move(dx, dy, worldMap)
}

func (p *Player) moveRight(dist float64, worldMap [][]int) {
	// TODO: I'm not totally sure how this works
	dx := -math.Sin(p.rot) * dist
	dy := math.Cos(p.rot) * dist

	move(dx, dy, worldMap)
}

func (p *Player) moveBack(dist float64, worldMap [][]int) {
	// TODO: I'm not totally sure how this works
	dx := math.Cos(p.rot) * -dist
	dy := math.Sin(p.rot) * -dist

	move(dx, dy, worldMap)
}

// Takes rot in radians, returns the value between 0, 2pi
func rotWrap(rot float64) float64 {
	// Hopefully this is pass-by-val
	if rot >= (2 * math.Pi) {
		rot -= (2 * math.Pi)
	}
	if rot <= 0 {
		rot += (2 * math.Pi)
	}
	return rot
}

func (p *Player) rotate(rad float64) {
	p.rot += rad
	p.rot = rotWrap(p.rot)
}

//func (p *Player) generateRays() {
//	for i := 0; i < p.rayCount; i++ {
//
//	}
//}
