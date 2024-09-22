package linalg

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v *Vec2) Add(v2 Vec2) *Vec2 {
	return &Vec2{v.X + v2.X, v.Y + v2.Y}
}

func (v *Vec2) Sub(v2 Vec2) *Vec2 {
	return &Vec2{v.X - v2.X, v.Y - v2.Y}
}

func (v *Vec2) Mul(scalar float64) *Vec2 {
	return &Vec2{v.X * scalar, v.Y * scalar}
}

func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vec2) Normalized() *Vec2 {
	mag := v.Magnitude()
	return &Vec2{v.X / mag, v.Y / mag}
}

func (v *Vec2) Dot(v2 Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

type Vec3 struct {
	x float64
	y float64
	z float64
}

// TODO: I think this is is specifically the angle from (1,0)
func Vec2FromAngle(angle float64) *Vec2 {
	return &Vec2{math.Cos(angle), math.Sin(angle)}
}

// Retorns the angle in radians
func (v *Vec2) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}
