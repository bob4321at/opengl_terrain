package world

import (
	"unsafe"

	"terrain/camera"
	"terrain/fastnoise"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var noise = fastnoise.New[float32]()

func init() {
	noise.NoiseType(fastnoise.OpenSimplex2S)
}

type Terrain struct {
	Verticies []float32
	Indicies  []uint32
	Color     []float32
	Pos       mgl32.Vec3
	Model     mgl32.Mat4
	Vao       uint32
	Ebo       uint32
}

func NewTerrain() Terrain {
	return Terrain{}
}

func (terrain *Terrain) GenVertecies(w, h int) {
	index_verts := []float32{}

	for i := range w / 3 {
		for j := range h / 3 {
			height := noise.Noise2D(i, j)
			height *= 10
			nextheight := noise.Noise2D(i+1, j)
			nextheight *= 10
			// nextnextheight := noise.Noise2D(i, j+1)
			// nextnextheight *= 10
			normal := mgl32.Vec3{float32(i), height, float32(j)}.Cross(mgl32.Vec3{float32(i) + 1, nextheight, float32(j)}).Normalize()
			terrain.Verticies = append(terrain.Verticies, float32(i), height-100, float32(j))
			index_verts = append(index_verts, float32(i), height-100, float32(j))
			terrain.Verticies = append(terrain.Verticies, normal.X(), normal.Y(), normal.Z())
			// terrain.Verticies = append(terrain.Verticies, float32(i), height, float32(j))
			// terrain.Verticies = append(terrain.Verticies, normal.X(), normal.Y(), normal.Z())
			// terrain.Verticies = append(terrain.Verticies, float32(i), nextnextheight, float32(j+1))
			// terrain.Verticies = append(terrain.Verticies, normal.X(), normal.Y(), normal.Z())
		}
	}

	for v := range len(index_verts) {
		if v/w < w && v/h < h {
			terrain.Indicies = append(terrain.Indicies, uint32(v), uint32(v+1), uint32(v+w))
			// terrain.Indicies = append(terrain.Indicies, uint32(v+w+2), uint32(v+w+1), uint32(v+1))
		}
	}

}

func (terrain *Terrain) Draw(program uint32) {
	gl.UseProgram(program)

	modelUniformLoc := gl.GetUniformLocation(program, gl.Str("model\x00"))
	viewUniformLoc := gl.GetUniformLocation(program, gl.Str("view\x00"))
	projectionUniformLoc := gl.GetUniformLocation(program, gl.Str("projection\x00"))

	terrain.Model = mgl32.Translate3D(terrain.Pos.X(), terrain.Pos.Y(), terrain.Pos.Z())

	gl.UniformMatrix4fv(modelUniformLoc, 1, false, &terrain.Model[0])
	gl.UniformMatrix4fv(viewUniformLoc, 1, false, &camera.CurrentCamera.View[0])
	gl.UniformMatrix4fv(projectionUniformLoc, 1, false, &camera.CurrentCamera.Perspective[0])

	gl.BindVertexArray(terrain.Vao)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, terrain.Ebo)
	gl.DrawElements(gl.TRIANGLES, int32(len(terrain.Indicies)), gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)
}

func (terrain *Terrain) GenVao() {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(terrain.Verticies), gl.Ptr(terrain.Verticies), gl.STATIC_DRAW)

	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(terrain.Indicies), unsafe.Pointer(&terrain.Indicies[0]), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, unsafe.Pointer(nil))
	gl.EnableVertexArrayAttrib(vao, 0)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, unsafe.Pointer(uintptr(3*4)))
	gl.EnableVertexArrayAttrib(vao, 1)

	terrain.Vao = vao
	terrain.Ebo = ebo
}
