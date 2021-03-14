package renderer

import "github.com/go-gl/gl/v4.1-core/gl"

// VertexArray class
type VertexArray struct {
	id          uint32
	bufferCount uint32
}

// InitVertexArray - constructor
func InitVertexArray() (v *VertexArray) {
	v = new(VertexArray)
	gl.GenVertexArrays(1, &v.id)
	return v
}
func (v *VertexArray) AddBuffer(vb *VertexBuffer) {
	vb.Bind()

	gl.EnableVertexAttribArray(v.bufferCount)
	gl.VertexAttribPointer(v.bufferCount, vb.vectorSize, gl.FLOAT, false, 0, nil)

	v.bufferCount++
}

// Bind vertex array
func (v *VertexArray) Bind() {
	gl.BindVertexArray(v.id)
}

// Unbind vertex array
func (v *VertexArray) Unbind() {
	gl.BindVertexArray(0)
}
