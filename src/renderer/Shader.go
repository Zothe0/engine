package renderer

import (
	"io/ioutil"
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
)

// Shader ...
type Shader struct {
	ID             uint32
	vertexSource   string
	fragmentSource string
}

// NewShader constructor for shader
func NewShader(vertexPath, fragmentPath string) (s *Shader) {
	s = new(Shader)
	s.vertexSource = readShader(vertexPath)
	s.fragmentSource = readShader(fragmentPath)
	s.loadShaderProgram()
	return s
}

// Compile shaders needs in one scope with the shader program linking
func (s *Shader) loadShaderProgram() {
	// Compile shaders
	vertexShader := setShaderSource(s.vertexSource, gl.VERTEX_SHADER)
	fragmentShader := setShaderSource(s.fragmentSource, gl.FRAGMENT_SHADER)
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
	s.ID = gl.CreateProgram()
	gl.AttachShader(s.ID, vertexShader)
	gl.AttachShader(s.ID, fragmentShader)
	gl.LinkProgram(s.ID)
	gl.GetShaderiv(s.ID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var infoLog uint8
		gl.GetProgramInfoLog(s.ID, 512, nil, &infoLog)
		log.Println(("Error: shaderProgram linking failed: "))
		log.Fatal(gl.GoStr(&infoLog))
	}
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
}

// Use shader program
func (s *Shader) Use() {
	gl.UseProgram(s.ID)
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
