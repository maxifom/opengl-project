package main

import "github.com/go-gl/mathgl/mgl32"

type Parallelepiped struct {
	vertices []float32
	indices  []uint32
	position mgl32.Vec3
	rotation float32
}

func NewParallelepiped(a, b, c float32, pos mgl32.Vec3, rotation float32) *Parallelepiped {
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

	var indices []uint32
	for i := range vertices {
		indices = append(indices, uint32(i))
	}

	return &Parallelepiped{vertices: vertices, position: pos, rotation: rotation, indices: indices}
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

func (c *Parallelepiped) Rotation() float32 {
	return c.rotation
}

func (c *Parallelepiped) SetRotation(f float32) {
	c.rotation = f
}

func (c *Parallelepiped) Indices() []uint32 {
	return c.indices
}
