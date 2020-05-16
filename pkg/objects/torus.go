package objects

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type Torus struct {
	vertices     []float32
	indices      []uint32
	rotation     float32
	position     mgl32.Vec3
	rotationAxes mgl32.Vec3
}

func NewTorus(R, r float64, position mgl32.Vec3, rotation float32, rotationAxes mgl32.Vec3) *Torus {
	var vertices []float32
	var indices []uint32

	dw := 1.0
	dphi := 1.0

	for w := -180.0; w < 180.0; w += dw {
		for phi := -180.0; phi < 180.0; phi += dphi {
			//x0,y0,z0,u0,v0
			vertices = append(vertices, float32((R+r*math.Cos(phi))*math.Sin(w)))
			vertices = append(vertices, float32((R+r*math.Cos(phi))*math.Cos(w)))
			vertices = append(vertices, float32(r*math.Sin(phi)))
			vertices = append(vertices, 1, 1)

			//x1,y1,z1,u1,v1
			vertices = append(vertices, float32((R+r*math.Cos(phi+dphi))*math.Sin(w)))
			vertices = append(vertices, float32((R+r*math.Cos(phi+dphi))*math.Cos(w)))
			vertices = append(vertices, float32(r*math.Sin(phi+dphi)))
			vertices = append(vertices, 1, 1)

			//x2,y2,z2,u2,v2
			vertices = append(vertices, float32((R+r*math.Cos(phi+dphi))*math.Sin(w+dw)))
			vertices = append(vertices, float32((R+r*math.Cos(phi+dphi))*math.Cos(w+dw)))
			vertices = append(vertices, float32(r*math.Sin(phi+dphi)))
			vertices = append(vertices, 1, 1)

			// 0123 четырехугольник = 012 + 230
			//x2,y2,z2,u2,v2
			vertices = append(vertices, float32((R+r*math.Cos(phi+dphi))*math.Sin(w+dw)))
			vertices = append(vertices, float32((R+r*math.Cos(phi+dphi))*math.Cos(w+dw)))
			vertices = append(vertices, float32(r*math.Sin(phi+dphi)))
			vertices = append(vertices, 1, 1)

			//x3,y3,z3,u3,v3
			//x2,y2,z2,u2,v2
			vertices = append(vertices, float32((R+r*math.Cos(phi))*math.Sin(w+dw)))
			vertices = append(vertices, float32((R+r*math.Cos(phi))*math.Cos(w+dw)))
			vertices = append(vertices, float32(r*math.Sin(phi)))
			vertices = append(vertices, 1, 1)

			//x0,y0,z0,u0,v0
			vertices = append(vertices, float32((R+r*math.Cos(phi))*math.Sin(w)))
			vertices = append(vertices, float32((R+r*math.Cos(phi))*math.Cos(w)))
			vertices = append(vertices, float32(r*math.Sin(phi)))
			vertices = append(vertices, 1, 1)

		}
	}

	for i := range vertices {
		indices = append(indices, uint32(i))
	}

	return &Torus{
		vertices:     vertices,
		indices:      indices,
		rotation:     rotation,
		position:     position,
		rotationAxes: rotationAxes,
	}
}

func (c *Torus) Vertices() []float32 {
	return c.vertices
}

func (c *Torus) Indices() []uint32 {
	return c.indices
}

func (c *Torus) Position() mgl32.Vec3 {
	return c.position
}

func (c *Torus) SetPosition(vec3 mgl32.Vec3) {
	c.position = vec3
}

func (c *Torus) Rotation() float32 {
	return c.rotation
}

func (c *Torus) SetRotation(f float32) {
	c.rotation = f
}

func (c *Torus) RotationAxes() mgl32.Vec3 {
	return c.rotationAxes
}

func (c *Torus) DrawMode() uint32 {
	return gl.TRIANGLES
}

func (c *Torus) Texture() uint32 {
	return 0
}

func (c *Torus) SetRotationAxes(vec3 mgl32.Vec3) {
	c.rotationAxes = vec3
}
