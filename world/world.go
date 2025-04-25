package world

import (
	"math"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Terrain struct {
	Verticies []float32
	Color     []float32
	Vao       uint32
}

func NewTerrain() Terrain {
	return Terrain{}
}

func (terrain *Terrain) GenVertecies(w, h int) {
	for i := range w {
		for j := range h {
			terrain.Verticies = append(terrain.Verticies, float32(i), float32(math.Sin(math.Cos(float64(i+j)))), float32(j))
			terrain.Verticies = append(terrain.Verticies, 0, 0, 0)
			terrain.Verticies = append(terrain.Verticies, float32(i)+1, float32(math.Sin(math.Cos(float64(i+j)))), float32(j))
			terrain.Verticies = append(terrain.Verticies, 0, 0, 0)
			terrain.Verticies = append(terrain.Verticies, float32(i), float32(math.Sin(math.Cos(float64(i+j)))), float32(j+1))
			terrain.Verticies = append(terrain.Verticies, 0, 0, 0)
		}
	}
	for i := range w {
		for j := range h {
			terrain.Verticies = append(terrain.Verticies, float32(i), float32(math.Sin(math.Cos(float64(i+j)))), float32(j)+1)
			terrain.Verticies = append(terrain.Verticies, 0, 0, 0)
			terrain.Verticies = append(terrain.Verticies, float32(i)+1, float32(math.Sin(math.Cos(float64(i+j)))), float32(j)+1)
			terrain.Verticies = append(terrain.Verticies, 0, 0, 0)
			terrain.Verticies = append(terrain.Verticies, float32(i)+1, float32(math.Sin(math.Cos(float64(i+j)))), float32(j))
			terrain.Verticies = append(terrain.Verticies, 0, 0, 0)
		}
	}
}

func (terrain *Terrain) GenVao() {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(terrain.Verticies), gl.Ptr(terrain.Verticies), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexArrayAttrib(vao, 0)
	gl.EnableVertexArrayAttrib(vao, 1)
	gl.BindBuffer(gl.ARRAY_BUFFER, vao)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, unsafe.Pointer(nil))
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, unsafe.Pointer(uintptr(3*4)))

	terrain.Vao = vao
}
