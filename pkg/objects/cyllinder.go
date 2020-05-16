package objects

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

// Цилиндр без крышек
type Cyllinder struct {
	vertices     []float32
	indices      []uint32
	rotation     float32
	position     mgl32.Vec3
	rotationAxes mgl32.Vec3
}

func NewCyllinder(H, R float64, position mgl32.Vec3, rotation float32, rotationAxes mgl32.Vec3) *Cyllinder {
	var vertices []float32
	var indices []uint32

	dh := H / 200.0
	dl := 1.0
	for h := -0.5; h < 0.5; h += dh {
		for l := float64(0); l < 360; l += dl {

			//x0,y0,z0,u0,v0
			vertices = append(vertices, float32(R*math.Sin(l)))
			vertices = append(vertices, float32(R*math.Cos(l)))
			vertices = append(vertices, float32(h*H))
			vertices = append(vertices, 1, 1)

			//x1,y1,z1,u1,v1
			vertices = append(vertices, float32(R*math.Sin(l)))
			vertices = append(vertices, float32(R*math.Cos(l)))
			vertices = append(vertices, float32((h+dh)*H))
			vertices = append(vertices, 1, 1)

			//x2,y2,z2,u2,v2
			vertices = append(vertices, float32(R*math.Sin(l+dl)))
			vertices = append(vertices, float32(R*math.Cos(l+dl)))
			vertices = append(vertices, float32((h+dh)*H))
			vertices = append(vertices, 1, 1)

			// 0123 четырехугольник = 012 + 230
			//x2,y2,z2,u2,v2
			vertices = append(vertices, float32(R*math.Sin(l+dl)))
			vertices = append(vertices, float32(R*math.Cos(l+dl)))
			vertices = append(vertices, float32((h+dh)*H))
			vertices = append(vertices, 1, 1)

			//x3,y3,z3,u3,v3
			vertices = append(vertices, float32(R*math.Sin(l+dl)))
			vertices = append(vertices, float32(R*math.Cos(l+dl)))
			vertices = append(vertices, float32(h*H))
			vertices = append(vertices, 1, 1)

			//x0,y0,z0,u0,v0
			vertices = append(vertices, float32(R*math.Sin(l)))
			vertices = append(vertices, float32(R*math.Cos(l)))
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
		rotation:     rotation,
		position:     position,
		rotationAxes: rotationAxes,
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

func (c *Cyllinder) DrawMode() uint32 {
	return gl.TRIANGLES
}

func (c *Cyllinder) Texture() uint32 {
	return 1
}

func (c *Cyllinder) SetRotationAxes(vec3 mgl32.Vec3) {
	c.rotationAxes = vec3
}
