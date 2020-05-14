package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type Cyllinder struct {
	vertices     []float32
	indices      []uint32
	rotation     float32
	position     mgl32.Vec3
	rotationAxes mgl32.Vec3
}

func NewCyllinder() *Cyllinder {
	var vertices []float32
	var indices []uint32

	var H = 2.0
	var R = 1.0
	var dh = 0.001
	var dl = 1
	for h := -0.5; h < 0.5; h += dh {
		for l := 0; l < 360; l += dl {

			//x0,y0,z0,u0,v0
			vertices = append(vertices, float32(R*math.Sin(float64(l))))
			vertices = append(vertices, float32(R*math.Cos(float64(l))))
			vertices = append(vertices, float32(h*H))
			vertices = append(vertices, 1, 1)

			//x1,y1,z1,u1,v1
			vertices = append(vertices, float32(R*math.Sin(float64(l))))
			vertices = append(vertices, float32(R*math.Cos(float64(l))))
			vertices = append(vertices, float32((h+dh)*H))
			vertices = append(vertices, 1, 1)

			//x2,y2,z2,u2,v2
			vertices = append(vertices, float32(R*math.Sin(float64(l+dl))))
			vertices = append(vertices, float32(R*math.Cos(float64(l+dl))))
			vertices = append(vertices, float32((h+dh)*H))
			vertices = append(vertices, 1, 1)

			// 0123 четырехугольник = 012 + 230
			//x2,y2,z2,u2,v2
			vertices = append(vertices, float32(R*math.Sin(float64(l+dl))))
			vertices = append(vertices, float32(R*math.Cos(float64(l+dl))))
			vertices = append(vertices, float32((h+dh)*H))
			vertices = append(vertices, 1, 1)

			//x3,y3,z3,u3,v3
			vertices = append(vertices, float32(R*math.Sin(float64(l+dl))))
			vertices = append(vertices, float32(R*math.Cos(float64(l+dl))))
			vertices = append(vertices, float32(h*H))
			vertices = append(vertices, 1, 1)

			//x0,y0,z0,u0,v0
			vertices = append(vertices, float32(R*math.Sin(float64(l))))
			vertices = append(vertices, float32(R*math.Cos(float64(l))))
			vertices = append(vertices, float32(h*H))
			vertices = append(vertices, 1, 1)

		}
	}

	for i := range vertices {
		indices = append(indices, uint32(i))
	}

	return &Cyllinder{
		vertices:     vertices,
		indices:      indices,
		rotation:     0,
		position:     mgl32.Vec3{},
		rotationAxes: mgl32.Vec3{0, 1, 0},
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

func (c *Cyllinder) RotationAxes() mgl32.Vec3 {
	return c.rotationAxes
}
