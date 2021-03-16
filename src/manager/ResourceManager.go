package manager

import (
	"engine/src/render"
	"engine/src/utils"
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
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
	cubeVertices = []float32{
		-0.5, -0.5, 0.5, // bottom left
		-0.5, 0.5, 0.5, // top left
		0.5, 0.5, 0.5, // top right
		0.5, -0.5, 0.5, // bottom right

		-0.5, -0.5, -0.5, // bottom left
		-0.5, 0.5, -0.5, // top left
		0.5, 0.5, -0.5, // top right
		0.5, -0.5, -0.5, // bottom right
	}
	cubeColors = []float32{
		1, 0.5, 0.5, // bottom left
		0.5, 1, 0.5, // top left
		0.5, 0.5, 1, // top right
		0.5, 0, 1, // bottom right

		0.5, 0.5, 1, // bottom left
		0.5, 1, 0.5, // top left
		1, 0.5, 0.5, // top right
		0.5, 0, 0.5, // bottom right
	}
	cubeIndexes = []uint32{ // note that we start from 0!
		// front
		0, 1, 2,
		2, 3, 0,
		// left
		0, 1, 5,
		5, 4, 0,
		// bottom
		0, 4, 7,
		7, 3, 0,
		// top
		6, 2, 1,
		1, 5, 6,
		// back
		6, 2, 3,
		3, 7, 6,
		// right
		6, 5, 4,
		4, 7, 6,
	}
)

// ResourceManager class
type ResourceManager struct {
	shadersMap  map[string]*render.Shader
	texturesMap map[string]*render.Texture
	spritesMap  map[string]*render.Sprite
}

// InitResourceManager - constructor
func InitResourceManager() (rm *ResourceManager) {
	rm = new(ResourceManager)
	rm.shadersMap = make(map[string]*render.Shader, utils.Sizeof(render.Shader{}))
	rm.texturesMap = make(map[string]*render.Texture, utils.Sizeof(render.Texture{}))
	rm.spritesMap = make(map[string]*render.Sprite, utils.Sizeof(render.Sprite{}))
	return rm
}
func (r *ResourceManager) AddShader(name, vertexPath, fragmentPath string) {
	r.shadersMap[name] = render.InitShader(vertexPath, fragmentPath)
}
func (r *ResourceManager) GetShader(name string) *render.Shader {
	s := r.shadersMap[name]
	if s == nil {
		log.Fatal("Error: shader not found")
	}
	return s
}
func (r *ResourceManager) AddTexture(name, path string, subTextures *[]string, subTexWidth, subTexHeight float32) {
	r.texturesMap[name] = render.InitTexture(path, gl.LINEAR_MIPMAP_LINEAR, gl.LINEAR, gl.CLAMP_TO_EDGE)
	texture := r.texturesMap[name]
	xStep, yStep := subTexWidth/texture.Width, (subTexHeight / texture.Height)
	var x, y float32
	for _, val := range *subTextures {
		texture.AddSubTexture(val, mgl.Vec2{x, y}, mgl.Vec2{x + xStep, y + yStep})
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
func (r *ResourceManager) GetTexture(name string) *render.Texture {
	t := r.texturesMap[name]
	if t == nil {
		log.Fatal("Error: texture not found")
	}
	return t
}
func (r *ResourceManager) AddSprite(name, shader, texture, subTexture string) *render.Sprite {
	r.spritesMap[name] = render.InitSprite(r.shadersMap[shader], r.texturesMap[texture], subTexture, &cubeVertices, &cubeColors, &cubeIndexes)
	return r.spritesMap[name]
}
func (r *ResourceManager) GetSprite(name string) *render.Sprite {
	s := r.spritesMap[name]
	if s == nil {
		log.Fatal("Error: shader not found")
	}
	return s
}
