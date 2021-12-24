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

// Adds the current vector instance to the 'other'
func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{v.x + other.x, v.y + other.y}
}

// Subtracts the 'other' vector from the current instance
func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{v.x - other.x, v.y - other.y}
}

func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{v.x + other.x, v.y + other.y, v.z + other.z}
}

// Subtracts the 'other' vector from the current instance
func (v Vector3) Sub(other Vector3) Vector3 {
	return Vector3{v.x - other.x, v.y - other.y, v.z - other.z}
}
