package opengl

import (
	"fmt"
	"terrain/shader"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func InitOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL Version: " + version)

	VertexShader, err := shader.CompileShader(shader.VertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	FragmentShader, err := shader.CompileShader(shader.FragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, VertexShader)
	gl.AttachShader(prog, FragmentShader)
	gl.LinkProgram(prog)
	gl.Enable(gl.DEPTH_TEST)
	return prog
}
