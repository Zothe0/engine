package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"rpg/src/renderer"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("RPG", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	window.GLCreateContext()
	defer window.Destroy()
	err = gl.Init()
	if err != nil {
		log.Fatal("Gl init error: ", err)
	}
	shader := renderer.LoadShader("../res/shaders/defaultVertex.glsl", "../res/shaders/defaultFragment.glsl")
	texCoords := []float32{
		0, 0, // bottom left
		0, 1, // top left
		1, 1, // top right
		1, 0, // bottom right
	}
	bytes, width, height := LoadImage("../res/textures/flint.png")
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(width), int32(height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(bytes))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	vertices := []float32{
		-0.5, -0.5, 0.0, // bottom left
		-0.5, 0.5, 0.0, // top left
		0.5, 0.5, 0.0, // top right
		0.5, -0.5, 0.0, // bottom right
	}

	indices := []uint32{ // note that we start from 0!
		0, 1, 2, // first triangle
		2, 3, 0, // first triangle
	}
	vdo := renderer.LoadVertexDataObject(shader, texture)
	vdo.AddVBO(vertices, 3, gl.STATIC_DRAW, 1)
	vdo.AddVBO(texCoords, 2, gl.STATIC_DRAW, 1)
	vdo.AddEBO(indices, gl.STATIC_DRAW)

	fmt.Println("OpenGL version:", gl.GoStr(gl.GetString(gl.VERSION)))

	shader.Use()
	trans := mgl32.Ident4()
	// trans = mgl32.HomogRotate3DZ(mgl32.DegToRad(45)).Mul4(trans)
	trans = mgl32.Translate3D(0.5, -0.5, 0).Mul4(trans)
	m4 := [16]float32(trans)
	transformLoc := gl.GetUniformLocation(shader.ShaderProgram, gl.Str("transform"+"\x00"))
	gl.UniformMatrix4fv(transformLoc, 1, false, &m4[0])

	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	gl.ClearColor(0, 0, 0.3, 1)
	timer := sdl.GetTicks()
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		gl.Clear(gl.COLOR_BUFFER_BIT)

		vdo.Render()
		// vdo1.Render()
		updateTimer := sdl.GetTicks()
		if updateTimer-timer > 1000 {
			timer = updateTimer
		}
		window.GLSwap()
	}
}

// LoadImage ...
func LoadImage(path string) (pixels []byte, w, h int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	w = img.Bounds().Max.X
	h = img.Bounds().Max.Y
	pixels = make([]byte, w*h*4)
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
	return pixels, w, h
}
