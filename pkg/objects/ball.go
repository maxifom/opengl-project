package objects

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

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

func NewBall(r float64, pos mgl32.Vec3, rotation float32, rotationAxes mgl32.Vec3) *Ball {
	var vertices []float32
	var indices []uint32

	db := mgl64.DegToRad(1.0)
	dl := mgl64.DegToRad(1.0)
	// db, dl - шаги широты и долготы

	// TODO: image texture will not work, only solid color (в книге описано как сделать)
	for b := mgl64.DegToRad(0); b < mgl64.DegToRad(1440); b += db {
		for l := mgl64.DegToRad(0); l < mgl64.DegToRad(360); l += dl {
			//x0,y0,z0,u0,v0
			x := float32(r * math.Cos(b) * math.Sin(l))
			y := float32(r * math.Cos(b) * math.Cos(l))
			z := float32(r * math.Sin(b))
			u := float32(1.0)
			v := float32(1.0)
			vertices = append(vertices, x, y, z, u, v)
			indice0 := uint32(len(vertices)/5 - 1)

			//x1,y1,z1,u1,v1
			x = float32(r * math.Cos(b+db) * math.Sin(l))
			y = float32(r * math.Cos(b+db) * math.Cos(l))
			z = float32(r * math.Sin(b+db))
			u = float32(1.0)
			v = float32(1.0)
			vertices = append(vertices, x, y, z, u, v)
			indice1 := uint32(len(vertices)/5 - 1)

			//x2,y2,z2,u2,v2
			x = float32(r * math.Cos(b+db) * math.Sin(l+dl))
			y = float32(r * math.Cos(b+db) * math.Cos(l+dl))
			z = float32(r * math.Sin(b+db))
			u = float32(1.0)
			v = float32(1.0)
			vertices = append(vertices, x, y, z, u, v)
			indice2 := uint32(len(vertices)/5 - 1)

			//x3,y3,z3,u3,v3
			x = float32(r * math.Cos(b) * math.Sin(l+dl))
			y = float32(r * math.Cos(b) * math.Cos(l+dl))
			z = float32(r * math.Sin(b))
			u = float32(1.0)
			v = float32(1.0)
			vertices = append(vertices, x, y, z, u, v)
			indice3 := uint32(len(vertices)/5 - 1)

			indices = append(indices,
				indice0, indice1, indice2,
				indice2, indice3, indice0,
			)
		}
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
