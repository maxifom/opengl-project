package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type Cyllinder struct {
	baseRadius float32
	topRadius  float32
	height     float32
	sectors    int
	stacks     int

	vertices     []float32
	unitVertices []float32

	indices []uint32

	rotation float32
	position mgl32.Vec3
}

func NewCyllinder(baseRadius float32, topRadius float32, height float32, sectors int, stacks int, rotation float32, position mgl32.Vec3) *Cyllinder {
	c := &Cyllinder{baseRadius: baseRadius, topRadius: topRadius, height: height, sectors: sectors, stacks: stacks, rotation: rotation, position: position}
	c.GenerateVertices()

	return c
}

func (c *Cyllinder) GenerateUnitVertices() {
	sectorStep := 2 * math.Pi / float64(c.sectors)

	var sectorAngle float64

	for i := 0; i <= c.sectors; i++ {
		sectorAngle = float64(i) * sectorStep
		c.unitVertices = append(c.unitVertices, float32(math.Cos(sectorAngle)))
		c.unitVertices = append(c.unitVertices, float32(math.Sin(sectorAngle)))
		c.unitVertices = append(c.unitVertices, 0)
	}
}

func (c *Cyllinder) GenerateVertices() {
	c.GenerateUnitVertices()

	//var x, y, z, radius float32
	var z, radius float32

	for i := 0; i <= c.stacks; i++ {
		z := (-c.height * 0.5) + float32(i)
		radius = c.baseRadius + float32(i)/float32(c.stacks)*(c.topRadius-c.baseRadius)
		t := 1 - float32(i)/float32(c.stacks)

		var k = 0

		for j := 0; j <= c.sectors; j++ {
			x := c.unitVertices[k]
			y := c.unitVertices[k+1]
			c.vertices = append(c.vertices, x*radius, y*radius, z)

			c.vertices = append(c.vertices, float32(j)/float32(c.sectors), t)

			k += 3
		}
	}

	//baseVertexIndex := len(c.vertices) / 3
	//z := -c.height * 0.5

	c.vertices = append(c.vertices, 0, 0, z)
	c.vertices = append(c.vertices, 0.5, 0.5)

	var j = 0
	for i := 0; i < c.sectors; i++ {
		x := c.unitVertices[j]
		y := c.unitVertices[j+1]
		c.vertices = append(c.vertices, x*c.baseRadius, y*c.baseRadius, z)
		c.vertices = append(c.vertices, -x*0.5+0.5, -y*0.5+0.5)

		j += 3
	}
}

func (c *Cyllinder) Vertices() []float32 {
	return c.vertices
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
