package shader

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

var VertexShaderSource = `
	#version 410
    	layout (location = 0) in vec3 aPos;
    	layout (location = 1) in vec3 aNormal;

	uniform mat4 model;
	uniform mat4 view;
	uniform mat4 projection;

    	out vec3 Pos;
    	out vec3 Normal;

	void main()
	{
    		gl_Position = projection * model * view * vec4(aPos, 1.0);
		Pos = aPos;
		Normal = aNormal;
	}

` + "\x00"

var FragmentShaderSource = `
    	#version 410
    	out vec4 frag_colour;
    	in vec3 Pos;
		in vec3 Normal;

    	void main() {
    		// float shade = Normal.x + Normal.y + Normal.z / 4-0.2;
        	// frag_colour = vec4(shade/4, shade+0.1/4, shade/4, 1);
        	frag_colour = vec4(Normal.x*100, Normal.y*100, Normal.z*100, 1);
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
