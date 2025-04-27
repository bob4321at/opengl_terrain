package main

import (
	"terrain/world"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func draw(terrain *world.Terrain, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	mod.Draw(program)

	terrain.Draw(program)

	glfw.PollEvents()
	window.SwapBuffers()
}
