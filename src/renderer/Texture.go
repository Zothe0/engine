package renderer

import (
	"image/png"
	"log"
	"os"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Texture class
type Texture struct {
	ID          uint32
	width       uint32
	height      uint32
	subTextures map[string]*subTexture
}

// NewTexture constructor
func NewTexture(path string, textureCount uint32, minFilter, magFilter, wrapMode int32) *Texture {
	texture := new(Texture)
	pixels, w, h := loadImage(path)
	gl.GenTextures(1, &texture.ID)
	texture.Bind()
	gl.ActiveTexture(gl.TEXTURE0 + textureCount)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, wrapMode)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, wrapMode)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, minFilter)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, magFilter)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(w), int32(h), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	texture.width = w
	texture.height = h

	return texture
}

// Bind texture
func (t *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.ID)
}

// AddSubTexture ...
func (t *Texture) AddSubTexture(name string, leftBottomXY, rightTopXY mgl32.Vec2) {
	t.subTextures[name] = &subTexture{
		leftBottomXY: leftBottomXY,
		rightTopXY:   rightTopXY,
	}
}

type subTexture struct {
	leftBottomXY mgl32.Vec2
	rightTopXY   mgl32.Vec2
}

func newSubTexture() *subTexture {
	var instance *subTexture

	return instance
}

func loadImage(path string) ([]byte, uint32, uint32) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	pixels := make([]byte, w*h*4)
	bIndex := 0
	for y := 0; y < int(h); y++ {
		for x := 0; x < int(w); x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[bIndex] = byte(r / 256)
			bIndex++
			pixels[bIndex] = byte(g / 256)
			bIndex++
			pixels[bIndex] = byte(b / 256)
			bIndex++
			pixels[bIndex] = byte(a / 256)
			bIndex++
		}
	}
	if h%2 == 0 {
		for i, j := w*h*2+1, (w*h*2+1)-w*4; i < (h*w*4+1)-w*4; i, j = i+w*4, j-w*4 {
			for x := 0; x < int(w)*4; x++ {
				pixels[i+x], pixels[j+x] = pixels[j+x], pixels[i+x]
			}
		}
	} else {
		middle := (h-1)*w*2 + 1
		for i, j := middle+w*4, middle-w*4; i < (h*w*4+1)-w*4; i, j = i+w*4, j-w*4 {
			for x := 0; x < int(w)*4; x++ {
				pixels[i+x], pixels[j+x] = pixels[j+x], pixels[i+x]
			}
		}
	}
	return pixels, uint32(w), uint32(h)
}
