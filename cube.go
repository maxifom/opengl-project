package main

import "github.com/go-gl/mathgl/mgl32"

type Cube struct {
	position     mgl32.Vec3
	vertices     []float32
	indices      []uint32
	rotation     float32
	rotationAxes mgl32.Vec3
}

func NewCube(a float32, pos mgl32.Vec3, rotation float32, rotationAxes mgl32.Vec3) *Cube {
	vertices := []float32{
		//  X, Y, Z, U, V
		// Bottom
		-a, -a, -a, 0.0, 0.0,
		a, -a, -a, 1.0, 0.0,
		-a, -a, a, 0.0, 1.0,
		a, -a, -a, 1.0, 0.0,
		a, -a, a, 1.0, 1.0,
		-a, -a, a, 0.0, 1.0,

		// Top
		-a, a, -a, 0.0, 0.0,
		-a, a, a, 0.0, 1.0,
		a, a, -a, 1.0, 0.0,
		a, a, -a, 1.0, 0.0,
		-a, a, a, 0.0, 1.0,
		a, a, a, 1.0, 1.0,

		// Front
		-a, -a, a, 1.0, 0.0,
		a, -a, a, 0.0, 0.0,
		-a, a, a, 1.0, 1.0,
		a, -a, a, 0.0, 0.0,
		a, a, a, 0.0, 1.0,
		-a, a, a, 1.0, 1.0,

		// Back
		-a, -a, -a, 0.0, 0.0,
		-a, a, -a, 0.0, 1.0,
		a, -a, -a, 1.0, 0.0,
		a, -a, -a, 1.0, 0.0,
		-a, a, -a, 0.0, 1.0,
		a, a, -a, 1.0, 1.0,

		// Left
		-a, -a, a, 0.0, 1.0,
		-a, a, -a, 1.0, 0.0,
		-a, -a, -a, 0.0, 0.0,
		-a, -a, a, 0.0, 1.0,
		-a, a, a, 1.0, 1.0,
		-a, a, -a, 1.0, 0.0,

		// Right
		a, -a, a, 1.0, 1.0,
		a, -a, -a, 1.0, 0.0,
		a, a, -a, 0.0, 0.0,
		a, -a, a, 1.0, 1.0,
		a, a, -a, 0.0, 0.0,
		a, a, a, 0.0, 1.0,
	}

	var indices []uint32

	for i := range vertices {
		indices = append(indices, uint32(i))
	}

	return &Cube{vertices: vertices, position: pos, rotation: rotation, indices: indices, rotationAxes: rotationAxes}
}

func (c *Cube) Vertices() []float32 {
	return c.vertices
}

func (c *Cube) Position() mgl32.Vec3 {
	return c.position
}
func (c *Cube) SetPosition(pos mgl32.Vec3) {
	c.position = pos
}

func (c *Cube) Rotation() float32 {
	return c.rotation
}

func (c *Cube) SetRotation(f float32) {
	c.rotation = f
}

func (c *Cube) Indices() []uint32 {
	return c.indices
}

func (c *Cube) RotationAxes() mgl32.Vec3 {
	return c.rotationAxes
}
