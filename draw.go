package main

import (
	"terrain/camera"
	"terrain/world"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func draw(terrain *world.Terrain, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	modelUniformLoc := gl.GetUniformLocation(program, gl.Str("model\x00"))
	viewUniformLoc := gl.GetUniformLocation(program, gl.Str("view\x00"))
	projectionUniformLoc := gl.GetUniformLocation(program, gl.Str("projection\x00"))

	gl.UniformMatrix4fv(modelUniformLoc, 1, false, &camera.CurrentCamera.Model[0])
	gl.UniformMatrix4fv(viewUniformLoc, 1, false, &camera.CurrentCamera.View[0])
	gl.UniformMatrix4fv(projectionUniformLoc, 1, false, &camera.CurrentCamera.Perspective[0])

	gl.BindVertexArray(terrain.Vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(terrain.Verticies)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}
