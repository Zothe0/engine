package renderer

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

// Sprite ...
type Sprite struct {
	shader        *Shader
	texture       *Texture
	// Vertex Array Object is array of buffers which contains data necessary for draw vertices like texture, verticies coords etc.
	vao           *VertexArray
	elementsCount int32
	vboCount      uint32
}

// LoadSprite constructor for Sprite
func LoadSprite(shader *Shader, texture *Texture) (vdo *Sprite) {
	vdo = new(Sprite)
	vdo.shader = shader
	vdo.texture = texture
	vdo.vao = NewVertexArray()
	return vdo
}

// AddVBO ...
func (v *Sprite) AddVBO(data *[]float32, vectorSize int32, drawMode uint32) {
	v.vao.Bind()

	buffer := NewVertexBuffer(data, vectorSize, drawMode)
	gl.EnableVertexAttribArray(v.vboCount)
	gl.VertexAttribPointer(v.vboCount, vectorSize, gl.FLOAT, false, 0, nil)

	v.vboCount++
	v.vao.Unbind()
	buffer.Unbind()
}

// AddEBO ...
func (v *Sprite) AddEBO(data *[]uint32, drawMode uint32) {
	v.elementsCount = int32(len(*data))
	v.vao.Bind()

	buffer := NewIndexBuffer(data, drawMode)
	// var buffer uint32
	// gl.GenBuffers(1, &buffer)
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffer)
	// gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(*data)*4, gl.Ptr(data), drawMode)

	v.vao.Unbind()
	buffer.Unbind()
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

// Render ...
func (v *Sprite) Render() {
	v.shader.Use()
	v.texture.Bind()
	v.vao.Bind()
	gl.DrawElements(gl.TRIANGLES, v.elementsCount, gl.UNSIGNED_INT, nil)
}
