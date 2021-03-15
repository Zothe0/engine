package game

import (
	"engine/src/manager"
	"fmt"
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
	"github.com/veandco/go-sdl2/sdl"
)

type state int

const (
	Run state = iota
	Stop
)

type Game struct {
	gameState      state
	width          int32
	height         int32
	cameraSpeed    float32
	KeyMap         map[sdl.Scancode]bool
	pWindow        *sdl.Window
	rm             *manager.ResourceManager
	cameraPosition mgl.Vec3
	viewMat        mgl.Mat4
	projectionMat  mgl.Mat4
}

func InitGame() {
	game := new(Game)
	// Init simple variables
	game.gameState = Stop
	game.width = 800
	game.height = 600
	game.cameraSpeed = 0.25
	game.KeyMap = make(map[sdl.Scancode]bool, 512)

	game.init()
	game.mainloop()
	game.cleanup()
}
func (g *Game) init() {
	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 6)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)

	g.pWindow, err = sdl.CreateWindow("Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		g.width, g.height, sdl.WINDOW_OPENGL|sdl.WINDOW_RESIZABLE)
	if err != nil {
		log.Fatal(err)
	}
	g.pWindow.GLCreateContext()
	g.pWindow.SetMinimumSize(g.width, g.height)

	err = gl.Init()
	if err != nil {
		log.Fatal("Gl init error: ", err)
	}
	g.rm = manager.InitResourceManager()
	subTextures := []string{"violet", "blue", "brown"}
	g.rm.AddShader("default", "./res/shaders/defaultVertex.glsl", "./res/shaders/defaultFragment.glsl")
	g.rm.AddShader("color", "./res/shaders/colorV.glsl", "./res/shaders/colorF.glsl")
	g.rm.AddTexture("spirits", "./res/textures/spirits.png", &subTextures, 170, 220)
	g.rm.AddSprite("unit", "color", "", "")

	fmt.Println("OpenGL version:", gl.GoStr(gl.GetString(gl.VERSION)))

	g.projectionMat = mgl.Perspective(mgl.DegToRad(60), float32(g.width/g.height), 0.1, 100)
	g.cameraPosition = mgl.Vec3{1, 1, 10}
	// cameraFront := mgl.Vec3{0, 1, 0}
	// cameraUp := mgl.Vec3{0, 1, 0}
	// g.viewMat = mgl.LookAtV(g.cameraPosition, g.cameraPosition.Add(cameraFront), cameraUp)
	g.viewMat = mgl.Ident4().Mul4(mgl.Translate3D(0, 0, -5))

	gl.ClearColor(0, 0, 0.3, 1)
	gl.Enable(gl.DEPTH_TEST) // For 3D correct drawn
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
}
func (g *Game) mainloop() {
	positionArr := []mgl.Vec3{
		{0, 0, 0},
		{2, 5, -15},
		{-2.4, 0.4, 3.5},
		{-1.3, 1, -3},
	}
	g.gameState = Run
	timer := sdl.GetTicks()
	run := true
	curSprite := g.rm.GetSprite("unit")
	for run {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				run = false
			case *sdl.WindowEvent:
				if t.Event == sdl.WINDOWEVENT_RESIZED || t.Event == sdl.WINDOWEVENT_SIZE_CHANGED || t.Event == sdl.WINDOWEVENT_MAXIMIZED {
					g.width, g.height = g.pWindow.GetSize()
					g.projectionMat = mgl.Perspective(mgl.DegToRad(60), float32(g.width/g.height), 0.1, 100)
					gl.Viewport(0, 0, g.width, g.height)
				}
			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN {
					g.KeyMap[t.Keysym.Scancode] = true
				}
				if t.Type == sdl.KEYUP {
					g.KeyMap[t.Keysym.Scancode] = false
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
		curSprite.Shader.SetMat4("projection", g.projectionMat)
		curSprite.Shader.SetMat4("view", g.viewMat)
		for i, v := range positionArr {
			model := mgl.Ident4()
			model = model.Mul4(mgl.Translate3D(v.X(), v.Y(), v.Z()))
			model = model.Mul4(mgl.HomogRotate3D(mgl.DegToRad(30*float32(i)), mgl.Vec3{1, 0.3, 0.5}))
			curSprite.Shader.SetMat4("model", model)
			// Draw sprite
			curSprite.Render()
		}

		g.pWindow.GLSwap()
	}
}
func (g *Game) cleanup() {
	sdl.Quit()
	g.pWindow.Destroy()
}
