package glfw

import (
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"golang.org/x/sys/windows"
	"os"
	"path/filepath"
	"runtime"
)

//go:embed glfw.dll
var dllData []byte

func init() {
	runtime.LockOSThread()
	runtime.GOMAXPROCS(runtime.NumCPU())
	dir := mylog.Check2(os.UserCacheDir())
	dir = filepath.Join(dir, "glfw3", "dll_cache")
	stream.CreatDirectory(dir)
	mylog.Check(windows.SetDllDirectory(dir))
	sha := sha256.Sum256(dllData)
	dllName := fmt.Sprintf("glfw3-%s.dll", base64.RawURLEncoding.EncodeToString(sha[:]))
	filePath := filepath.Join(dir, dllName)
	if !stream.IsFilePath(filePath) {
		stream.WriteTruncate(filePath, dllData)
	}
	mylog.Check2(GengoLibrary.LoadFrom(filePath))
}
