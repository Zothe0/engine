package main

import (
	"log"
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

	shader := renderer.LoadShader("../res/shaders/vertexShader.txt", "../res/shaders/fragmentShader.txt")

	vertices := []float32{
		-0.5, -0.5, 0.0, // bottom left
		-0.5, 0.5, 0.0, // top left
		0.5, 0.5, 0.0, // top right
		0.5, -0.5, 0.0, // bottom right
	}

	indices := []uint32{ // note that we start from 0!
		0, 1, 2, // first triangle
		2, 3, 1, // second triangle
		0, 1, 3, // first triangle
	}
	// vertices2 := []float32{
	// 	0.5, -1, 0,
	// 	0.5, -0.5, 0,
	// 	1, -0.5, 0,
	// 	1, -1, 0,
	// }
	// indices2 := []uint32{ // note that we start from 0!
	// 	0, 1, 3, // first triangle
	// 	1, 2, 3, // second triangle
	// }
	vdo := renderer.LoadVertexDataObject(shader)
	vdo.AddVBO(vertices, gl.STATIC_DRAW)
	vdo.AddEBO(indices, gl.STATIC_DRAW)

	gl.ClearColor(0, 0, 0, 1)
	log.Print("OpenGL version: ", gl.GoStr(gl.GetString(gl.VERSION)))
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		gl.Clear(gl.COLOR_BUFFER_BIT)

		vdo.Render()
		window.GLSwap()
	}
}
