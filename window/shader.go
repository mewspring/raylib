package window

// #include <stdlib.h>
// #include <raylib.h>
import "C"

import (
	"runtime"
	"unsafe"
)

// Shader holds a vertex or fragment shader.
type Shader struct {
	// underlying raylib shader.
	_shader C.Shader
}

// LoadShader loads the given vertex and fragment shader.
func LoadShader(vsPath, fsPath string) (*Shader, error) {
	_vsPath := C.CString(vsPath)
	defer C.free(unsafe.Pointer(_vsPath))
	_fsPath := C.CString(fsPath)
	defer C.free(unsafe.Pointer(_fsPath))
	_shader := C.LoadShader(_vsPath, _fsPath)
	// TODO: figure out how to check error.
	shader := &Shader{
		_shader: _shader,
	}
	// Set finalizer to free shader.
	free := func(obj any) {
		C.UnloadShader(_shader)
	}
	runtime.SetFinalizer(shader, free)
	return shader, nil
}

// Enable enables drawing of the shader.
func (shader *Shader) Enable() {
	C.BeginShaderMode(shader._shader)
}

// Disable disables drawing of the shader.
func (shader *Shader) Disable() {
	C.EndShaderMode()
}
