package objects

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

// Цилиндр без крышек
type XYEllipse struct {
	vertices     []float32
	indices      []uint32
	rotation     float32
	position     mgl32.Vec3
	rotationAxes mgl32.Vec3
	texture      uint32
}

func NewXYEllipse(cx, cy, rx, ry, z float32, angle float64, position mgl32.Vec3, rotation float32, rotationAxes mgl32.Vec3, texture uint32) *XYEllipse {
	var vertices []float32
	var indices []uint32

	da := mgl64.DegToRad(angle)
	vertices = append(vertices, cx, cy, z, 1, 1)
	indices = append(indices, 0)

	for angle := mgl64.DegToRad(0); angle <= mgl64.DegToRad(2160); angle += da {
		sin, cos := math.Sincos(angle)
		dx := rx * float32(cos)
		dy := ry * float32(sin)

		vertices = append(vertices, dx+cx, dy+cy, z, 1, 1)
		indices = append(indices, uint32(len(vertices)/5-1))

		sin, cos = math.Sincos(angle + da)
		dx = rx * float32(cos)
		dy = ry * float32(sin)

		vertices = append(vertices, dx+cx, dy+cy, z, 1, 1)
		indices = append(indices, uint32(len(vertices)/5-1))
	}

	return &XYEllipse{
		vertices:     vertices,
		indices:      indices,
		rotation:     rotation,
		position:     position,
		rotationAxes: rotationAxes,
		texture:      texture,
	}
}

func (c *XYEllipse) Vertices() []float32 {
	return c.vertices
}

func (c *XYEllipse) Indices() []uint32 {
	return c.indices
}

func (c *XYEllipse) Position() mgl32.Vec3 {
	return c.position
}

func (c *XYEllipse) SetPosition(vec3 mgl32.Vec3) {
	c.position = vec3
}

func (c *XYEllipse) Rotation() float32 {
	return c.rotation
}

func (c *XYEllipse) SetRotation(f float32) {
	c.rotation = f
}

func (c *XYEllipse) RotationAxes() mgl32.Vec3 {
	return c.rotationAxes
}

func (c *XYEllipse) DrawMode() uint32 {
	return gl.TRIANGLE_FAN
}

func (c *XYEllipse) Texture() uint32 {
	return c.texture
}

func (c *XYEllipse) SetRotationAxes(vec3 mgl32.Vec3) {
	c.rotationAxes = vec3
}
