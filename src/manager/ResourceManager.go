package manager

import (
	"engine/src/renderer"
	"engine/src/utils"
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	defVertices = []float32{
		-0.5, -0.5, 0.0, // bottom left
		-0.5, 0.5, 0.0, // top left
		0.5, 0.5, 0.0, // top right
		0.5, -0.5, 0.0, // bottom right
	}

	defIndexes = []uint32{ // note that we start from 0!
		0, 1, 2, // first triangle
		2, 3, 0, // first triangle
	}
)

// ResourceManager class
type ResourceManager struct {
	shadersMap  map[string]*renderer.Shader
	texturesMap map[string]*renderer.Texture
	spritesMap  map[string]*renderer.Sprite
}

func NewResourceManager() (rm *ResourceManager) {
	rm = new(ResourceManager)
	rm.shadersMap = make(map[string]*renderer.Shader, utils.Sizeof(renderer.Shader{}))
	rm.texturesMap = make(map[string]*renderer.Texture, utils.Sizeof(renderer.Texture{}))
	rm.spritesMap = make(map[string]*renderer.Sprite, utils.Sizeof(renderer.Sprite{}))
	return rm
}
func (r *ResourceManager) AddShader(name, vertexPath, fragmentPath string) {
	r.shadersMap[name] = renderer.NewShader(vertexPath, fragmentPath)
}
func (r *ResourceManager) GetShader(name string) *renderer.Shader {
	s := r.shadersMap[name]
	if s == nil {
		log.Fatal("Error: shader not found")
	}
	return s
}
func (r *ResourceManager) AddTexture(name, path string, subTextures *[]string, subTexWidth, subTexHeight float32) {
	r.texturesMap[name] = renderer.NewTexture(path, gl.LINEAR_MIPMAP_LINEAR, gl.LINEAR, gl.CLAMP_TO_EDGE)
	texture := r.texturesMap[name]
	xStep, yStep := subTexWidth/texture.Width, (subTexHeight / texture.Height)
	var x, y float32
	for _, val := range *subTextures {
		texture.AddSubTexture(val, mgl32.Vec2{x, y}, mgl32.Vec2{x + xStep, y + yStep})
		if x+xStep == texture.Width {
			x = 0
			if y+yStep < texture.Height {
				y += yStep
			} else {
				log.Fatal("Error: haven't space for subTexture: ", val)
			}
		} else {
			x += xStep
		}
	}
}
func (r *ResourceManager) GetTexture(name string) *renderer.Texture {
	t := r.texturesMap[name]
	if t == nil {
		log.Fatal("Error: texture not found")
	}
	return t
}
func (r *ResourceManager) AddSprite(name, shader, texture, subTexture string) {
	r.spritesMap[name] = renderer.NewSprite(r.shadersMap[shader], r.texturesMap[texture], subTexture, &defVertices, &defIndexes)
}
func (r *ResourceManager) GetSprite(name string) *renderer.Sprite {
	s := r.spritesMap[name]
	if s == nil {
		log.Fatal("Error: shader not found")
	}
	return s
}
