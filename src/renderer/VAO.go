package renderer

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

// VertexDataObject ...
type VertexDataObject struct {
	// Vertex Array Object is array of buffers which contains data necessary for draw vertices
	vao uint32
}

// LoadBufferObject constructor for VertexDataObject
func LoadBufferObject() (vdo VertexDataObject) {
	gl.GenVertexArrays(1, &vdo.vao)
	return vdo
}

func (v VertexDataObject) addVBO(data []float32, drawMode uint32) {
	v.Bind()
	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), drawMode)
}

// Bind VAO
func (v VertexDataObject) Bind() {
	gl.BindVertexArray(v.vao)
}

// Unbind VAO
func (v VertexDataObject) Unbind() {
	gl.BindVertexArray(0)
}
