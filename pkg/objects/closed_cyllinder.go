package objects

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

//TODO: сделать через треугольники?
func NewClosedCyllinder(H, R float64, position mgl32.Vec3, rotation float32, rotationAxes mgl32.Vec3) *ClosedCyllinder {
	var vertices []float32
	var indices []uint32

	dx := R / 500.0
	dh := H / 500.0

	var xc = 0.0
	var yc = 0.0

	r2 := R * R
	for sinB := -0.5; sinB < 0.5; sinB += dh {
		h := float32(H * sinB)
		for x := 0.0; x <= R; x += dx {
			y := math.Sqrt(r2 - x*x)

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
