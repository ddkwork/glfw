package main

import (
	"fmt"
	"github.com/ddkwork/glfw"
	"github.com/ebitengine/purego"
	"image"
	"image/color"
	"unsafe"
)

func main() {
	glfw.Init()
	defer glfw.Terminate()
	w := glfw.CreateWindow(640, 480, StringToBytePointer("Custom Cursor"), nil, nil)
	glfw.MakeContextCurrent(w)

	// Creating a custom cursor.
	callback := purego.NewCallback(func(img *image.NRGBA) *image.NRGBA {
		c := color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
		const size = 16
		m := image.NewNRGBA(image.Rect(0, 0, size, size))
		for y := 0; y < size; y++ {
			for x := 0; x < size-y; x++ {
				m.SetNRGBA(x, y, c)
			}
		}
		return m
	})
	cursor := glfw.CreateCursor(unsafe.Pointer(callback), 0, 0)
	glfw.SetCursor(w, cursor)

	newCallback := purego.NewCallback(func(window uintptr, entered int) uintptr {
		fmt.Println("Cursor Enter")
		return uintptr(w)
	})
	glfw.SetCursorEnterCallback(w, glfw.Cursorenterfun(newCallback))

	//typedef void (* GLFWdropfun)(GLFWwindow* window, int path_count, const char* paths[]);
	//c := purego.NewCallback(func(win uintptr, path_count int, files **byte) uintptr {
	//	paths := make([]string, path_count)
	//	for i := 0; i < path_count; i++ {
	//		paths[i] = BytePointerToString((files)) //todo not working
	//	}
	//	return uintptr(w)
	//})
	//glfw.SetDropCallback(w, glfw.Dropfun(c))

	//typedef void (* GLFWdropfun)(GLFWwindow* window, int path_count, const char* paths[]);
	c := purego.NewCallback(func(win uintptr, path_count int, paths **byte) uintptr {
		convertedPaths := make([]string, path_count)
		for i := 0; i < path_count; i++ {
			convertedPaths[i] = BytePointerToString(*(**byte)(unsafe.Pointer(uintptr(unsafe.Pointer(paths)) + uintptr(i)*unsafe.Sizeof(*paths))))
		}
		fmt.Println(convertedPaths) // 打印路径以验证
		return uintptr(w)
	})
	glfw.SetDropCallback(w, glfw.Dropfun(c))

	glfw.SetClipboardString(w, StringToBytePointer("Hello, World!"))
	glfw.SetCharCallback(w, glfw.Charfun(purego.NewCallback(func(s *byte) uintptr {
		println(BytePointerToString(s))
		return uintptr(w)
	})))

	glfw.SetKeyCallback(w, glfw.Keyfun(purego.NewCallback(func() uintptr {
		fmt.Println("Key")
		return uintptr(w)
	})))

	// Setting a custom cursor.
	//w.SetIcon([]image.Image{whiteTriangle})
	//glfw.SetWindowIcon(w, 1, unsafe.Pointer(whiteTriangle))

	for !Boolean2Bool(glfw.WindowShouldClose(w)) {
		glfw.SwapBuffers(w)
		glfw.PollEvents()
	}

	//for {
	//	// PostEmptyEvent()
	//	glfw.PollEvents()
	//	glfw.SwapBuffers(w)
	//	if glfw.WindowShouldClose(w) != 0 {
	//		glfw.DestroyWindow(w)
	//		break
	//	}
	//}

	// WindowHint(Visible, False)
	// WindowHint(Resizable,Enable(!Resizable))
	// WindowHint(Decorated,Enable(!w.Decorated))
	// WindowHint(Floating,Enable(Floating))
	// WindowHint(AutoIconify, False)
	// WindowHint(TransparentFramebuffer, False)
	// WindowHint(FocusOnShow, False)
	// WindowHint(ScaleToMonitor, False)

	// PostEmptyEvent()

	// SetCursorEnterCallback(w, func() {})
	// SetCursorPosCallback(w, func() {})
	// SetMouseButtonCallback(w, func() {})
	// SetWindowFocusCallback(w, func() {})
	// SetWindowCloseCallback(w, func() {})
	// SetWindowSizeCallback(w, func() {})
	// SetWindowRefreshCallback(w, func() {})
	// SetScrollCallback(w, func() {})
	// SetKeyCallback(w, func() {})
	// SetCharCallback(w, func() {})
	// SetDropCallback(w, func() {})
	// SetWindowIcon(w,32, func() {})

}
func Boolean2Bool(b glfw.Bool) bool {
	return int(b) == glfw.True
}

var whiteTriangle = func() *image.NRGBA {
	c := color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
	const size = 16
	m := image.NewNRGBA(image.Rect(0, 0, size, size))
	for y := 0; y < size; y++ {
		for x := 0; x < size-y; x++ {
			m.SetNRGBA(x, y, c)
		}
	}
	return m
}()

func StringToBytePointer(s string) *byte {
	bytes := []byte(s)
	ptr := &bytes[0]
	return ptr
}

func BytePointerToString(ptr *byte) string {
	var bytes []byte
	for *ptr != 0 {
		bytes = append(bytes, *ptr)
		ptr = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + 1))
	}
	return string(bytes)
}
