package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

// b - широта от -90 до 90
// l - долгота от 0 до 360
// x = R * cosB * sinL
// y = R * cosB * cosL
// z = R * sinB

type Ball struct {
	vertices     []float32
	indices      []uint32
	rotation     float32
	position     mgl32.Vec3
	rotationAxes mgl32.Vec3
}

func NewBall(r, db, dl float64, pos mgl32.Vec3, rotation float32, rotationAxes mgl32.Vec3) *Ball {
	var vertices []float32
	var indices []uint32

	// b - широта от -90 до 90
	// l - долгота от 0 до 360
	// db, dl - шаги широты и долготы

	// TODO: image texture will not work, only solid color (в книге описано как сделать)
	for b := float64(-90); b < 90; b += db {
		for l := float64(0); l < 360; l += dl {
			//x0,y0,z0,u0,v0
			vertices = append(vertices, float32(r*math.Cos(b)*math.Sin(l)))
			vertices = append(vertices, float32(r*math.Cos(b)*math.Cos(l)))
			vertices = append(vertices, float32(r*math.Sin(b)))
			vertices = append(vertices, 1, 1)

			//x1,y1,z1,u1,v1
			vertices = append(vertices, float32(r*math.Cos(b+db)*math.Sin(l)))
			vertices = append(vertices, float32(r*math.Cos(b+db)*math.Cos(l)))
			vertices = append(vertices, float32(r*math.Sin(b+db)))
			vertices = append(vertices, 1, 1)

			//x2,y2,z2,u2,v2
			vertices = append(vertices, float32(r*math.Cos(b+db)*math.Sin(l+dl)))
			vertices = append(vertices, float32(r*math.Cos(b+db)*math.Cos(l+dl)))
			vertices = append(vertices, float32(r*math.Sin(b+db)))
			vertices = append(vertices, 1, 1)

			// 0123 четырехугольник = 012 + 230
			//x2,y2,z2,u2,v2
			vertices = append(vertices, float32(r*math.Cos(b+db)*math.Sin(l+dl)))
			vertices = append(vertices, float32(r*math.Cos(b+db)*math.Cos(l+dl)))
			vertices = append(vertices, float32(r*math.Sin(b+db)))
			vertices = append(vertices, 1, 1)

			//x3,y3,z3,u3,v3
			vertices = append(vertices, float32(r*math.Cos(b)*math.Sin(l+dl)))
			vertices = append(vertices, float32(r*math.Cos(b)*math.Cos(l+dl)))
			vertices = append(vertices, float32(r*math.Sin(b)))
			vertices = append(vertices, 1, 1)

			//x0,y0,z0,u0,v0
			vertices = append(vertices, float32(r*math.Cos(b)*math.Sin(l)))
			vertices = append(vertices, float32(r*math.Cos(b)*math.Cos(l)))
			vertices = append(vertices, float32(r*math.Sin(b)))
			vertices = append(vertices, 1, 1)
		}
	}

	for i := range vertices {
		indices = append(indices, uint32(i))
	}

	return &Ball{
		vertices:     vertices,
		indices:      indices,
		rotation:     rotation,
		position:     pos,
		rotationAxes: rotationAxes,
	}
}

func (b *Ball) Vertices() []float32 {
	return b.vertices
}

func (b *Ball) Indices() []uint32 {
	return b.indices
}

func (b *Ball) Position() mgl32.Vec3 {
	return b.position
}

func (b *Ball) SetPosition(vec3 mgl32.Vec3) {
	b.position = vec3
}

func (b *Ball) Rotation() float32 {
	return b.rotation
}

func (b *Ball) SetRotation(f float32) {
	b.rotation = f
}

func (b *Ball) RotationAxes() mgl32.Vec3 {
	return b.rotationAxes
}

func (b *Ball) DrawMode() uint32 {
	return gl.TRIANGLES
}

func (b *Ball) Texture() uint32 {
	return 1
}

func (b *Ball) SetRotationAxes(vec3 mgl32.Vec3) {
	b.rotationAxes = vec3
}
