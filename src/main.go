package main

import (
	"engine/src/renderer"
	"fmt"
	"log"

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
	shader := renderer.NewShader("../res/shaders/defaultVertex.glsl", "../res/shaders/defaultFragment.glsl")
	texCoords := []float32{
		0, 0, // bottom left
		0, 1, // top left
		1, 1, // top right
		1, 0, // bottom right
	}
	texture := renderer.NewTexture("../res/textures/flint.png", gl.LINEAR_MIPMAP_LINEAR, gl.NEAREST, gl.CLAMP_TO_EDGE)
	texture1 := renderer.NewTexture("../res/textures/test.png", gl.LINEAR_MIPMAP_LINEAR, gl.NEAREST, gl.CLAMP_TO_EDGE)

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
	texture.Bind()
	vdo.AddVBO(vertices, 3, gl.STATIC_DRAW, 1)
	vdo.AddVBO(texCoords, 2, gl.STATIC_DRAW, 1)
	vdo.AddEBO(indices, gl.STATIC_DRAW)
	
	vdo1 := renderer.LoadVertexDataObject(shader, texture1)
	vdo1.AddVBO(vertices, 3, gl.STATIC_DRAW, 1)
	vdo1.AddVBO(texCoords, 2, gl.STATIC_DRAW, 1)
	vdo1.AddEBO(indices, gl.STATIC_DRAW)

	fmt.Println("OpenGL version:", gl.GoStr(gl.GetString(gl.VERSION)))

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

		m4 := translateMatrix(0.5, -0.5)
		gl.UniformMatrix4fv(gl.GetUniformLocation(shader.ID, gl.Str("transform"+"\x00")), 1, false, &m4[0])
		vdo.Render()

		m4 = translateMatrix(-0.5, 0.5)
		gl.UniformMatrix4fv(gl.GetUniformLocation(shader.ID, gl.Str("transform"+"\x00")), 1, false, &m4[0])
		vdo1.Render()

		// Timer
		updateTimer := sdl.GetTicks()
		if updateTimer-timer > 1000 {
			timer = updateTimer
		}
		window.GLSwap()
	}
}

func translateMatrix(x, y float32) *[16]float32 {
	matrix := mgl32.Ident4()
	matrix = mgl32.Translate3D(x, y, 0).Mul4(matrix)
	m4 := [16]float32(matrix)
	return &m4
}
