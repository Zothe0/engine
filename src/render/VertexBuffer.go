package render

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

// VertexBuffer class
type VertexBuffer struct {
	id         uint32
	vectorSize int32
	size       int
}

// InitVertexBuffer - constructor.
func InitVertexBuffer(data *[]float32, vectorSize int32, drawMode uint32) (vb *VertexBuffer) {
	vb = new(VertexBuffer)
	vb.size = len(*data) * 4
	vb.vectorSize = vectorSize
	gl.GenBuffers(1, &vb.id)
	vb.Bind()
	gl.BufferData(gl.ARRAY_BUFFER, vb.size, gl.Ptr(*data), drawMode)
	vb.Unbind()
	return vb
}

// Update - by default used to change subTexture coordinates for animation
func (v *VertexBuffer) Update(data *[]float32) {
	v.Bind()
	v.size = len(*data) * 4
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, v.size, gl.Ptr(*data))
	v.Unbind()
}

// Bind ...
func (v *VertexBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, v.id)
}

// Unbind ...
func (v *VertexBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}
