package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func ReadShaderFromFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("failed to read shader from file %s: %s", filename, err)
	}

	return string(bytes)
}
