package utils

import (
	"reflect"

	"github.com/go-gl/gl/v2.1/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
)

// Sizeof computing byte size of structure
func Sizeof(T interface{}) int {
	return int(reflect.TypeOf(T).Size())
}
func Cstr(str string) *uint8 {
	return gl.Str(str + "\x00")
}
func MatAddress(matrix mgl.Mat4) *float32 {
	matrix = [16]float32(matrix)
	return &matrix[0]
}
