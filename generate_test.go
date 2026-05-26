package glfw

import (
	"strings"
	"testing"

	"github.com/ddkwork/bindgen/c2go"
)

func TestGenerate(t *testing.T) {
	c2go.Generate(t, []c2go.BindgenConfig{{
		HeadersDir:     "clone/glfw/include",
		OutputDir:      ".",
		PackageName:    "glfw",
		RecurseHeaders: true,
		HeaderOrder:    []string{"GLFW/glfw3.h"},
		BindDll:        true,
		DllName:        "glfw3.dll",
		Predefined: `
#define GLFW_INCLUDE_NONE
#define GLFW_DLL
`,
		DllFuncFilter: func(name string) bool {
			return strings.HasPrefix(name, "glfw")
		},
	}})
}
