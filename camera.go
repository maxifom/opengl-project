package main

import (
	"fmt"
	"log"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type CameraDirection string

const (
	FORWARD  = "forward"
	BACKWARD = "backward"
	LEFT     = "left"
	RIGHT    = "right"

	XPLUS  = "x_plus"
	XMINUS = "x_minus"

	YPLUS  = "y_plus"
	YMINUS = "y_minus"

	ZPLUS  = "z_plus"
	ZMINUS = "z_minus"
)

type Camera struct {
	Yaw, Pitch, MovementSpeed, MouseSensitivity, Zoom float32

	Position, Front, Up, Right, WorldUp mgl32.Vec3
}

func NewCamera(yaw float32, pitch float32, position mgl32.Vec3, up mgl32.Vec3) Camera {
	camera := Camera{
		Yaw:              yaw,
		Pitch:            pitch,
		Position:         position,
		WorldUp:          up,
		Front:            mgl32.Vec3{0, 0, -1},
		MovementSpeed:    5,
		MouseSensitivity: 0.1,
		Zoom:             45.0,
	}

	camera.UpdateCameraVectors()

	return camera
}

func (c *Camera) GetViewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(c.Position, c.Position.Add(c.Front), c.Up)
}

func (c *Camera) ProcessKeyboard(direction CameraDirection, deltaTime float32) {
	velocity := c.MovementSpeed * deltaTime
	switch direction {
	case FORWARD:
		c.Position = c.Position.Add(c.Front.Mul(velocity))
	case BACKWARD:
		c.Position = c.Position.Sub(c.Front.Mul(velocity))
	case LEFT:
		c.Position = c.Position.Sub(c.Right.Mul(velocity))
	case RIGHT:
		c.Position = c.Position.Add(c.Right.Mul(velocity))
	case ZPLUS:
		c.Position = c.Position.Sub(c.Up.Mul(velocity))
	case ZMINUS:
		c.Position = c.Position.Add(c.Up.Mul(velocity))
	case YPLUS:
		c.ProcessMouseMovement(0, 5, true)
	case YMINUS:
		c.ProcessMouseMovement(0, -5, true)
	case XMINUS:
		c.ProcessMouseMovement(-5, 0, true)
	case XPLUS:
		c.ProcessMouseMovement(5, 0, true)
	default:
		panic("invalid direction " + direction)
	}
}

func (c *Camera) ProcessMouseMovement(xOffset, yOffset float32, constrainPitch bool) {
	xOffset *= c.MouseSensitivity
	yOffset *= c.MouseSensitivity

	c.Yaw += xOffset
	c.Pitch += yOffset

	if constrainPitch {
		if c.Pitch > 89 {
			c.Pitch = 89
		}
		if c.Pitch < -89 {
			c.Pitch = -89
		}
	}

	c.UpdateCameraVectors()
	log.Println(c)
}

func (c *Camera) ProcessMouseScroll(yOffset float32) {
	if c.Zoom >= 1 && c.Zoom <= 90 {
		c.Zoom -= yOffset
	}
	if c.Zoom <= 1 {
		c.Zoom = 1
	}
	if c.Zoom >= 90 {
		c.Zoom = 90
	}
}

func (c *Camera) UpdateCameraVectors() {
	c.Front = mgl32.Vec3{
		float32(math.Cos(float64(mgl32.DegToRad(c.Yaw))) * math.Cos(float64(mgl32.DegToRad(c.Pitch)))),
		float32(math.Sin(float64(mgl32.DegToRad(c.Pitch)))),
		float32(math.Sin(float64(mgl32.DegToRad(c.Yaw))) * math.Cos(float64(mgl32.DegToRad(c.Pitch)))),
	}.Normalize()
	c.Right = c.Front.Cross(c.WorldUp).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()
}

func (c *Camera) String() string {
	return fmt.Sprintf("Camera: \nPosition: %#v\nFront: %#v\nRight:%#v\nUp:%#v\nWorldUp:%#v\nPitch:%.8f\nYaw:%.8f\nZoom:%.8f\nMouseSensitivity:%.8f\nMovement Speed: %.2f\n", c.Position, c.Front, c.Right, c.Up, c.WorldUp, c.Pitch, c.Yaw, c.Zoom, c.MouseSensitivity, c.MovementSpeed)
}
