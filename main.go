package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {

	err := glfw.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	window, err := glfw.CreateWindow(500, 400, "RPG", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	window.MakeContextCurrent()
	defer window.Destroy()

	err = gl.Init()
	if err != nil {
		log.Fatal(err)
	}
	fragmentShaderSource :=
		`#version 460 core
		out vec4 FragColor;

		void main()
		{
			FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
		}`
	vertexShaderSource :=
		`#version 460 core
		layout (location = 0) in vec3 aPos;

		void main()
		{
			gl_Position = vec4(aPos, 1.0);
		}`

	shaderProgram := loadShaderProgram(fragmentShaderSource, vertexShaderSource)

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0}

	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.ClearColor(0, 0, 0, 1)
	log.Print("OpenGL version: ", gl.GoStr(gl.GetString(gl.VERSION)))
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(shaderProgram)
		gl.BindVertexArray(VAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func loadShaderProgram(fragmentShaderSource, vertexShaderSource string) (shaderProgram uint32){
	// Load shaders
	fragmentShader := setShaderSource(fragmentShaderSource, gl.FRAGMENT_SHADER)
	vertexShader := setShaderSource(vertexShaderSource, gl.VERTEX_SHADER)
	var status int32
	gl.CompileShader(fragmentShader)
	gl.GetShaderiv(fragmentShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var infoLog uint8
		gl.GetShaderInfoLog(fragmentShader, gl.INFO_LOG_LENGTH, nil, &infoLog)
		log.Println(("Error: fragmentShader compilation failed: "))
		log.Fatal(gl.GoStr(&infoLog))
	}
	gl.CompileShader(vertexShader)
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var infoLog uint8
		gl.GetShaderInfoLog(vertexShader, gl.INFO_LOG_LENGTH, nil, &infoLog)
		log.Println(("Error: vertexShader compilation failed: "))
		log.Fatal(gl.GoStr(&infoLog))
	}
	// Linking
	shaderProgram = gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)
	gl.GetShaderiv(shaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var infoLog uint8
		gl.GetProgramInfoLog(shaderProgram, 512, nil, &infoLog)
		log.Println(("Error: shaderProgram linking failed: "))
		log.Fatal(gl.GoStr(&infoLog))
	}
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	return shaderProgram
}

func setShaderSource(source string, shaderType uint32) uint32 {
	shader := gl.CreateShader(shaderType)
	cSource, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, cSource, nil)
	free()
	var status int32
	gl.CompileShader(shader)
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var infoLog uint8
		gl.GetShaderInfoLog(shader, gl.INFO_LOG_LENGTH, nil, &infoLog)
		if shaderType == gl.VERTEX_SHADER {
			log.Println(("Error: vertexShader compilation failed: "))
		} else {
			log.Println(("Error: fragmentShader compilation failed: "))
		}
		log.Fatal(gl.GoStr(&infoLog))
	}
	return shader
}
