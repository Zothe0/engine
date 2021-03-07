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
		0.5, 0.5, 0.0, // top right
		0.5, -0.5, 0.0, // bottom right
		-0.5, -0.5, 0.0, // bottom let
		-0.5, 0.5, 0.0, // top left
	}

	indices := []uint32{ // note that we start from 0!
		0, 1, 3, // first triangle
		1, 2, 3, // second triangle
	}
	vertices2 := []float32{
		0.5, -1, 0,
		0.5, -0.5, 0,
		1, -0.5, 0,
		1, -1, 0,
	}
	indices2 := []uint32{ // note that we start from 0!
		0, 1, 3, // first triangle
		1, 2, 3, // second triangle
	}
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	var EBO uint32
	gl.GenBuffers(1, &EBO)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	var VAO2 uint32
	gl.GenVertexArrays(1, &VAO2)
	gl.BindVertexArray(VAO2)

	var VBO2 uint32
	gl.GenBuffers(1, &VBO2)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO2)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices2)*4, gl.Ptr(vertices2), gl.STATIC_DRAW)

	var EBO2 uint32
	gl.GenBuffers(1, &EBO2)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO2)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices2)*4, gl.Ptr(indices2), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

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
		gl.ClearColor(0, 0, 0, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// gl.UseProgram(shaderProgram)
		gl.BindVertexArray(VAO)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
		shader.Use()
		gl.BindVertexArray(VAO2)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
		window.GLSwap()
	}
}