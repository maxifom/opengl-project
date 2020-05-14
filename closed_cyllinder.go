package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

// Cyllinder с крышками
type ClosedCyllinder struct {
	vertices     []float32
	indices      []uint32
	rotation     float32
	position     mgl32.Vec3
	rotationAxes mgl32.Vec3
}

func NewClosedCyllinder(H, R, dh, dx float64, position mgl32.Vec3, rotation float32, rotationAxes mgl32.Vec3) *ClosedCyllinder {
	var vertices []float32
	var indices []uint32

	//TODO: dx dh чтобы было константное кол-во шагов (в зависимости от радиуса)
	//var H = 4.0
	//var R = 1.0
	//var dh = 0.005
	//var dx = 0.001
	var xc = 0.0
	var yc = 0.0
	for sinB := -0.5; sinB < 0.5; sinB += dh {
		for x := 0.0; x <= R; x += dx {
			y := math.Sqrt(R*R - x*x)
			h := float32(H * sinB)

			// Вертикаль от xc+x, yc+y, -h/2 до xc+x, yc+y, h/2 - боковая
			//vertices = append(vertices, float32(xc+x), float32(yc+y), -h/2, 1, 1)
			//vertices = append(vertices, float32(xc+x), float32(yc+y), h/2, 1, 1)

			// Вертикаль от xc-x, yc+y,-h/2 до xc-x, yc+y, h/2 - боковая
			//vertices = append(vertices, float32(xc-x), float32(yc+y), -h/2, 1, 1)
			//vertices = append(vertices, float32(xc-x), float32(yc+y), h/2, 1, 1)

			// Вертикаль от xc+x, yc-y, -h/2 до xc+x, yc+y, -h/2 - вертикаль основания
			vertices = append(vertices, float32(xc+x), float32(yc-y), -h/2, 1, 1)
			vertices = append(vertices, float32(xc+x), float32(yc+y), -h/2, 1, 1)

			// Вертикаль от xc-x, yc-y, -h/2 до xc-x, yc+y, -h/2 - вертикаль основания
			vertices = append(vertices, float32(xc-x), float32(yc-y), -h/2, 1, 1)
			vertices = append(vertices, float32(xc-x), float32(yc+y), -h/2, 1, 1)
		}
	}

	for i := range vertices {
		indices = append(indices, uint32(i))
	}

	return &ClosedCyllinder{
		vertices:     vertices,
		indices:      indices,
		rotation:     rotation,
		position:     position,
		rotationAxes: rotationAxes,
	}
}

func (c *ClosedCyllinder) Vertices() []float32 {
	return c.vertices
}

func (c *ClosedCyllinder) Indices() []uint32 {
	return c.indices
}

func (c *ClosedCyllinder) Position() mgl32.Vec3 {
	return c.position
}

func (c *ClosedCyllinder) SetPosition(vec3 mgl32.Vec3) {
	c.position = vec3
}

func (c *ClosedCyllinder) Rotation() float32 {
	return c.rotation
}

func (c *ClosedCyllinder) SetRotation(f float32) {
	c.rotation = f
}

func (c *ClosedCyllinder) RotationAxes() mgl32.Vec3 {
	return c.rotationAxes
}

func (c *ClosedCyllinder) DrawMode() uint32 {
	return gl.LINES
}

func (c *ClosedCyllinder) Texture() uint32 {
	return 1
}

func (c *ClosedCyllinder) SetRotationAxes(vec3 mgl32.Vec3) {
	c.rotationAxes = vec3
}
