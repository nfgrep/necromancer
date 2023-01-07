package player

import (
	"math"

	"github.com/gdamore/tcell/v2"
	"github.com/nfgrep/necromancer/gfx"
	"github.com/nfgrep/necromancer/linalg"
	"github.com/nfgrep/necromancer/world"
)

type Player struct {
	Position *linalg.Vec2
	// Fwd      *linalg.Vec2
	Rotation *float64
	ViewRays []*ViewRay
	//	rayCount    int
	//	rayAngleGap float64
	ViewLen float64
}

func (p *Player) Draw(screen tcell.Screen, style tcell.Style) {
	// Draw the player
	gfx.SetContentEqualWidth(screen, int(p.Position.X), int(p.Position.Y), ' ', nil, style)

	// Some debug
	//gfx.DrawDebugText(0, fmt.Sprintf("player Fwd: %v", p.Fwd))
	//drawDebugText(screen, fmt.Sprintf("player dir: %v", player.Fwd))
}

// Returns an array of the distances to the walls for each ray
// TODO: the screen is just to draw debug rays, maybe remove?
func (p *Player) CastViewRays(worldMap world.Map, screen tcell.Screen, rayStyle tcell.Style) []float64 {
	dists := []float64{}
	for i, viewRay := range p.ViewRays {
		// Get endpoints for ray
		//rEnd := ray.GetEnd(player.ViewLen)
		// rx1 := player.pos.X + math.Cos(ray.rot+player.rot)*float64(player.ViewLen)
		// ry1 := player.pos.Y + math.Sin(ray.rot+player.rot)*float64(player.ViewLen)
		//intersect := castRay(screen, world.WorldMap, int(player.x), int(player.y), int(rx1), int(ry1), player.rayStyle)

		// Define a function that runs on each iteration of the raycast, and checks if we've hit a wall
		intersects := func(x, y float64) bool {
			// Ignore negative and out-of-bounds values
			if x < 0 || y < 0 || x > float64(worldMap.Width()) || y > float64(worldMap.Height()) {
				return false
			}
			gfx.SetContentEqualWidth(screen, int(x), int(y), ' ', nil, rayStyle) // Draw the ray in the top-down map view

			return worldMap.WallExistsAt(x, y)
		}

		// TODO: combine the offset witht the base ray and cast that.
		intersect := viewRay.Cast(p.ViewLen, intersects, i)
		//gfx.DrawDebugText(i, fmt.Sprintf("ray: %v, intersect: %v", i, intersect))

		rayDist := math.Sqrt(math.Pow(float64(p.Position.X-intersect.X), 2) + math.Pow(float64(p.Position.Y-intersect.Y), 2))

		dists = append(dists, rayDist)
	}
	return dists
}

// Moves player if possible, i.e. if there's no wall in the way
func (p *Player) move(dx, dy float64, worldMap world.Map) {
	newX := p.Position.X + dx
	newY := p.Position.Y + dy
	if !worldMap.WallExistsAt(newX, newY) {
		p.Position.X = newX
		p.Position.Y = newY
	}
}

func (p *Player) MoveFwd(dist float64, worldMap world.Map) {
	// dx := p.Fwd.X * dist
	// dy := p.Fwd.Y * dist

	// Get the forward vector and multiply by the distance
	fwdVec := p.ForwardVec().Mul(dist)
	// Move the player
	p.move(fwdVec.X, fwdVec.Y, worldMap)
	// p.move(dx, dy, worldMap)
}

func (p *Player) MoveLeft(dist float64, worldMap world.Map) {
	// // Get the left-vector
	// // Flip the x and y, and negate the x
	// dx := p.Fwd.Y * dist
	// dy := -p.Fwd.X * dist
	// p.move(dx, dy, worldMap)

	// Get the right vector, negate it to get left vector, and multiply by the distance
	leftVec := p.RightVec().Mul(-dist)
	// Move the player
	p.move(leftVec.X, leftVec.Y, worldMap)
}

func (p *Player) MoveRight(dist float64, worldMap world.Map) {
	// // Get the right-vector
	// // Flip the x and y, and negate the y
	// dx := -p.Fwd.Y * dist
	// dy := p.Fwd.X * dist

	// Get the right vector and multiply by the distance
	rightVec := p.RightVec().Mul(dist)
	// Move the player
	p.move(rightVec.X, rightVec.Y, worldMap)
	//p.move(dx, dy, worldMap)
}

func (p *Player) MoveBack(dist float64, worldMap world.Map) {
	// dx := -p.Fwd.X * dist
	// dy := -p.Fwd.Y * dist

	// Get the forward vector, negate it to get back vector, and multiply by the distance
	backVec := p.ForwardVec().Mul(-dist)
	// Move the player
	p.move(backVec.X, backVec.Y, worldMap)
	// p.move(dx, dy, worldMap)
}

func (p *Player) RightVec() *linalg.Vec2 {
	// Subtract pi/2 (90 degrees) from the rotation
	rightRot := rotWrap(*p.Rotation + (math.Pi / 2))
	// Get the right vector, make sure it's normalized
	rightVec := linalg.Vec2FromAngle(rightRot).Normalized()
	return rightVec
}

func (p *Player) ForwardVec() *linalg.Vec2 {
	// Get the forward vector, make sure it's normalized
	forwardVec := linalg.Vec2FromAngle(*p.Rotation).Normalized()
	return forwardVec
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

func (p *Player) Rotate(rad float64) {
	// // Get existing rotation from fwd vector
	// currRot := p.Fwd.Angle()
	// // Add the rotation
	// newRot := currRot + rad
	// // Make sure to wrap it
	// rot := rotWrap(newRot)
	// // Set our new fwd vector
	// newFwd := linalg.Vec2FromAngle(rot).Normalized()
	// p.Fwd.X = newFwd.X
	// p.Fwd.Y = newFwd.Y

	// Add the rotation
	newRot := *p.Rotation + rad
	// Make sure to wrap it
	rot := rotWrap(newRot)
	// Set our new rotation
	*p.Rotation = rot // TODO: this feels icky, is it?
}

// TODO: take in basis ray
func GenerateRays(count int, fov float64, origin *linalg.Vec2, playerRot *float64) []*ViewRay {
	rays := []*ViewRay{}
	rayGap := fov / float64(count)
	halfCount := count / 2

	// if even
	if count%2 == 0 {
		for i := -halfCount; i < count; i++ {
			rayRot := rayGap * float64(i)
			// ray := &gfx.Ray{Origin: origin, Direction: fwd}                      // This one changes dynamically
			// viewRay := &ViewRay{Ray: ray, offset: *linalg.Vec2FromAngle(rayRot)} // This one stores a ref to former and a static offset
			viewRay := &ViewRay{
				Ray: &gfx.Ray{
					Origin:   origin,
					Rotation: playerRot,
				},
				offsetAngle: rayRot,
			}
			rays = append(rays, viewRay)
		}
	} else {
		for i := 0; i < halfCount; i++ {
			rayRot := rayGap * float64(i)
			// posRay := &gfx.Ray{Origin: origin, Direction: fwd}
			// posViewRay := &ViewRay{Ray: posRay, offset: *linalg.Vec2FromAngle(rayRot)}
			// negRay := &gfx.Ray{Origin: origin, Direction: fwd}
			// negViewRay := &ViewRay{Ray: negRay, offset: *linalg.Vec2FromAngle(-rayRot)}
			posViewRay := &ViewRay{
				Ray: &gfx.Ray{
					Origin:   origin,
					Rotation: playerRot,
				},
				offsetAngle: rayRot,
			}
			negViewRay := &ViewRay{
				Ray: &gfx.Ray{
					Origin:   origin,
					Rotation: playerRot,
				},
				offsetAngle: -rayRot,
			}
			// Prepend the negative rays, so that the rays are ordered from left to right
			rays = append([]*ViewRay{negViewRay}, rays...)
			// Then append the positive ray
			rays = append(rays, posViewRay)
		}
	}

	return rays
}
