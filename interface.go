package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Object interface {
	Vertices() []float32
	Position() mgl32.Vec3
	SetPosition(vec3 mgl32.Vec3)
	Rotation() float32
	SetRotation(float32)
}
