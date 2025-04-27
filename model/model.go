package model

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"terrain/camera"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func RemoveArrayElement[T any](index_to_remove int, slice *[]T) {
	*slice = append((*slice)[:index_to_remove], (*slice)[index_to_remove+1:]...)
}

func getVertsFromFile(file_string string) (verts map[int][]float32) {
	verts = make(map[int][]float32)
	origonal_verts, _, _ := strings.Cut(file_string, "vn ")
	verticies := strings.Split(origonal_verts, "v ")
	RemoveArrayElement(0, &verticies)
	new_test := ""
	for _, thing := range verticies {
		thing = strings.ReplaceAll(thing, "\n", " ")
		new_test += thing
	}
	testss := strings.Split(new_test, " ")

	for vert, element := range testss {
		if element != "" {
			val, err := strconv.ParseFloat(element, 32)
			if err != nil {
				log.Fatal(err)
			}
			verts[int(vert/3)] = append(verts[vert/3], float32(val))
		}
	}

	return verts
}

func getNormsFromFile(file_string string) (norms map[int][]float32) {
	norms = make(map[int][]float32)
	origonal_normals, _, _ := strings.Cut(file_string, "s ")
	_, origonal_normals, _ = strings.Cut(origonal_normals, "vn ")
	origonal_normals, _, _ = strings.Cut(origonal_normals, "vt ")
	origonal_normals = strings.ReplaceAll(origonal_normals, "vn ", "")
	normals := strings.Split(origonal_normals, "v ")
	new_normals := ""
	for _, thing := range normals {
		thing = strings.ReplaceAll(thing, "\n", " ")
		new_normals += thing
	}
	real_normals := strings.Split(new_normals, " ")

	for norm, element := range real_normals {
		if element != "" {
			val, err := strconv.ParseFloat(element, 32)
			if err != nil {
				log.Fatal(err)
			}
			norms[norm/3] = append(norms[norm/3], float32(val))
		}
	}

	return norms
}

type Face struct {
	ID       int
	Indicies []uint32
	Normals  []int
}

func GetIndiciesForFaces(faces *[]Face, file_string string) {
	_, origonal_indicies, _ := strings.Cut(file_string, "s 0")
	indicie := strings.Split(origonal_indicies, "f ")
	combined := ""
	for _, val := range indicie {
		val = strings.ReplaceAll(val, "\n", " ")
		combined += val
	}
	indicie = strings.Split(combined, " ")
	temp_indicie := []string{}
	for _, val := range indicie {
		val, _, _ = strings.Cut(val, "/")
		temp_indicie = append(temp_indicie, val)
	}
	indicie = temp_indicie

	for f := range len(indicie) / 4 {
		*faces = append(*faces, Face{f, []uint32{}, []int{}})
	}

	for f, element := range indicie {
		if element != "" {
			val, err := strconv.Atoi(element)
			if err != nil {
				log.Fatal(err)
			}
			(*faces)[f/5].Indicies = append((*faces)[f/5].Indicies, uint32(val))
		}
	}
}

func getFacesFromFile(file_string string) (faces []Face) {
	GetIndiciesForFaces(&faces, file_string)

	_, origonal_Normals, _ := strings.Cut(file_string, "s 0")
	Normal := strings.Split(origonal_Normals, "f ")
	combined := ""
	for _, val := range Normal {
		val = strings.ReplaceAll(val, "\n", " ")
		combined += val
	}
	Normal = strings.Split(combined, " ")
	temp_Normal := []string{}
	for _, val := range Normal {
		_, val, _ = strings.Cut(val, "/")
		temp_Normal = append(temp_Normal, val)
	}
	temp_temp_Normal := []string{}
	for _, val := range temp_Normal {
		_, val, _ = strings.Cut(val, "/")
		temp_temp_Normal = append(temp_temp_Normal, val)
	}

	Normal = temp_temp_Normal

	for f, element := range Normal {
		if element != "" {
			val, err := strconv.Atoi(element)
			if err != nil {
				log.Fatal(err)
			}
			faces[f/5].Normals = append(faces[f/5].Normals, int(val))
		}
	}

	return faces
}

func ObjParser(path string) (verts map[int][]float32, norms map[int][]float32, faces []Face) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	file_string := string(file)

	verts = getVertsFromFile(file_string)
	norms = getNormsFromFile(file_string)
	faces = getFacesFromFile(file_string)

	return verts, norms, faces
}

type Model struct {
	Verticies []float32
	Normals   []float32
	Indicies  []uint32
	Vao       uint32
	Ebo       uint32
}

func NewModel(path string) (m Model) {
	verts, norms, faces := ObjParser("./model/test.obj")

	indicies := []uint32{}

	temp_verts := []float32{}
	for _, face := range faces {
		for i := range len(face.Indicies) {
			posv := verts[int(face.Indicies[i])-1]
			posn := norms[face.Normals[i]-1]

			xv := posv[0]
			xn := posn[0]
			zv := posv[1]
			zn := posn[1]
			yv := posv[2]
			yn := posn[2]

			temp_verts = append(temp_verts, xv, xn, yv, yn, zv, zn)
			indicies = append(indicies, uint32(i))
		}
	}

	m.Verticies = temp_verts
	m.Indicies = indicies

	fmt.Println(m.Verticies)

	return m
}

func (model *Model) GenVao() {

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(model.Verticies), gl.Ptr(model.Verticies), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(model.Indicies), unsafe.Pointer(&model.Indicies[0]), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, unsafe.Pointer(nil))
	gl.EnableVertexArrayAttrib(vao, 0)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, unsafe.Pointer(uintptr(3*4)))
	gl.EnableVertexArrayAttrib(vao, 1)

	model.Vao = vao
	model.Ebo = ebo
}

func (model *Model) Draw(program uint32) {
	gl.UseProgram(program)

	modelUniformLoc := gl.GetUniformLocation(program, gl.Str("model\x00"))
	viewUniformLoc := gl.GetUniformLocation(program, gl.Str("view\x00"))
	projectionUniformLoc := gl.GetUniformLocation(program, gl.Str("projection\x00"))

	gl.UniformMatrix4fv(modelUniformLoc, 1, false, &camera.CurrentCamera.Model[0])
	gl.UniformMatrix4fv(viewUniformLoc, 1, false, &camera.CurrentCamera.View[0])
	gl.UniformMatrix4fv(projectionUniformLoc, 1, false, &camera.CurrentCamera.Perspective[0])

	gl.BindVertexArray(model.Vao)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, model.Ebo)
	gl.DrawElements(gl.TRIANGLES, int32(len(model.Verticies)), gl.UNSIGNED_INT, nil)
	// gl.DrawArrays(gl.TRIANGLES, 0, int32(len(model.Verticies)))
	gl.BindVertexArray(0)
}

// func (model *Model) Draw(program uint32) {
// 	gl.UseProgram(program)

// 	// Check for duplicate indices
// 	indices := make(map[int]Vertex)
// 	for _, index := range model.Indicies {
// 		if _, ok := indices[index]; !ok {
// 			indices[index] = model.Vertices[0]
// 		} else {
// 			// Handle duplicate index error
// 			log.Fatal("Duplicate index found in model")
// 		}
// 	}

// 	// Map indices to vertex data
// 	indicesData := make([]int, 0)
// 	for _, index := range model.Indicies {
// 		indicesData = append(indicesData, index)
// 	}

// 	gl.UniformMatrix4fv(modelUniformLoc, 1, false, &camera.CurrentCamera.Model[0])
// 	gl.UniformMatrix4fv(viewUniformLoc, 1, false, &camera.CurrentCamera.View[0])
// 	gl.UniformMatrix4fv(projectionUniformLoc, 1, false, &camera.CurrentCamera.Perspective[0])

// 	gl.BindVertexArray(model.Vao)
// 	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, gl.NewBuffer(gl.ELEMENT_ARRAY_BUFFER))

// 	// Copy indices data to buffer
// 	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicesData)*4, indicesData[:], gl.STATIC_DRAW)

// 	gl.DrawElements(gl.TRIANGLES, len(indicesData), gl.UNSIGNED_INT, nil)
// 	gl.BindVertexArray(0)
// }
