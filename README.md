# GLFW Go Bindings

基于 [GLFW 3.5](https://www.glfw.org/) 的 Go 语言绑定，**纯 Go 实现，无 CGO 依赖**。

## 特性

- **🚀 无 CGO**: 纯 Go 实现，使用 `golang.org/x/sys/windows` 动态加载 DLL，编译速度快，交叉编译简单
- **自动生成**: 从 GLFW 官方头文件自动生成所有绑定代码
- **DLL 嵌入**: `glfw3.dll` 嵌入到 Go 二进制文件中，运行时自动解压到用户缓存目录，无需手动分发 DLL
- **完整覆盖**: 包含所有 GLFW 函数、结构体、类型定义和常量

## 生成内容

| 类型 | 数量 | 说明 |
|------|------|------|
| 函数 | 120 | 所有 `glfw*` API 函数 |
| 结构体 | 8 | `GLFWmonitor`, `GLFWwindow`, `GLFWcursor`, `GLFWvidmode`, `GLFWgammaramp`, `GLFWimage`, `GLFWgamepadstate`, `GLFWallocator` |
| 类型定义 | 25 | 回调函数类型 (`GLFWerrorfun`, `GLFWkeyfun` 等) |
| 常量 | 300+ | 键码、鼠标按钮、窗口提示、错误码等 |

## 使用方法

```go
package main

import (
    "fmt"
    "github.com/ddkwork/glfw"
)

func main() {
    g := &glfw.Glfw{}
    
    // 初始化 GLFW
    if g.Init() == glfw.GlfwFalse {
        panic("Failed to initialize GLFW")
    }
    defer g.Terminate()
    
    // 获取版本
    var major, minor, revision int32
    g.GetVersion(&major, &minor, &revision)
    fmt.Printf("GLFW version: %d.%d.%d\n", major, minor, revision)
    
    // 创建窗口
    g.WindowHint(glfw.GlfwContextVersionMajor, 3)
    g.WindowHint(glfw.GlfwContextVersionMinor, 3)
    g.WindowHint(glfw.GlfwOpenglProfile, glfw.GlfwOpenglCoreProfile)
    
    window := g.CreateWindow(800, 600, "GLFW Window", nil, nil)
    if window == nil {
        panic("Failed to create window")
    }
    defer g.DestroyWindow(window)
    
    // 主循环
    for g.WindowShouldClose(window) == glfw.GlfwFalse {
        g.SwapBuffers(window)
        g.PollEvents()
    }
}
```

## 回调函数

GLFW 回调通过函数指针类型实现：

```go
// 键盘回调
g.SetKeyCallback(window, func(w *glfw.GLFWwindow, key, scancode, action, mods int32) uintptr {
    if key == glfw.GlfwKeyEscape && action == glfw.GlfwPress {
        g.SetWindowShouldClose(w, glfw.GlfwTrue)
    }
    return 0
})

// 窗口大小回调
g.SetWindowSizeCallback(window, func(w *glfw.GLFWwindow, width, height int32) uintptr {
    fmt.Printf("Window resized: %dx%d\n", width, height)
    return 0
})
```

## 常用常量

```go
// 布尔值
glfw.GlfwTrue   // 1
glfw.GlfwFalse  // 0

// 输入状态
glfw.GlfwRelease  // 0
glfw.GlfwPress    // 1
glfw.GlfwRepeat   // 2

// 键盘按键
glfw.GlfwKeyEscape, glfw.GlfwKeySpace, glfw.GlfwKeyA, ...

// 鼠标按钮
glfw.GlfwMouseButtonLeft, glfw.GlfwMouseButtonRight, glfw.GlfwMouseButtonMiddle

// 窗口提示
glfw.GlfwVisible, glfw.GlfwResizable, glfw.GlfwContextVersionMajor, ...

// 平台
glfw.GlfwPlatformWin32, glfw.GlfwPlatformCocoa, glfw.GlfwPlatformX11, glfw.GlfwPlatformWayland
```

## 目录结构

```
glfw/
├── glfw3.go          # 自动生成的绑定代码 (1097 行)
├── dll.go            # DLL 加载和嵌入逻辑
├── glfw3.dll         # GLFW 动态库 (嵌入到 dll.go)
├── dll_test.go       # 测试用例
├── generate_test.go  # 绑定生成配置
├── clone/            # GLFW 源码和头文件
│   └── glfw/
│       └── include/GLFW/glfw3.h
└── build/            # 编译输出目录
```

## 生成绑定

修改 `generate_test.go` 配置后运行：

```bash
go test -run TestGenerate -v
```

配置说明：

```go
c2go.BindgenConfig{
    HeadersDir:     "glfw/include",      // GLFW 头文件目录
    OutputDir:      "../",               // 输出目录
    PackageName:    "glfw",              // Go 包名
    RecurseHeaders: true,                // 递归处理头文件
    HeaderOrder:    []string{"GLFW/glfw3.h"}, // 处理顺序
    BindDll:        true,                // 生成 DLL 绑定
    DllName:        "glfw3.dll",         // DLL 名称
    Predefined: `
#define GLFW_INCLUDE_NONE  // 不自动包含 OpenGL 头文件
#define GLFW_DLL           // 使用 DLL 链接
`,
    DllFuncFilter: func(name string) bool {
        return strings.HasPrefix(name, "glfw") // 只导出 glfw* 函数
    },
}
```

## 构建 GLFW DLL

使用 CMake 构建 GLFW 动态库：

```bash
cd clone
cmake -S . -B build -DBUILD_SHARED_LIBS=ON -DGLFW_BUILD_EXAMPLES=OFF -DGLFW_BUILD_TESTS=OFF
cmake --build build --config Release
cp build/src/Release/glfw3.dll ../
```

## 测试

```bash
go test -v
```

测试覆盖：
- GLFW 初始化和终止
- 版本查询
- 平台检测
- 错误回调
- 显示器枚举
- 窗口创建和操作

## 版本信息

- GLFW: 3.5.0
- Go 1.26+
- 平台: Windows (Win32 WGL)

## 许可证

- Go 绑定代码: MIT
- GLFW: zlib/libpng license (见 `clone/glfw/LICENSE.md`)
