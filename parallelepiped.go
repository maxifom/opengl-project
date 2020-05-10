package main

import "github.com/go-gl/mathgl/mgl32"

type Parallelepiped struct {
	vertices []float32
	position mgl32.Vec3
}

func NewParallelepiped(a, b, c float32, pos mgl32.Vec3) *Parallelepiped {
	vertices := []float32{
		//  X, Y, Z, U, V
		// Bottom
		-a, -b, -c, 0.0, 0.0,
		a, -b, -c, 1.0, 0.0,
		-a, -b, c, 0.0, 1.0,
		a, -b, -c, 1.0, 0.0,
		a, -b, c, 1.0, 1.0,
		-a, -b, c, 0.0, 1.0,

		// Top
		-a, b, -c, 0.0, 0.0,
		-a, b, c, 0.0, 1.0,
		a, b, -c, 1.0, 0.0,
		a, b, -c, 1.0, 0.0,
		-a, b, c, 0.0, 1.0,
		a, b, c, 1.0, 1.0,

		// Front
		-a, -b, c, 1.0, 0.0,
		a, -b, c, 0.0, 0.0,
		-a, b, c, 1.0, 1.0,
		a, -b, c, 0.0, 0.0,
		a, b, c, 0.0, 1.0,
		-a, b, c, 1.0, 1.0,

		// Back
		-a, -b, -c, 0.0, 0.0,
		-a, b, -c, 0.0, 1.0,
		a, -b, -c, 1.0, 0.0,
		a, -b, -c, 1.0, 0.0,
		-a, b, -c, 0.0, 1.0,
		a, b, -c, 1.0, 1.0,

		// Left
		-a, -b, c, 0.0, 1.0,
		-a, b, -c, 1.0, 0.0,
		-a, -b, -c, 0.0, 0.0,
		-a, -b, c, 0.0, 1.0,
		-a, b, c, 1.0, 1.0,
		-a, b, -c, 1.0, 0.0,

		// Right
		a, -b, c, 1.0, 1.0,
		a, -b, -c, 1.0, 0.0,
		a, b, -c, 0.0, 0.0,
		a, -b, c, 1.0, 1.0,
		a, b, -c, 0.0, 0.0,
		a, b, c, 0.0, 1.0,
	}

	return &Parallelepiped{vertices: vertices, position: pos}
}

func (c *Parallelepiped) Vertices() []float32 {
	return c.vertices
}

func (c *Parallelepiped) Position() mgl32.Vec3 {
	return c.position
}

func (c *Parallelepiped) SetPosition(pos mgl32.Vec3) {
	c.position = pos
}
