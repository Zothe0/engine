package renderer

import "github.com/go-gl/gl/v4.6-core/gl"

// VertexArray class
type VertexArray struct {
	id uint32
}

// NewVertexArray constructor
func NewVertexArray() (vertexArray *VertexArray) {
	vertexArray = new(VertexArray)
	gl.GenVertexArrays(1, &vertexArray.id)
	return vertexArray
}

// Bind vertex array
func (v *VertexArray) Bind() {
	gl.BindVertexArray(v.id)
}

// Unbind vertex array
func (v *VertexArray) Unbind() {
	gl.BindVertexArray(0)
}
