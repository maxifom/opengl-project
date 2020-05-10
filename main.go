package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const float32Size = 4
const uint32Size = 4
const windowWidth = 800
const windowHeight = 600

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	c := NewCamera(-90, 0, mgl32.Vec3{0, 0, 20}, mgl32.Vec3{0, 1, 0})
	var lastXPosition *float64
	var lastYPosition *float64

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "OpenGL Tank", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		gl.Viewport(0, 0, int32(width), int32(height))
	})
	window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		if lastXPosition == nil {
			lastXPosition = &xpos
			lastYPosition = &ypos
		}

		xOffset := xpos - *lastXPosition
		yOffset := *lastYPosition - ypos

		lastXPosition = &xpos
		lastYPosition = &ypos
		c.ProcessMouseMovement(float32(xOffset), float32(yOffset), true)
	})
	window.SetScrollCallback(func(w *glfw.Window, xoff float64, yoff float64) {
		c.ProcessMouseScroll(float32(yoff))
	})

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// Configure the vertex and fragment shaders
	vertexShader := ReadShaderFromFile("shaders/vertex")
	fragmentShader := ReadShaderFromFile("shaders/fragment")
	program, err := newProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}

	gl.UseProgram(program)

	projection := mgl32.Perspective(mgl32.DegToRad(c.Zoom), float32(windowWidth)/windowHeight, 0.1, 100.0)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	mat4 := c.GetViewMatrix()
	gl.UniformMatrix4fv(cameraUniform, 1, false, &mat4[0])

	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)

	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))

	// Load the texture
	texture, err := newTexture("tank.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	// Configure the vertex data
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var objects []Object
	objects = append(objects, NewParallelepiped(3, 4, 5, mgl32.Vec3{0, 0, 0}))
	objects = append(objects, NewParallelepiped(2, 3, 4, mgl32.Vec3{4, 0, 0}))

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*float32Size, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*float32Size, gl.PtrOffset(3*float32Size))

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	angle := 0.0
	previousTime := glfw.GetTime()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Update
		time := glfw.GetTime()
		elapsed := time - previousTime
		previousTime = time

		ProcessInput(&c, window, float32(elapsed), objects)

		angle += elapsed
		angle = 0
		model = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})

		// Render
		gl.UseProgram(program)
		for _, object := range objects {
			v := object.Vertices()
			gl.BufferData(gl.ARRAY_BUFFER, len(v)*float32Size, gl.Ptr(v), gl.STATIC_DRAW)
			model = mgl32.Ident4()
			model = TranslateMat4Vec3(model, object.Position())
			gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

			gl.BindVertexArray(vao)

			gl.ActiveTexture(gl.TEXTURE0)
			gl.BindTexture(gl.TEXTURE_2D, texture)

			projection = mgl32.Perspective(mgl32.DegToRad(c.Zoom), float32(windowWidth)/windowHeight, 0.1, 100.0)
			gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

			gl.DrawArrays(gl.TRIANGLES, 0, int32(len(v)))
		}

		mat4 := c.GetViewMatrix()
		gl.UniformMatrix4fv(cameraUniform, 1, false, &mat4[0])

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
