package renderer

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

// VertexDataObject ...
type VertexDataObject struct {
	// Vertex Array Object is array of buffers which contains data necessary for draw vertices
	Shader        *Shader
	vao           uint32
	elementsCount int32
	vboCount      uint32
}

// LoadVertexDataObject constructor for VertexDataObject
func LoadVertexDataObject(shader *Shader) *VertexDataObject {
	var vdo VertexDataObject
	vdo.Shader = shader
	gl.GenVertexArrays(1, &vdo.vao)
	return &vdo
}

// AddVBO ...
func (v *VertexDataObject) AddVBO(data []float32, drawMode uint32) {
	v.bind()
	var buffer uint32
	
	gl.GenBuffers(1, &buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), drawMode)

	gl.EnableVertexAttribArray(v.vboCount)
	gl.VertexAttribPointer(v.vboCount, 3, gl.FLOAT, false, 3*4, nil)
	v.vboCount++

	v.unbind()
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// AddEBO ...
func (v *VertexDataObject) AddEBO(data []uint32, drawMode uint32) {
	v.elementsCount = int32(len(data))

	v.bind()
	var buffer uint32
	gl.GenBuffers(1, &buffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(data)*4, gl.Ptr(data), drawMode)

	v.unbind()
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// Render ...
func (v *VertexDataObject) Render() {
	v.Shader.Use()
	v.bind()
	gl.DrawElements(gl.TRIANGLES, v.elementsCount, gl.UNSIGNED_INT, nil)
}

// Bind VAO
func (v *VertexDataObject) bind() {
	gl.BindVertexArray(v.vao)
}

// Unbind VAO
func (v *VertexDataObject) unbind() {
	gl.BindVertexArray(0)
}
