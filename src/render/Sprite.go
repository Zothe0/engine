package render

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
)

// Sprite ...
type Sprite struct {
	Shader  *Shader
	texture *Texture
	// Vertex Array Object is array of buffers which contains data necessary for draw vertices like texture, verticies coords etc.
	vao           *VertexArray
	elementsCount int32
	position      mgl.Vec3
	size          mgl.Vec3
	rotation      mgl.Vec3
}

// InitSprite - constructor for Sprite
func InitSprite(shader *Shader, texture *Texture, subTextureName string, vertices, colors *[]float32, indexes *[]uint32) (s *Sprite) {
	s = new(Sprite)
	s.vao = InitVertexArray()
	s.Shader = shader
	s.texture = texture
	s.elementsCount = int32(len(*indexes))

	// subTexture := s.texture.GetSubTexture(subTextureName)
	// lbx, lby := subTexture.leftBottomXY.X(), subTexture.leftBottomXY.Y()
	// rtx, rty := subTexture.rightTopXY.X(), subTexture.rightTopXY.Y()
	// texCoords := []float32{
	// 	lbx, lby, // bottom left
	// 	lbx, rty, // top left
	// 	rtx, rty, // top right
	// 	rtx, lby, // bottom right

	// 	lbx, lby, // bottom left
	// 	lbx, rty, // top left
	// 	rtx, rty, // top right
	// 	rtx, lby, // bottom right
	// }

	s.vao.Bind()
	vertexCoordsBuffer := InitVertexBuffer(vertices, 3, gl.DYNAMIC_DRAW)
	s.vao.AddBuffer(vertexCoordsBuffer)
	texCoordsBuffer := InitVertexBuffer(colors, 3, gl.STATIC_DRAW)
	s.vao.AddBuffer(texCoordsBuffer)

	ib := InitIndexBuffer(indexes, gl.STATIC_DRAW)

	s.vao.Unbind()
	ib.Unbind()
	return s
}

// Render ...
func (v *Sprite) Render() {
	v.Shader.Use()
	if v.texture != nil {
		v.texture.Bind()
	}
	v.vao.Bind()

	gl.DrawElements(gl.TRIANGLES, v.elementsCount, gl.UNSIGNED_INT, nil)
}
