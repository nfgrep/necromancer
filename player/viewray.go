package player

import (
	"github.com/nfgrep/necromancer/gfx"
	"github.com/nfgrep/necromancer/linalg"
)

// The plan is to just have the ray in this datatype be a pointer to the fwd vec, and add the offset to that at runtime.
type ViewRay struct {
	*gfx.Ray
	//offset linalg.Vec2
	offsetAngle float64
}

// TODO: this is making alot of garbage, should be seperate array of offsets or soemthing?
// Combines underlying Ray with offset
func (v *ViewRay) Solve() gfx.Ray {
	newRotation := *v.Ray.Rotation + v.offsetAngle
	return gfx.Ray{
		Origin: v.Ray.Origin,
		//Direction: v.Ray.Direction.Add(v.offset).Normalized(),
		Rotation: &newRotation,
	}
}

func (v *ViewRay) Cast(dist float64, intersects func(x, y float64) bool, rayIdx int) *linalg.Vec2 {
	rayToCast := v.Solve()
	return rayToCast.Cast(dist, intersects, rayIdx)
}
