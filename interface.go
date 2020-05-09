package main

type Object interface {
	Vertices() []float32
	Indices() []uint32
}
