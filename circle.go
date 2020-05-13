package main

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type Circle struct {
	vertices []float32
	radius   float32
	position mgl32.Vec3
	rotation float32
}

func NewCircle(numberOfSectors int, radius float32, position mgl32.Vec3, rotation float32) *Circle {
	deltaAngle := 2 * math.Pi / float64(numberOfSectors)
	var vertices = make([]float32, 0, numberOfSectors*5*3)
	var lastVertices []float32
	for i := 0.0; i < 360; i += 360.0 / float64(numberOfSectors) {
		vertices = append(vertices, position.X(), position.Y(), position.Z(), 0, 0)
		rad := float64(mgl32.DegToRad(float32(i)))
		vertices = append(vertices, lastVertices...)
		lastVertices = []float32{
			float32(math.Cos(rad) * float64(radius)),
			float32(math.Sin(rad) * float64(radius)),
			position.Z(),
			float32((math.Cos(deltaAngle*i) + 1.0) * 0.5),
			float32((math.Sin(deltaAngle*i) + 1.0) * 0.5)}
		vertices = append(vertices, lastVertices...)
	}

	return &Circle{position: position, radius: radius, vertices: vertices, rotation: rotation}
}

func (c *Circle) Vertices() []float32 {
	return c.vertices
}

func (c *Circle) Position() mgl32.Vec3 {
	return c.position
}

func (c *Circle) SetPosition(pos mgl32.Vec3) {
	c.position = pos
}

func (c *Circle) Rotation() float32 {
	return c.rotation
}

func (c *Circle) SetRotation(f float32) {
	c.rotation = f
}
