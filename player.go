package main

import (
	"math"

	"github.com/gdamore/tcell/v2"
)

type Ray struct {
	rot float64 // Relative to player rot
}

type Player struct {
	x        float64
	y        float64
	rot      float64
	rays     []*Ray
	rayStyle tcell.Style
	//	rayCount    int
	//	rayAngleGap float64
	viewLen int
}

var player = &Player{
	x: 2, y: 2,
	rot:      math.Pi / 2,
	rays:     generateRays(30, 1.0),
	rayStyle: tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed),
	viewLen:  70,
}

func Init() {
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

func generateRays(count int, fov float64) []*Ray {
	rays := []*Ray{}
	rayGap := fov / float64(count)
	halfCount := count / 2

	// if even
	if count%2 == 0 {
		for i := -halfCount; i < count; i++ {
			rays = append(rays, &Ray{rot: rayGap * float64(i)})
		}
		//// negative half of the view cone
		//for i := halfCount; i > -2; i-- {
		//	rayRot := rayGap * float64(i)
		//	rays = append(rays, &Ray{rot: -rayRot})
		//}

		//centerRaysRot := rayGap / 2.0
		////centerNegRayRot := -(centerRaysRot * float64(halfCount-1))
		////centerPosRayRot := (centerRaysRot * float64(halfCount))
		//rays = append(rays, &Ray{rot: -centerRaysRot * })
		//rays = append(rays, &Ray{rot: centerRaysRot})

		//// positive half of the view cone
		//for i := 2; i <= halfCount; i++ {
		//	rayRot := rayGap * float64(i)
		//	rays = append(rays, &Ray{rot: rayRot})
		//}
	} else {
		for i := 0; i < halfCount; i++ {
			rayRot := rayGap * float64(i)
			rays = append(rays, &Ray{rot: rayRot}, &Ray{rot: -rayRot})
		}
	}

	return rays
}
