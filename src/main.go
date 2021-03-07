package main

import (
	"fmt"
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

	shader := renderer.LoadShader("../res/shaders/defaultVertex.glsl", "../res/shaders/defaultFragment.glsl")

	vertices := []float32{
		-0.5, -0.5, 0.0, // bottom left
		-0.5, 0.5, 0.0, // top left
		0.5, 0.5, 0.0, // top right
		0.5, -0.5, 0.0, // bottom right
	}

	indices := []uint32{ // note that we start from 0!
		0, 1, 2, // first triangle
		0, 2, 3, // first triangle
	}
	// vertices1 := []float32{
	// 	// positions         // colors
	// 	0.5, -0.5, 0.0, 1.0, 0.0, 0.0, // bottom right
	// 	-0.5, -0.5, 0.0, 0.0, 1.0, 0.0, // bottom left
	// 	0.0, 0.5, 0.0, 0.0, 0.0, 1.0, // top
	// }
	// indices2 := []uint32{ // note that we start from 0!
	// 	0, 1, 3, // first triangle
	// 	1, 2, 3, // second triangle
	// }
	vdo := renderer.LoadVertexDataObject(shader)
	vdo.AddVBO(vertices, gl.STATIC_DRAW)
	vdo.AddEBO(indices, gl.STATIC_DRAW)
	// vdo1 := renderer.LoadVertexDataObject(shader)
	// vdo1.AddVBO(vertices1, gl.STATIC_DRAW)
	// vdo1.AddEBO(indices2, gl.STATIC_DRAW)

	var nrAttrbutes int32
	gl.GetIntegerv(gl.MAX_VERTEX_ATTRIBS, &nrAttrbutes)
	fmt.Println("OpenGL version:", gl.GoStr(gl.GetString(gl.VERSION)))
	fmt.Println("Maximum nr of vertex attributes supported: ", nrAttrbutes)

	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	gl.ClearColor(0, 0, 0.3, 1)
	timer := sdl.GetTicks()
	// vertexColorLocation := gl.GetUniformLocation(vdo.Shader.ShaderProgram, gl.Str("ourColor"+"\x00"))
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
