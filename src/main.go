package main

import (
	"engine/src/manager"
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
	defer window.Destroy()
	window.GLCreateContext()
	window.SetResizable(true)
	err = gl.Init()
	if err != nil {
		log.Fatal("Gl init error: ", err)
	}
	rm := manager.NewResourceManager()
	subTextures := []string{"violet", "blue", "brown"}
	rm.AddShader("default", "../res/shaders/defaultVertex.glsl", "../res/shaders/defaultFragment.glsl")
	rm.AddTexture("spirits", "../res/textures/spirits.png", &subTextures, 170, 220)
	rm.AddSprite("unit", "default", "spirits", "blue")
	rm.AddSprite("unitV", "default", "spirits", "violet")
	sprite := rm.GetSprite("unit")
	sprite1 := rm.GetSprite("unitV")

	fmt.Println("OpenGL version:", gl.GoStr(gl.GetString(gl.VERSION)))

	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	gl.ClearColor(0, 0, 0.3, 1)
	timer := sdl.GetTicks()

	run := true
	for run {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				run = false
			case *sdl.WindowEvent:
				if t.Event == sdl.WINDOWEVENT_RESIZED {
					w, h := window.GetSize()
					gl.Viewport(0, 0, w, h)
				}
			}
		}
		gl.Clear(gl.COLOR_BUFFER_BIT)

		m4 := translateMatrix(0.5, -0.5)
		gl.UniformMatrix4fv(gl.GetUniformLocation(sprite.Shader.ID, gl.Str("transform"+"\x00")), 1, false, &m4[0])
		sprite.Render()

		m4 = translateMatrix(-0.5, 0.5)
		gl.UniformMatrix4fv(gl.GetUniformLocation(sprite.Shader.ID, gl.Str("transform"+"\x00")), 1, false, &m4[0])
		sprite.Render()

		m4 = translateMatrix(0.5, 0.5)
		gl.UniformMatrix4fv(gl.GetUniformLocation(sprite1.Shader.ID, gl.Str("transform"+"\x00")), 1, false, &m4[0])
		sprite1.Render()

		m4 = translateMatrix(-0.5, -0.5)
		gl.UniformMatrix4fv(gl.GetUniformLocation(sprite1.Shader.ID, gl.Str("transform"+"\x00")), 1, false, &m4[0])
		sprite1.Render()

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
