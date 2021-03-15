package main

import (
	"engine/src/manager"
	"fmt"
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	width          float32 = 800
	height         float32 = 600
	cameraPosition mgl.Vec3
	view           mgl.Mat4
)

func main() {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)

	window, err := sdl.CreateWindow("Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(width), int32(height), sdl.WINDOW_OPENGL|sdl.WINDOW_RESIZABLE)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()
	window.GLCreateContext()
	window.SetResizable(true)
	window.SetMinimumSize(800, 600)

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
	// First way to interpret camera
	// cameraPosition := mgl.Vec3{0, 0, -10}
	// cameraTarget := mgl.Vec3{0, 0, 0}
	// cameraDirection := cameraPosition.Sub(cameraTarget).Normalize()
	// up := mgl.Vec3{0, 1, 0}
	// cameraRight := up.Cross(cameraDirection).Normalize()
	// cameraUp := cameraDirection.Cross(cameraRight)
	// view := mgl.LookAtV(cameraPosition, cameraTarget, cameraUp)
	cameraPosition = mgl.Vec3{1, 1, 10}
	cameraFront := mgl.Vec3{-3, 3, 0}
	cameraUp := mgl.Vec3{0, 1, 0}
	var cameraSpeed float32 = 0.25
	view = mgl.LookAtV(cameraPosition, cameraPosition.Add(cameraFront), cameraUp)
	// view := mgl.Ident4().Mul4(mgl.Translate3D(0, 0, -5))

	model1 := mgl.Ident4().Mul4(mgl.Translate3D(-5, -2, 1)).Mul4(mgl.Scale3D(0.5, 0.5, 0.5)).Mul4(mgl.HomogRotate3DX(mgl.DegToRad(180)))
	positionArr := []mgl.Vec3{
		{0, 0, 0},
		{2, 5, -15},
		{-2.4, 0.4, 3.5},
		{-1.3, 1, -3},
	}
	keys := make(map[sdl.Scancode]bool, 1024)

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
				if t.Event == sdl.WINDOWEVENT_RESIZED || t.Event == sdl.WINDOWEVENT_SIZE_CHANGED {
					w, h := window.GetSize()
					width, height = float32(w), float32(h)
					projection = mgl.Perspective(mgl.DegToRad(60), width/height, 0.1, 100)
					gl.Viewport(0, 0, w, h)
				}
			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN {
					keys[t.Keysym.Scancode] = true
					// if t.Keysym.Scancode == sdl.SCANCODE_W {
					// 	cameraPosition = cameraPosition.Add(cameraFront.Mul(cameraSpeed))
					// 	view = mgl.LookAtV(cameraPosition, cameraPosition.Add(cameraFront), cameraUp)
					// }
					// if t.Keysym.Scancode == sdl.SCANCODE_S {
					// 	cameraPosition = cameraPosition.Sub(cameraFront.Mul(cameraSpeed))
					// 	view = mgl.LookAtV(cameraPosition, cameraPosition.Add(cameraFront), cameraUp)
					// }
					// if t.Keysym.Scancode == sdl.SCANCODE_A {
					// 	cameraPosition = cameraPosition.Sub(cameraFront.Cross(cameraUp).Mul(cameraSpeed))
					// 	view = mgl.LookAtV(cameraPosition, cameraPosition.Add(cameraFront), cameraUp)
					// }
					// if t.Keysym.Scancode == sdl.SCANCODE_D {
					// 	cameraPosition = cameraPosition.Add(cameraFront.Cross(cameraUp).Mul(cameraSpeed))
					// 	view = mgl.LookAtV(cameraPosition, cameraPosition.Add(cameraFront), cameraUp)
					// }
					// if t.Keysym.Scancode == sdl.SCANCODE_LSHIFT || t.Keysym.Scancode == sdl.SCANCODE_C {
					// 	view = view.Mul4(mgl.Translate3D(0, 1, 0))
					// }
					// if t.Keysym.Scancode == sdl.SCANCODE_SPACE {
					// 	view = view.Mul4(mgl.Translate3D(0, -1, 0))
					// }
					// if t.Keysym.Scancode == sdl.SCANCODE_F {
					// 	window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
					// }
					// if t.Keysym.Scancode == sdl.SCANCODE_ESCAPE {
					// 	window.SetFullscreen(0)
					// }
				}
				if t.Type == sdl.KEYUP {
					keys[t.Keysym.Scancode] = false
				}
			}
		}
		// Timer
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		updateTimer := sdl.GetTicks()
		if updateTimer-timer > 100 {
			// model = model.Mul4(mgl.HomogRotate3DY(mgl.DegToRad(1)))
			// model1 = model1.Mul4(mgl.HomogRotate3DY(mgl.DegToRad(2)))
			timer = updateTimer
		}

		// Upload data to shader
		sprite.Shader.SetMat4("projection", projection)
		sprite.Shader.SetMat4("view", view)
		for i, v := range positionArr {
			model := mgl.Ident4()
			model = model.Mul4(mgl.Translate3D(v.X(), v.Y(), v.Z()))
			model = model.Mul4(mgl.HomogRotate3D(mgl.DegToRad(30*float32(i)), mgl.Vec3{1, 0.3, 0.5}))
			sprite.Shader.SetMat4("model", model)
			// Draw sprite
			sprite.Render()
		}

		sprite.Shader.SetMat4("model", model1)
		sprite.Render()

		window.GLSwap()
	}
}
