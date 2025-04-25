package camera

import (
	"math"
	"terrain/utils"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Pos         mgl32.Vec3
	Target      mgl32.Vec3
	Direction   mgl32.Vec3
	WorldUp     mgl32.Vec3
	Right       mgl32.Vec3
	Up          mgl32.Vec3
	View        mgl32.Mat4
	Radius      float64
	CamZ        float32
	CamY        float32
	CamX        float32
	Perspective mgl32.Mat4
	Model       mgl32.Mat4

	CameraRotation mgl32.Vec3
}

func NewCamera(Pos mgl32.Vec3, Target_Pos mgl32.Vec3) (c Camera) {
	c.Pos = Pos
	c.Target = Target_Pos
	dir := mgl32.Vec3{Pos.X() - Target_Pos.X(), Pos.Y() - Target_Pos.Y(), Pos.Z() - Target_Pos.Z()}
	c.Direction = dir.Normalize()
	c.WorldUp = mgl32.Vec3{0, 1, 0}
	c.Right = mgl32.Vec3.Cross(c.WorldUp, c.Direction).Normalize()
	c.Up = mgl32.Vec3.Cross(c.Direction, c.Right)

	c.Radius = 7
	c.CamX = float32(math.Cos(float64(c.CameraRotation.Y())) * c.Radius)
	c.CamZ = float32(math.Sin(float64(c.CameraRotation.Y())) * c.Radius)
	c.View = mgl32.LookAtV(mgl32.Vec3{c.CamX, 0 + c.Pos.Y(), c.CamZ}, c.Target, c.Up)

	c.Perspective = mgl32.Perspective(mgl32.DegToRad(45), 1280/720, 0.1, 100)

	c.Model = mgl32.Mat4{1}
	c.Model = mgl32.Translate3D(0, -1, 0)

	return c
}

func (c *Camera) Update(window *glfw.Window) {
	dir := mgl32.Vec3{c.Pos.X() - c.Target.X(), c.Pos.Y() - c.Target.Y(), c.Pos.Z() - c.Target.Z()}
	c.Direction = dir.Normalize()
	c.WorldUp = mgl32.Vec3{0, 1, 0}
	c.Right = mgl32.Vec3.Cross(c.WorldUp, c.Direction).Normalize()
	c.Up = mgl32.Vec3.Cross(c.Direction, c.Right)

	c.View = mgl32.LookAtV(c.Pos, c.Target, c.Up)

	c.Perspective = mgl32.Perspective(mgl32.DegToRad(66), 1280/720, 0.1, 100)

	c.Model = mgl32.Mat4{1}
	c.Model = mgl32.Translate3D(0, 0, 0)
	c.CamX = float32(math.Cos(float64(c.CameraRotation.Y()))*c.Radius) * float32(math.Cos(float64(c.CameraRotation.X())))
	c.CamY = float32(math.Sin(float64(c.CameraRotation.X())) * c.Radius)
	c.CamZ = float32(math.Sin(float64(c.CameraRotation.Y()))*c.Radius) * float32(math.Cos(float64(c.CameraRotation.X())))
	c.Target = mgl32.Vec3{c.CamX, c.CamY, c.CamZ}.Add(c.Pos)
	c.View = mgl32.LookAtV(c.Pos, c.Target, c.Up)

	if window.GetKey(glfw.KeyH) == glfw.Press {
		c.CameraRotation = c.CameraRotation.Add(mgl32.Vec3{0, -0.03 * utils.DT, 0})
	}
	if window.GetKey(glfw.KeyL) == glfw.Press {
		c.CameraRotation = c.CameraRotation.Add(mgl32.Vec3{0, 0.03 * utils.DT, 0})
	}

	if window.GetKey(glfw.KeyJ) == glfw.Press {
		c.CameraRotation = c.CameraRotation.Add(mgl32.Vec3{-0.03 * utils.DT, 0, 0})
	}
	if window.GetKey(glfw.KeyK) == glfw.Press {
		c.CameraRotation = c.CameraRotation.Add(mgl32.Vec3{0.03 * utils.DT, 0, 0})
	}

	if window.GetKey(glfw.KeyW) == glfw.Press {
		c.Pos = c.Pos.Add(mgl32.Vec3{c.CamX / 100 * utils.DT, 0, c.CamZ / 100 * utils.DT})
	} else if window.GetKey(glfw.KeyS) == glfw.Press {
		c.Pos = c.Pos.Sub(mgl32.Vec3{c.CamX / 100 * utils.DT, 0, c.CamZ / 100 * utils.DT})
	}

	if window.GetKey(glfw.KeyD) == glfw.Press {
		c.Pos = c.Pos.Add(mgl32.Vec3{c.Right.X() * float32(c.Radius) / 100 * utils.DT, 0, c.Right.Z() * float32(c.Radius) / 100 * utils.DT})
	} else if window.GetKey(glfw.KeyA) == glfw.Press {
		c.Pos = c.Pos.Sub(mgl32.Vec3{c.Right.X() * float32(c.Radius) / 100 * utils.DT, 0, c.Right.Z() * float32(c.Radius) / 100 * utils.DT})
	}

	if window.GetKey(glfw.KeyQ) == glfw.Press {
		c.Pos = c.Pos.Add(mgl32.Vec3{0, 1 * utils.DT, 0})
	} else if window.GetKey(glfw.KeyE) == glfw.Press {
		c.Pos = c.Pos.Add(mgl32.Vec3{0, -1 * utils.DT, 0})
	}
}

var CurrentCamera = NewCamera(mgl32.Vec3{1, 0, 0.0}, mgl32.Vec3{0, 0, 0})
