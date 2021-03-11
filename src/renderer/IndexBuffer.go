package renderer

import "github.com/go-gl/gl/v4.6-core/gl"

// IndexBuffer class
type IndexBuffer struct {
	id uint32
}

// NewIndexBuffer constructor
func NewIndexBuffer(data *[]uint32, drawMode uint32) (indexBuffer *IndexBuffer) {
	indexBuffer = new(IndexBuffer)
	gl.GenBuffers(1, &indexBuffer.id)
	indexBuffer.Bind()
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(*data)*4, gl.Ptr(*data), drawMode)
	return indexBuffer
}

// Bind ...
func (i *IndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, i.id)
}

// Unbind ...
func (i *IndexBuffer) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}
