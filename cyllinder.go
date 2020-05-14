package main

import "github.com/go-gl/mathgl/mgl32"

type Cyllinder struct {
	vertices []float32
	indices  []uint32
	rotation float32
	position mgl32.Vec3
}

func NewCyllinder() *Cyllinder {
	var vertices []float32
	var indices []uint32

	for i := range vertices {
		indices = append(indices, uint32(i))
	}

	return &Cyllinder{
		vertices: vertices,
		indices:  indices,
		rotation: 0,
		position: mgl32.Vec3{},
	}
}

func (c *Cyllinder) Vertices() []float32 {
	return c.vertices
}

func (c *Cyllinder) Indices() []uint32 {
	return c.indices
}

func (c *Cyllinder) Position() mgl32.Vec3 {
	return c.position
}

func (c *Cyllinder) SetPosition(vec3 mgl32.Vec3) {
	c.position = vec3
}

func (c *Cyllinder) Rotation() float32 {
	return c.rotation
}

func (c *Cyllinder) SetRotation(f float32) {
	c.rotation = f
}
