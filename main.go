package main

import (
	"runtime"
	"terrain/camera"
	"terrain/opengl"
	"terrain/utils"
	"terrain/window"
	"terrain/world"

	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	runtime.LockOSThread()

	window := window.MakeWindow()
	defer glfw.Terminate()

	program := opengl.InitOpenGL()

	terrain := world.NewTerrain()
	terrain.GenVertecies(1000, 1000)
	terrain.GenVao()

	for !window.ShouldClose() {
		camera.CurrentCamera.Update(window)

		utils.DT = float32(glfw.GetTime()) - utils.LastTime
		utils.DT *= 60
		utils.LastTime = float32(glfw.GetTime())

		draw(&terrain, window, program)
	}
}
