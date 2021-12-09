package main

import "math"

type Vector2 struct
{
	x, y float64
}

type Vector3 struct
{
	x, y, z float64
}

func (v Vector2) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2));
}

func (v Vector3) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2));
}
