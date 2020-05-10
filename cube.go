package main

type Cube struct {
	vertices []float32
}

func NewCube(a float32) Cube {
	// TODO: Fix texture
	vertices := []float32{
		//  X, Y, Z, U, V
		// Bottom
		-a, -a, -a, 0.0, 0.0,
		a, -a, -a, 1.0, 0.0,
		-a, -a, a, 0.0, 1.0,
		a, -a, -a, 1.0, 0.0,
		a, -a, a, 1.0, 1.0,
		-a, -a, a, 0.0, 1.0,

		// Top
		-a, a, -a, 0.0, 0.0,
		-a, a, a, 0.0, 1.0,
		a, a, -a, 1.0, 0.0,
		a, a, -a, 1.0, 0.0,
		-a, a, a, 0.0, 1.0,
		a, a, a, 1.0, 1.0,

		// Front
		-a, -a, a, 1.0, 0.0,
		a, -a, a, 0.0, 0.0,
		-a, a, a, 1.0, 1.0,
		a, -a, a, 0.0, 0.0,
		a, a, a, 0.0, 1.0,
		-a, a, a, 1.0, 1.0,

		// Back
		-a, -a, -a, 0.0, 0.0,
		-a, a, -a, 0.0, 1.0,
		a, -a, -a, 1.0, 0.0,
		a, -a, -a, 1.0, 0.0,
		-a, a, -a, 0.0, 1.0,
		a, a, -a, 1.0, 1.0,

		// Left
		-a, -a, a, 0.0, 1.0,
		-a, a, -a, 1.0, 0.0,
		-a, -a, -a, 0.0, 0.0,
		-a, -a, a, 0.0, 1.0,
		-a, a, a, 1.0, 1.0,
		-a, a, -a, 1.0, 0.0,

		// Right
		a, -a, a, 1.0, 1.0,
		a, -a, -a, 1.0, 0.0,
		a, a, -a, 0.0, 0.0,
		a, -a, a, 1.0, 1.0,
		a, a, -a, 0.0, 0.0,
		a, a, a, 0.0, 1.0,
	}

	return Cube{vertices: vertices}
}

func (c Cube) Vertices() []float32 {
	return c.vertices
}
