package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"rpg/src/renderer"

	"github.com/go-gl/gl/v4.6-core/gl"
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
		0, 1, // top left
		0, 0, // bottom left
		1, 0, // bottom right
		1, 1, // top right
		// 1, 1, // top right
		// 0, 0, // bottom left
	}
	bytes, width, height := LoadImage("../res/textures/flint.png")
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(bytes))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	vertices := []float32{
		-0.5, -0.5, 0.0, // bottom left
		-0.5, 0.5, 0.0, // top left
		0.5, 0.5, 0.0, // top right
		0.5, -0.5, 0.0, // bottom right
	}
	// colors := []float32{
	// 	-0.5, -0.5, 1.0, // bottom left
	// 	-0.5, 0.5, 0.0, // top left
	// 	0.5, 0.5, 1.0, // top right
	// 	0.5, -0.5, 0.0, // bottom right
	// }

	indices := []uint32{ // note that we start from 0!
		0, 1, 2, // first triangle
		2, 3, 0, // first triangle
	}
	// vertices1 := []float32{
	// 	// positions         // colors
	// 	-0.5, -0.5, 0.0, 1.0, 0.0, 0.0, // bottom  left
	// 	0.0, 0.5, 0.0, 0.0, 0.0, 1.0, // top
	// 	0.5, -0.5, 0.0, 0.0, 1.0, 0.0, // bottom right
	// }
	// indices2 := []uint32{ // note that we start from 0!
	// 	0, 1, 2, // first triangle
	// }
	vdo := renderer.LoadVertexDataObject(shader, texture)
	vdo.AddVBO(vertices, 3, gl.STATIC_DRAW, 1)
	vdo.AddVBO(texCoords, 2, gl.STATIC_DRAW, 1)
	vdo.AddEBO(indices, gl.STATIC_DRAW)
	// vdo1 := renderer.LoadVertexDataObject(shader, texture)
	// vdo1.AddVBO(vertices1, gl.STATIC_DRAW, 2)
	// vdo1.AddVBO(texCoords, gl.STATIC_DRAW, 1)
	// vdo1.AddEBO(indices2, gl.STATIC_DRAW)

	var nrAttrbutes int32
	gl.GetIntegerv(gl.MAX_VERTEX_ATTRIBS, &nrAttrbutes)
	fmt.Println("OpenGL version:", gl.GoStr(gl.GetString(gl.VERSION)))
	fmt.Println("Maximum nr of vertex attributes supported: ", nrAttrbutes)

	// projectionMatLocation := gl.GetUniformLocation(vdo1.Shader.ShaderProgram, gl.Str("projectionMat"+"\x00"))
	// gl.UniformMatrix4f()
	// projectionMat := mat4.Zero
	// projectionMat.TranslateX(0.5)

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
func LoadImage(path string) (pixels []byte, width, height int32) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	width = int32(img.Bounds().Max.X)
	height = int32(img.Bounds().Max.Y)
	pixels = make([]byte, width*height*4)
	bIndex := 0
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
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

	return pixels, width, height
}
