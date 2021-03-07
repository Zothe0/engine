package renderer

import (
	"io/ioutil"
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
)

// Shader ...
type Shader struct {
	vertexSource   string
	fragmentSource string
	shaderProgram  uint32
}

// LoadShader constructor for shader
func LoadShader(vertexPath, fragmentPath string) *Shader {
	var shader Shader
	shader.vertexSource = readShader(vertexPath)
	shader.fragmentSource = readShader(fragmentPath)
	shader.loadShaderProgram()
	return &shader
}

// Compile shaders needs in one scope with the shader program linking
func (s *Shader) loadShaderProgram() {
	// Compile shaders
	fragmentShader := setShaderSource(s.fragmentSource, gl.FRAGMENT_SHADER)
	vertexShader := setShaderSource(s.vertexSource, gl.VERTEX_SHADER)
	var status int32
	gl.CompileShader(vertexShader)
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var infoLog uint8
		gl.GetShaderInfoLog(vertexShader, gl.INFO_LOG_LENGTH, nil, &infoLog)
		log.Println(("Error: vertexShader compilation failed: "))
		log.Fatal(gl.GoStr(&infoLog))
	}
	gl.CompileShader(fragmentShader)
	gl.GetShaderiv(fragmentShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var infoLog uint8
		gl.GetShaderInfoLog(fragmentShader, gl.INFO_LOG_LENGTH, nil, &infoLog)
		log.Println(("Error: fragmentShader compilation failed: "))
		log.Fatal(gl.GoStr(&infoLog))
	}
	// Linking
	s.shaderProgram = gl.CreateProgram()
	gl.AttachShader(s.shaderProgram, vertexShader)
	gl.AttachShader(s.shaderProgram, fragmentShader)
	gl.LinkProgram(s.shaderProgram)
	gl.GetShaderiv(s.shaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var infoLog uint8
		gl.GetProgramInfoLog(s.shaderProgram, 512, nil, &infoLog)
		log.Println(("Error: shaderProgram linking failed: "))
		log.Fatal(gl.GoStr(&infoLog))
	}
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
}

// Use shader program
func (s *Shader) Use() {
	gl.UseProgram(s.shaderProgram)
}

func readShader(path string) string {
	source, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(source)
}

func setShaderSource(source string, shaderType uint32) uint32 {
	shader := gl.CreateShader(shaderType)
	cSource, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, cSource, nil)
	free()
	return shader
}
