package renderer

import "github.com/go-gl/gl/v4.6-core/gl"

// VertexBuffer class
type VertexBuffer struct {
	id uint32
}

// NewVertexBuffer constructor. On load it's stay binded
func NewVertexBuffer(data *[]float32, vectorSize int32, drawMode uint32) (vbo *VertexBuffer) {
	vbo = new(VertexBuffer)

	gl.GenBuffers(1, &vbo.id)
	vbo.Bind()
	gl.BufferData(gl.ARRAY_BUFFER, len(*data)*4, gl.Ptr(*data), drawMode)

	return vbo
}

// Update - by default used to change subTexture coordinates for animation
func (v *VertexBuffer) Update(data *[]float32) {
	v.Bind()
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(*data)*4, gl.Ptr(*data))
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
