package renderer

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

// Sprite ...
type Sprite struct {
	Shader  *Shader
	texture *Texture
	// Vertex Array Object is array of buffers which contains data necessary for draw vertices like texture, verticies coords etc.
	vao           *VertexArray
	elementsCount int32
}

// InitSprite - constructor for Sprite
func InitSprite(shader *Shader, texture *Texture, subTextureName string, vertices *[]float32, indexes *[]uint32) (s *Sprite) {
	s = new(Sprite)
	s.vao = InitVertexArray()
	s.Shader = shader
	s.texture = texture
	s.elementsCount = int32(len(*indexes))

	subTexture := s.texture.GetSubTexture(subTextureName)
	texCoords := []float32{
		subTexture.leftBottomXY.X(), subTexture.leftBottomXY.Y(), // bottom left
		subTexture.leftBottomXY.X(), subTexture.rightTopXY.Y(), // top left
		subTexture.rightTopXY.X(), subTexture.rightTopXY.Y(), // top right
		subTexture.rightTopXY.X(), subTexture.leftBottomXY.Y(), // bottom right
	}

	s.vao.Bind()
	vertexCoordsBuffer := InitVertexBuffer(vertices, 3, gl.STATIC_DRAW)
	s.vao.AddBuffer(vertexCoordsBuffer)
	texCoordsBuffer := InitVertexBuffer(&texCoords, 2, gl.STATIC_DRAW)
	s.vao.AddBuffer(texCoordsBuffer)

	ib := InitIndexBuffer(indexes, gl.STATIC_DRAW)

	s.vao.Unbind()
	ib.Unbind()
	return s
}

// Render ...
func (v *Sprite) Render() {
	v.Shader.Use()
	v.texture.Bind()
	v.vao.Bind()
	gl.DrawElements(gl.TRIANGLES, v.elementsCount, gl.UNSIGNED_INT, nil)
}
