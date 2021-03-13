package main

import (
	"engine/src/manager"
	"engine/src/utils"
	"fmt"
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	width  float32 = 800
	height float32 = 800
)

func main() {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(width), int32(height), sdl.WINDOW_OPENGL)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()
	window.GLCreateContext()
	window.SetResizable(true)
	err = gl.Init()
	if err != nil {
		log.Fatal("Gl init error: ", err)
	}
	rm := manager.InitResourceManager()
	subTextures := []string{"violet", "blue", "brown"}
	rm.AddShader("default", "../res/shaders/defaultVertex.glsl", "../res/shaders/defaultFragment.glsl")
	rm.AddTexture("spirits", "../res/textures/spirits.png", &subTextures, 170, 220)
	rm.AddSprite("unit", "default", "spirits", "blue")
	rm.AddSprite("unitV", "default", "spirits", "violet")
	sprite := rm.GetSprite("unit")
	// sprite1 := rm.GetSprite("unitV")

	fmt.Println("OpenGL version:", gl.GoStr(gl.GetString(gl.VERSION)))

	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	gl.ClearColor(0, 0, 0.3, 1)
	timer := sdl.GetTicks()
	modelLoc := gl.GetUniformLocation(sprite.Shader.ID, utils.Cstr("model"))
	viewLoc := gl.GetUniformLocation(sprite.Shader.ID, utils.Cstr("view"))
	projectionLoc := gl.GetUniformLocation(sprite.Shader.ID, utils.Cstr("projection"))

	projection := mgl32.Perspective(mgl32.DegToRad(45), width/height, 0.1, 100)
	view := mgl32.Translate3D(0, 0, -3).Mul4(mgl32.Ident4())
	model := mgl32.Translate3D(-1, -1, -10).Mul4(mgl32.Scale3D(0.5, 0.5, 0.5)).Mul4(mgl32.Ident4())

	run := true
	for run {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				run = false
			case *sdl.WindowEvent:
				if t.Event == sdl.WINDOWEVENT_RESIZED {
					w, h := window.GetSize()
					width, height = float32(w), float32(h)
					gl.Viewport(0, 0, w, h)
				}
			}
		}
		gl.Clear(gl.COLOR_BUFFER_BIT)
		// Upload data to shader
		gl.UniformMatrix4fv(projectionLoc, 1, false, utils.MatAddress(projection))
		gl.UniformMatrix4fv(viewLoc, 1, false, utils.MatAddress(view))
		gl.UniformMatrix4fv(modelLoc, 1, false, utils.MatAddress(model))
		// Draw sprite
		sprite.Render()

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
