package main

type Cube struct {
	vertices []float32
	indices  []uint32
}

func NewCube(a float32) Cube {
	// TODO: Fix texture
	vertices := []float32{
		0, 0, 0, 1, 0,
		0, a, 0, 1, 0,
		a, 0, 0, 0, 0,
		a, a, 0, 1, 0,
		0, 0, a, 1, 0,
		a, 0, a, 1, 0,
		a, a, a, 1, 0,
		0, a, a, 1, 0,
	}

	indices := []uint32{
		0, 1, 3,
		0, 2, 3,

		0, 4, 2,
		4, 5, 2,

		1, 7, 6,
		6, 3, 1,

		3, 2, 5,
		6, 5, 3,

		4, 5, 6,
		6, 4, 7,

		1, 0, 3,
		7, 3, 1,
	}

	return Cube{vertices: vertices, indices: indices}
}

func (c Cube) Vertices() []float32 {
	return c.vertices
}

func (c Cube) Indices() []uint32 {
	return c.indices
}
