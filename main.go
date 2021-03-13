package main

import (
	"engine/src/manager"
	"fmt"
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
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
	rm.AddShader("default", "./res/shaders/defaultVertex.glsl", "./res/shaders/defaultFragment.glsl")
	rm.AddShader("color", "./res/shaders/colorV.glsl", "./res/shaders/colorF.glsl")
	rm.AddTexture("spirits", "./res/textures/spirits.png", &subTextures, 170, 220)
	rm.AddSprite("unit", "color", "spirits", "blue")
	rm.AddSprite("unitV", "default", "spirits", "violet")
	sprite := rm.GetSprite("unit")
	// sprite1 := rm.GetSprite("unitV")

	fmt.Println("OpenGL version:", gl.GoStr(gl.GetString(gl.VERSION)))

	projection := mgl.Perspective(mgl.DegToRad(60), width/height, 0.1, 100)
	view := mgl.Translate3D(0, 0, -3).Mul4(mgl.Ident4())
	size := mgl.Vec3{2, 2, 2}
	model := mgl.Ident4()
	model = model.Mul4(mgl.Translate3D(2, 2, -10))
	model = model.Mul4(mgl.Scale3D(size.X(), size.Y(), size.Z()))
	model = model.Mul4(mgl.HomogRotate3DY(mgl.DegToRad(90)))

	gl.ClearColor(0, 0, 0.3, 1)
	gl.Enable(gl.DEPTH_TEST) // For 3D correct drawn
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
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
					width, height = float32(w), float32(h)
					gl.Viewport(0, 0, w, h)
				}
			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN {
					if t.Keysym.Scancode == sdl.SCANCODE_W {
						view = view.Mul4(mgl.Translate3D(0, 0, 1))
					}
					if t.Keysym.Scancode == sdl.SCANCODE_A {
						view = view.Mul4(mgl.Translate3D(1, 0, 0))
					}
					if t.Keysym.Scancode == sdl.SCANCODE_S {
						view = view.Mul4(mgl.Translate3D(0, 0, -1))
					}
					if t.Keysym.Scancode == sdl.SCANCODE_D {
						view = view.Mul4(mgl.Translate3D(-1, 0, 0))
					}
					if t.Keysym.Scancode == sdl.SCANCODE_LSHIFT || t.Keysym.Scancode == sdl.SCANCODE_C {
						view = view.Mul4(mgl.Translate3D(0, 1, 0))
					}
					if t.Keysym.Scancode == sdl.SCANCODE_SPACE {
						view = view.Mul4(mgl.Translate3D(0, -1, 0))
					}
				}
			}
		}
		// Timer
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		updateTimer := sdl.GetTicks()
		if updateTimer-timer > 10 {
			model = model.Mul4(mgl.HomogRotate3DY(mgl.DegToRad(1)))
			timer = updateTimer
		}
		// Upload data to shader
		sprite.Shader.SetMat4("projection", projection)
		sprite.Shader.SetMat4("view", view)
		sprite.Shader.SetMat4("model", model)
		// Draw sprite
		sprite.Render()
		window.GLSwap()
	}
}
