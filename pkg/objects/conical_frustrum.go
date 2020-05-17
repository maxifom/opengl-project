package objects

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

// Усеченный конус
type ConicalFrustrum struct {
	vertices     []float32
	indices      []uint32
	rotation     float32
	position     mgl32.Vec3
	rotationAxes mgl32.Vec3
	texture      uint32
}

func NewConicalFrustrum(H, Rmax, Rmin, angle float64, enableTop, enableBottom bool, position mgl32.Vec3, rotation float32, rotationAxes mgl32.Vec3, texture uint32) []Object {
	var vertices []float32
	var indices []uint32

	var r = Rmax
	dr := (Rmax - Rmin) / 100
	dh := 2. / 500
	dl := mgl64.DegToRad(angle)
	for h := -1.0; h < 1; h += dh {
		r -= dr
		for l := mgl64.DegToRad(0); l < mgl64.DegToRad(360); l += dl {
			//x0,y0,z0,u0,v0
			vertices = append(vertices, float32(r*math.Sin(l)))
			vertices = append(vertices, float32(r*math.Cos(l)))
			vertices = append(vertices, float32(h*H))
			vertices = append(vertices, 1, 1)
			indice0 := uint32(len(vertices)/5 - 1)

			//x1,y1,z1,u1,v1
			vertices = append(vertices, float32(r*math.Sin(l)))
			vertices = append(vertices, float32(r*math.Cos(l)))
			vertices = append(vertices, float32((h+dh)*H))
			vertices = append(vertices, 1, 1)
			indice1 := uint32(len(vertices)/5 - 1)

			//x2,y2,z2,u2,v2
			vertices = append(vertices, float32(r*math.Sin(l+dl)))
			vertices = append(vertices, float32(r*math.Cos(l+dl)))
			vertices = append(vertices, float32((h+dh)*H))
			vertices = append(vertices, 1, 1)
			indice2 := uint32(len(vertices)/5 - 1)

			//x3,y3,z3,u3,v3
			vertices = append(vertices, float32(r*math.Sin(l+dl)))
			vertices = append(vertices, float32(r*math.Cos(l+dl)))
			vertices = append(vertices, float32(h*H))
			vertices = append(vertices, 1, 1)
			indice3 := uint32(len(vertices)/5 - 1)

			indices = append(indices,
				indice0, indice1, indice2,
				indice2, indice3, indice0,
			)
		}
	}

	objects := []Object{
		&ConicalFrustrum{
			vertices:     vertices,
			indices:      indices,
			rotation:     rotation,
			position:     position,
			rotationAxes: rotationAxes,
			texture:      texture,
		},
	}

	if enableTop {
		objects = append(objects, NewXYEllipse(0, 0, float32(Rmax), float32(Rmax), float32(-H), angle, position, rotation, rotationAxes, texture))
	}
	if enableBottom {
		objects = append(objects, NewXYEllipse(0, 0, float32(Rmin), float32(Rmin), float32(-0.6*H), angle, position, rotation, rotationAxes, texture))
	}

	return objects
}

func (c *ConicalFrustrum) Vertices() []float32 {
	return c.vertices
}

func (c *ConicalFrustrum) Indices() []uint32 {
	return c.indices
}

func (c *ConicalFrustrum) Position() mgl32.Vec3 {
	return c.position
}

func (c *ConicalFrustrum) SetPosition(vec3 mgl32.Vec3) {
	c.position = vec3
}

func (c *ConicalFrustrum) Rotation() float32 {
	return c.rotation
}

func (c *ConicalFrustrum) RotationAxes() mgl32.Vec3 {
	return c.rotationAxes
}

func (c *ConicalFrustrum) SetRotation(f float32) {
	c.rotation = f
}

func (c *ConicalFrustrum) SetRotationAxes(vec3 mgl32.Vec3) {
	c.rotationAxes = vec3
}

func (c *ConicalFrustrum) DrawMode() uint32 {
	return gl.TRIANGLES
}

func (c *ConicalFrustrum) Texture() uint32 {
	return c.texture
}
