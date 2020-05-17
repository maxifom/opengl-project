package main

import (
	"example/pkg/objects"
	"fmt"
	"image"
	"image/draw"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func newTexture(file string, n int) (uint32, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return 0, fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	defer imgFile.Close()
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return 0, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*float32Size {
		return 0, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(uint32(gl.TEXTURE0 + n))
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture, nil
}

var currentObject = 0
var lastTimeChangedCurrentObject = time.Time{}

func ProcessInput(camera *Camera, window *glfw.Window, deltaTime float32, objects []objects.Object) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}

	if window.GetKey(glfw.KeyW) == glfw.Press {
		camera.ProcessKeyboard(FORWARD, deltaTime)
	}
	if window.GetKey(glfw.KeyS) == glfw.Press {
		camera.ProcessKeyboard(BACKWARD, deltaTime)
	}
	if window.GetKey(glfw.KeyA) == glfw.Press {
		camera.ProcessKeyboard(LEFT, deltaTime)
	}
	if window.GetKey(glfw.KeyD) == glfw.Press {
		camera.ProcessKeyboard(RIGHT, deltaTime)
	}
	if window.GetKey(glfw.KeyKP1) == glfw.Press {
		camera.ProcessKeyboard(ZMINUS, deltaTime)
	}
	if window.GetKey(glfw.KeyKP3) == glfw.Press {
		camera.ProcessKeyboard(ZPLUS, deltaTime)
	}
	if window.GetKey(glfw.KeyKP2) == glfw.Press {
		camera.ProcessKeyboard(YMINUS, deltaTime)
	}
	if window.GetKey(glfw.KeyKP4) == glfw.Press {
		camera.ProcessKeyboard(XMINUS, deltaTime)
	}
	if window.GetKey(glfw.KeyKP6) == glfw.Press {
		camera.ProcessKeyboard(XPLUS, deltaTime)
	}
	if window.GetKey(glfw.KeyKP8) == glfw.Press {
		camera.ProcessKeyboard(YPLUS, deltaTime)
	}
	if window.GetKey(glfw.KeyKPAdd) == glfw.Press {
		if time.Since(lastTimeChangedCurrentObject) < 300*time.Millisecond {
			return
		}
		if currentObject+1 == len(objects) {
			currentObject = 0
		} else {
			currentObject++
		}
		log.Println(currentObject)
		lastTimeChangedCurrentObject = time.Now()
	}

	if window.GetKey(glfw.KeyUp) == glfw.Press {
		objects[currentObject].SetPosition(objects[currentObject].Position().Add(mgl32.Vec3{0, 0.1, 0}))
	}
	if window.GetKey(glfw.KeyDown) == glfw.Press {
		objects[currentObject].SetPosition(objects[currentObject].Position().Add(mgl32.Vec3{0, -0.1, 0}))
	}
	if window.GetKey(glfw.KeyLeft) == glfw.Press {
		log.Println(objects[currentObject].Position().X())
		objects[currentObject].SetPosition(objects[currentObject].Position().Add(mgl32.Vec3{-0.1, 0, 0}))
	}
	if window.GetKey(glfw.KeyRight) == glfw.Press {
		objects[currentObject].SetPosition(objects[currentObject].Position().Add(mgl32.Vec3{0.1, 0, 0}))
	}
	if window.GetKey(glfw.KeyKP7) == glfw.Press {
		log.Println(objects[currentObject].Position().Z())
		objects[currentObject].SetPosition(objects[currentObject].Position().Add(mgl32.Vec3{0, 0, -0.1}))
	}
	if window.GetKey(glfw.KeyKP9) == glfw.Press {
		log.Println(objects[currentObject].Position().Z())
		objects[currentObject].SetPosition(objects[currentObject].Position().Add(mgl32.Vec3{0, 0, 0.1}))
	}

	if window.GetKey(glfw.KeyE) == glfw.Press {
		objects[currentObject].SetRotation(objects[currentObject].Rotation() + 1)
	}

	if window.GetKey(glfw.KeyQ) == glfw.Press {
		objects[currentObject].SetRotation(objects[currentObject].Rotation() - 1)
	}

	if window.GetKey(glfw.KeyZ) == glfw.Press {
		if objects[currentObject].RotationAxes().Z() != 1 {
			objects[currentObject].SetRotation(0)
		}
		objects[currentObject].SetRotationAxes(mgl32.Vec3{0, 0, 1})
		objects[currentObject].SetRotation(objects[currentObject].Rotation() + 1)
	}
	if window.GetKey(glfw.KeyY) == glfw.Press {
		if objects[currentObject].RotationAxes().Y() != 1 {
			objects[currentObject].SetRotation(0)
		}
		objects[currentObject].SetRotationAxes(mgl32.Vec3{0, 1, 0})
		objects[currentObject].SetRotation(objects[currentObject].Rotation() + 1)
	}
	if window.GetKey(glfw.KeyX) == glfw.Press {
		if objects[currentObject].RotationAxes().X() != 1 {
			objects[currentObject].SetRotation(0)
		}
		objects[currentObject].SetRotationAxes(mgl32.Vec3{1, 0, 0})
		objects[currentObject].SetRotation(objects[currentObject].Rotation() + 1)
	}
}

func TranslateMat4Vec3(mat4 mgl32.Mat4, vec3 mgl32.Vec3) mgl32.Mat4 {
	return mat4.Mul4(mgl32.Translate3D(vec3[0], vec3[1], vec3[2]))
}
