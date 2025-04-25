package shader

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

var VertexShaderSource = `
	#version 410
    	layout (location = 0) in vec3 aPos;
    	layout (location = 1) in vec3 aColor;

	uniform mat4 model;
	uniform mat4 view;
	uniform mat4 projection;

    	out vec4 Pos;
    	out vec3 Color;

	void main()
	{
    		gl_Position = projection * model * view * vec4(aPos, 1.0);
		Pos = model * vec4(aPos, 1);
		Color = aColor;
	}

` + "\x00"

var FragmentShaderSource = `
    	#version 410
    	out vec4 frag_colour;
    	in vec4 Pos;
	in vec3 Color;
    	void main() {
        	frag_colour = vec4(0, Pos.y+2.5, 0, 1);
    	}
` + "\x00"

func CompileShader(source string, shaderType uint32) (uint32, error) {
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
