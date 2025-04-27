package main

import (
	"runtime"
	"terrain/camera"
	"terrain/model"
	"terrain/opengl"
	"terrain/utils"
	"terrain/window"
	"terrain/world"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var mod model.Model

func main() {
	runtime.LockOSThread()

	window := window.MakeWindow()
	defer glfw.Terminate()

	program := opengl.InitOpenGL()

	terrain := world.NewTerrain()
	terrain.GenVertecies(100, 100)
	terrain.GenVao()

	mod = model.NewModel("./model/test.obj")
	mod.GenVao()

	for !window.ShouldClose() {
		camera.CurrentCamera.Update(window)

		if window.GetKey(glfw.Key1) == glfw.Press {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
			gl.ClearColor(0.35, 0.35, 1, 1)
		} else if window.GetKey(glfw.Key2) == glfw.Press {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
			gl.ClearColor(0.35, 0.35, 1, 1)
		} else if window.GetKey(glfw.Key3) == glfw.Press {
			gl.PolygonMode(gl.FRONT_AND_BACK, gl.POINT)
			// gl.ClearColor(0.2, 0.2, 0.2, 1)
		}

		utils.DT = float32(glfw.GetTime()) - utils.LastTime
		utils.DT *= 60
		utils.LastTime = float32(glfw.GetTime())

		draw(&terrain, window, program)
	}
}
