package glfw

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type GLFWglproc func() uintptr

type GLFWvkproc func() uintptr

type GLFWallocatefun func(uintptr, unsafe.Pointer) uintptr

type GLFWreallocatefun func(unsafe.Pointer, uintptr, unsafe.Pointer) uintptr

type GLFWdeallocatefun func(unsafe.Pointer, unsafe.Pointer) uintptr

type GLFWerrorfun func(int32, *int8) uintptr

type GLFWwindowposfun func(*GLFWwindow, int32, int32) uintptr

type GLFWwindowsizefun func(*GLFWwindow, int32, int32) uintptr

type GLFWwindowclosefun func(*GLFWwindow) uintptr

type GLFWwindowrefreshfun func(*GLFWwindow) uintptr

type GLFWwindowfocusfun func(*GLFWwindow, int32) uintptr

type GLFWwindowiconifyfun func(*GLFWwindow, int32) uintptr

type GLFWwindowmaximizefun func(*GLFWwindow, int32) uintptr

type GLFWframebuffersizefun func(*GLFWwindow, int32, int32) uintptr

type GLFWwindowcontentscalefun func(*GLFWwindow, float32, float32) uintptr

type GLFWmousebuttonfun func(*GLFWwindow, int32, int32, int32) uintptr

type GLFWcursorposfun func(*GLFWwindow, float64, float64) uintptr

type GLFWcursorenterfun func(*GLFWwindow, int32) uintptr

type GLFWscrollfun func(*GLFWwindow, float64, float64) uintptr

type GLFWkeyfun func(*GLFWwindow, int32, int32, int32, int32) uintptr

type GLFWcharfun func(*GLFWwindow, uint32) uintptr

type GLFWcharmodsfun func(*GLFWwindow, uint32, int32) uintptr

type GLFWdropfun func(*GLFWwindow, int32, **int8) uintptr

type GLFWmonitorfun func(*GLFWmonitor, int32) uintptr

type GLFWjoystickfun func(int32, int32) uintptr

type (
	GLFWmonitor struct{} // glfw3.h:1396 -> GLFWmonitor
	GLFWwindow  struct{} // glfw3.h:1408 -> GLFWwindow
	GLFWcursor  struct{} // glfw3.h:1420 -> GLFWcursor
	GLFWvidmode struct {
		Width       int32
		Height      int32
		RedBits     int32
		GreenBits   int32
		BlueBits    int32
		RefreshRate int32
	} // glfw3.h:2031 -> GLFWvidmode
	GLFWgammaramp struct {
		Red   *uint16
		Green *uint16
		Blue  *uint16
		Size  uint32
		_     [4]byte
	} // glfw3.h:2065 -> GLFWgammaramp
	GLFWimage struct {
		Width  int32
		Height int32
		Pixels *uint8
	} // glfw3.h:2094 -> GLFWimage
	GLFWgamepadstate struct {
		Buttons [15]uint8
		_       [1]byte
		Axes    [6]float32
	} // glfw3.h:2118 -> GLFWgamepadstate
	GLFWallocator struct {
		Allocate   GLFWallocatefun
		Reallocate GLFWreallocatefun
		Deallocate GLFWdeallocatefun
		User       unsafe.Pointer
	} // glfw3.h:2142 -> GLFWallocator
)

// Source: glfw3.h -> Macro constants
const (
	GlfwVersionMajor              uint32 = 3
	GlfwVersionMinor              uint32 = 5
	GlfwVersionRevision           uint32 = 0
	GlfwTrue                      uint32 = 1
	GlfwFalse                     uint32 = 0
	GlfwRelease                   uint32 = 0
	GlfwPress                     uint32 = 1
	GlfwRepeat                    uint32 = 2
	GlfwHatCentered               uint32 = 0
	GlfwHatUp                     uint32 = 1
	GlfwHatRight                  uint32 = 2
	GlfwHatDown                   uint32 = 4
	GlfwHatLeft                   uint32 = 8
	GlfwHatRightUp                uint32 = (GlfwHatRight | GlfwHatUp)
	GlfwHatRightDown              uint32 = (GlfwHatRight | GlfwHatDown)
	GlfwHatLeftUp                 uint32 = (GlfwHatLeft | GlfwHatUp)
	GlfwHatLeftDown               uint32 = (GlfwHatLeft | GlfwHatDown)
	GlfwKeyUnknown                int32  = -1
	GlfwKeySpace                  uint32 = 32
	GlfwKeyApostrophe             uint32 = 39
	GlfwKeyComma                  uint32 = 44
	GlfwKeyMinus                  uint32 = 45
	GlfwKeyPeriod                 uint32 = 46
	GlfwKeySlash                  uint32 = 47
	GlfwKey0                      uint32 = 48
	GlfwKey1                      uint32 = 49
	GlfwKey2                      uint32 = 50
	GlfwKey3                      uint32 = 51
	GlfwKey4                      uint32 = 52
	GlfwKey5                      uint32 = 53
	GlfwKey6                      uint32 = 54
	GlfwKey7                      uint32 = 55
	GlfwKey8                      uint32 = 56
	GlfwKey9                      uint32 = 57
	GlfwKeySemicolon              uint32 = 59
	GlfwKeyEqual                  uint32 = 61
	GlfwKeyA                      uint32 = 65
	GlfwKeyB                      uint32 = 66
	GlfwKeyC                      uint32 = 67
	GlfwKeyD                      uint32 = 68
	GlfwKeyE                      uint32 = 69
	GlfwKeyF                      uint32 = 70
	GlfwKeyG                      uint32 = 71
	GlfwKeyH                      uint32 = 72
	GlfwKeyI                      uint32 = 73
	GlfwKeyJ                      uint32 = 74
	GlfwKeyK                      uint32 = 75
	GlfwKeyL                      uint32 = 76
	GlfwKeyM                      uint32 = 77
	GlfwKeyN                      uint32 = 78
	GlfwKeyO                      uint32 = 79
	GlfwKeyP                      uint32 = 80
	GlfwKeyQ                      uint32 = 81
	GlfwKeyR                      uint32 = 82
	GlfwKeyS                      uint32 = 83
	GlfwKeyT                      uint32 = 84
	GlfwKeyU                      uint32 = 85
	GlfwKeyV                      uint32 = 86
	GlfwKeyW                      uint32 = 87
	GlfwKeyX                      uint32 = 88
	GlfwKeyY                      uint32 = 89
	GlfwKeyZ                      uint32 = 90
	GlfwKeyLeftBracket            uint32 = 91
	GlfwKeyBackslash              uint32 = 92
	GlfwKeyRightBracket           uint32 = 93
	GlfwKeyGraveAccent            uint32 = 96
	GlfwKeyWorld1                 uint32 = 161
	GlfwKeyWorld2                 uint32 = 162
	GlfwKeyEscape                 uint32 = 256
	GlfwKeyEnter                  uint32 = 257
	GlfwKeyTab                    uint32 = 258
	GlfwKeyBackspace              uint32 = 259
	GlfwKeyInsert                 uint32 = 260
	GlfwKeyDelete                 uint32 = 261
	GlfwKeyRight                  uint32 = 262
	GlfwKeyLeft                   uint32 = 263
	GlfwKeyDown                   uint32 = 264
	GlfwKeyUp                     uint32 = 265
	GlfwKeyPageUp                 uint32 = 266
	GlfwKeyPageDown               uint32 = 267
	GlfwKeyHome                   uint32 = 268
	GlfwKeyEnd                    uint32 = 269
	GlfwKeyCapsLock               uint32 = 280
	GlfwKeyScrollLock             uint32 = 281
	GlfwKeyNumLock                uint32 = 282
	GlfwKeyPrintScreen            uint32 = 283
	GlfwKeyPause                  uint32 = 284
	GlfwKeyF1                     uint32 = 290
	GlfwKeyF2                     uint32 = 291
	GlfwKeyF3                     uint32 = 292
	GlfwKeyF4                     uint32 = 293
	GlfwKeyF5                     uint32 = 294
	GlfwKeyF6                     uint32 = 295
	GlfwKeyF7                     uint32 = 296
	GlfwKeyF8                     uint32 = 297
	GlfwKeyF9                     uint32 = 298
	GlfwKeyF10                    uint32 = 299
	GlfwKeyF11                    uint32 = 300
	GlfwKeyF12                    uint32 = 301
	GlfwKeyF13                    uint32 = 302
	GlfwKeyF14                    uint32 = 303
	GlfwKeyF15                    uint32 = 304
	GlfwKeyF16                    uint32 = 305
	GlfwKeyF17                    uint32 = 306
	GlfwKeyF18                    uint32 = 307
	GlfwKeyF19                    uint32 = 308
	GlfwKeyF20                    uint32 = 309
	GlfwKeyF21                    uint32 = 310
	GlfwKeyF22                    uint32 = 311
	GlfwKeyF23                    uint32 = 312
	GlfwKeyF24                    uint32 = 313
	GlfwKeyF25                    uint32 = 314
	GlfwKeyKp0                    uint32 = 320
	GlfwKeyKp1                    uint32 = 321
	GlfwKeyKp2                    uint32 = 322
	GlfwKeyKp3                    uint32 = 323
	GlfwKeyKp4                    uint32 = 324
	GlfwKeyKp5                    uint32 = 325
	GlfwKeyKp6                    uint32 = 326
	GlfwKeyKp7                    uint32 = 327
	GlfwKeyKp8                    uint32 = 328
	GlfwKeyKp9                    uint32 = 329
	GlfwKeyKpDecimal              uint32 = 330
	GlfwKeyKpDivide               uint32 = 331
	GlfwKeyKpMultiply             uint32 = 332
	GlfwKeyKpSubtract             uint32 = 333
	GlfwKeyKpAdd                  uint32 = 334
	GlfwKeyKpEnter                uint32 = 335
	GlfwKeyKpEqual                uint32 = 336
	GlfwKeyLeftShift              uint32 = 340
	GlfwKeyLeftControl            uint32 = 341
	GlfwKeyLeftAlt                uint32 = 342
	GlfwKeyLeftSuper              uint32 = 343
	GlfwKeyRightShift             uint32 = 344
	GlfwKeyRightControl           uint32 = 345
	GlfwKeyRightAlt               uint32 = 346
	GlfwKeyRightSuper             uint32 = 347
	GlfwKeyMenu                   uint32 = 348
	GlfwKeyLast                   uint32 = GlfwKeyMenu
	GlfwModShift                  uint32 = 0x0001
	GlfwModControl                uint32 = 0x0002
	GlfwModAlt                    uint32 = 0x0004
	GlfwModSuper                  uint32 = 0x0008
	GlfwModCapsLock               uint32 = 0x0010
	GlfwModNumLock                uint32 = 0x0020
	GlfwMouseButton1              uint32 = 0
	GlfwMouseButton2              uint32 = 1
	GlfwMouseButton3              uint32 = 2
	GlfwMouseButton4              uint32 = 3
	GlfwMouseButton5              uint32 = 4
	GlfwMouseButton6              uint32 = 5
	GlfwMouseButton7              uint32 = 6
	GlfwMouseButton8              uint32 = 7
	GlfwMouseButtonLast           uint32 = GlfwMouseButton8
	GlfwMouseButtonLeft           uint32 = GlfwMouseButton1
	GlfwMouseButtonRight          uint32 = GlfwMouseButton2
	GlfwMouseButtonMiddle         uint32 = GlfwMouseButton3
	GlfwJoystick1                 uint32 = 0
	GlfwJoystick2                 uint32 = 1
	GlfwJoystick3                 uint32 = 2
	GlfwJoystick4                 uint32 = 3
	GlfwJoystick5                 uint32 = 4
	GlfwJoystick6                 uint32 = 5
	GlfwJoystick7                 uint32 = 6
	GlfwJoystick8                 uint32 = 7
	GlfwJoystick9                 uint32 = 8
	GlfwJoystick10                uint32 = 9
	GlfwJoystick11                uint32 = 10
	GlfwJoystick12                uint32 = 11
	GlfwJoystick13                uint32 = 12
	GlfwJoystick14                uint32 = 13
	GlfwJoystick15                uint32 = 14
	GlfwJoystick16                uint32 = 15
	GlfwJoystickLast              uint32 = GlfwJoystick16
	GlfwGamepadButtonA            uint32 = 0
	GlfwGamepadButtonB            uint32 = 1
	GlfwGamepadButtonX            uint32 = 2
	GlfwGamepadButtonY            uint32 = 3
	GlfwGamepadButtonLeftBumper   uint32 = 4
	GlfwGamepadButtonRightBumper  uint32 = 5
	GlfwGamepadButtonBack         uint32 = 6
	GlfwGamepadButtonStart        uint32 = 7
	GlfwGamepadButtonGuide        uint32 = 8
	GlfwGamepadButtonLeftThumb    uint32 = 9
	GlfwGamepadButtonRightThumb   uint32 = 10
	GlfwGamepadButtonDpadUp       uint32 = 11
	GlfwGamepadButtonDpadRight    uint32 = 12
	GlfwGamepadButtonDpadDown     uint32 = 13
	GlfwGamepadButtonDpadLeft     uint32 = 14
	GlfwGamepadButtonLast         uint32 = GlfwGamepadButtonDpadLeft
	GlfwGamepadButtonCross        uint32 = GlfwGamepadButtonA
	GlfwGamepadButtonCircle       uint32 = GlfwGamepadButtonB
	GlfwGamepadButtonSquare       uint32 = GlfwGamepadButtonX
	GlfwGamepadButtonTriangle     uint32 = GlfwGamepadButtonY
	GlfwGamepadAxisLeftX          uint32 = 0
	GlfwGamepadAxisLeftY          uint32 = 1
	GlfwGamepadAxisRightX         uint32 = 2
	GlfwGamepadAxisRightY         uint32 = 3
	GlfwGamepadAxisLeftTrigger    uint32 = 4
	GlfwGamepadAxisRightTrigger   uint32 = 5
	GlfwGamepadAxisLast           uint32 = GlfwGamepadAxisRightTrigger
	GlfwNoError                   uint32 = 0
	GlfwNotInitialized            uint32 = 0x00010001
	GlfwNoCurrentContext          uint32 = 0x00010002
	GlfwInvalidEnum               uint32 = 0x00010003
	GlfwInvalidValue              uint32 = 0x00010004
	GlfwOutOfMemory               uint32 = 0x00010005
	GlfwApiUnavailable            uint32 = 0x00010006
	GlfwVersionUnavailable        uint32 = 0x00010007
	GlfwPlatformError             uint32 = 0x00010008
	GlfwFormatUnavailable         uint32 = 0x00010009
	GlfwNoWindowContext           uint32 = 0x0001000A
	GlfwCursorUnavailable         uint32 = 0x0001000B
	GlfwFeatureUnavailable        uint32 = 0x0001000C
	GlfwFeatureUnimplemented      uint32 = 0x0001000D
	GlfwPlatformUnavailable       uint32 = 0x0001000E
	GlfwFocused                   uint32 = 0x00020001
	GlfwIconified                 uint32 = 0x00020002
	GlfwResizable                 uint32 = 0x00020003
	GlfwVisible                   uint32 = 0x00020004
	GlfwDecorated                 uint32 = 0x00020005
	GlfwAutoIconify               uint32 = 0x00020006
	GlfwFloating                  uint32 = 0x00020007
	GlfwMaximized                 uint32 = 0x00020008
	GlfwCenterCursor              uint32 = 0x00020009
	GlfwTransparentFramebuffer    uint32 = 0x0002000A
	GlfwHovered                   uint32 = 0x0002000B
	GlfwFocusOnShow               uint32 = 0x0002000C
	GlfwMousePassthrough          uint32 = 0x0002000D
	GlfwPositionX                 uint32 = 0x0002000E
	GlfwPositionY                 uint32 = 0x0002000F
	GlfwRedBits                   uint32 = 0x00021001
	GlfwGreenBits                 uint32 = 0x00021002
	GlfwBlueBits                  uint32 = 0x00021003
	GlfwAlphaBits                 uint32 = 0x00021004
	GlfwDepthBits                 uint32 = 0x00021005
	GlfwStencilBits               uint32 = 0x00021006
	GlfwAccumRedBits              uint32 = 0x00021007
	GlfwAccumGreenBits            uint32 = 0x00021008
	GlfwAccumBlueBits             uint32 = 0x00021009
	GlfwAccumAlphaBits            uint32 = 0x0002100A
	GlfwAuxBuffers                uint32 = 0x0002100B
	GlfwStereo                    uint32 = 0x0002100C
	GlfwSamples                   uint32 = 0x0002100D
	GlfwSrgbCapable               uint32 = 0x0002100E
	GlfwRefreshRate               uint32 = 0x0002100F
	GlfwDoublebuffer              uint32 = 0x00021010
	GlfwClientApi                 uint32 = 0x00022001
	GlfwContextVersionMajor       uint32 = 0x00022002
	GlfwContextVersionMinor       uint32 = 0x00022003
	GlfwContextRevision           uint32 = 0x00022004
	GlfwContextRobustness         uint32 = 0x00022005
	GlfwOpenglForwardCompat       uint32 = 0x00022006
	GlfwContextDebug              uint32 = 0x00022007
	GlfwOpenglDebugContext        uint32 = GlfwContextDebug
	GlfwOpenglProfile             uint32 = 0x00022008
	GlfwContextReleaseBehavior    uint32 = 0x00022009
	GlfwContextNoError            uint32 = 0x0002200A
	GlfwContextCreationApi        uint32 = 0x0002200B
	GlfwScaleToMonitor            uint32 = 0x0002200C
	GlfwScaleFramebuffer          uint32 = 0x0002200D
	GlfwCocoaRetinaFramebuffer    uint32 = 0x00023001
	GlfwCocoaFrameName            uint32 = 0x00023002
	GlfwCocoaGraphicsSwitching    uint32 = 0x00023003
	GlfwX11ClassName              uint32 = 0x00024001
	GlfwX11InstanceName           uint32 = 0x00024002
	GlfwWin32KeyboardMenu         uint32 = 0x00025001
	GlfwWin32Showdefault          uint32 = 0x00025002
	GlfwWaylandAppId              uint32 = 0x00026001
	GlfwNoApi                     uint32 = 0
	GlfwOpenglApi                 uint32 = 0x00030001
	GlfwOpenglEsApi               uint32 = 0x00030002
	GlfwNoRobustness              uint32 = 0
	GlfwNoResetNotification       uint32 = 0x00031001
	GlfwLoseContextOnReset        uint32 = 0x00031002
	GlfwOpenglAnyProfile          uint32 = 0
	GlfwOpenglCoreProfile         uint32 = 0x00032001
	GlfwOpenglCompatProfile       uint32 = 0x00032002
	GlfwCursor                    uint32 = 0x00033001
	GlfwStickyKeys                uint32 = 0x00033002
	GlfwStickyMouseButtons        uint32 = 0x00033003
	GlfwLockKeyMods               uint32 = 0x00033004
	GlfwRawMouseMotion            uint32 = 0x00033005
	GlfwUnlimitedMouseButtons     uint32 = 0x00033006
	GlfwCursorNormal              uint32 = 0x00034001
	GlfwCursorHidden              uint32 = 0x00034002
	GlfwCursorDisabled            uint32 = 0x00034003
	GlfwCursorCaptured            uint32 = 0x00034004
	GlfwAnyReleaseBehavior        uint32 = 0
	GlfwReleaseBehaviorFlush      uint32 = 0x00035001
	GlfwReleaseBehaviorNone       uint32 = 0x00035002
	GlfwNativeContextApi          uint32 = 0x00036001
	GlfwEglContextApi             uint32 = 0x00036002
	GlfwOsmesaContextApi          uint32 = 0x00036003
	GlfwAnglePlatformTypeNone     uint32 = 0x00037001
	GlfwAnglePlatformTypeOpengl   uint32 = 0x00037002
	GlfwAnglePlatformTypeOpengles uint32 = 0x00037003
	GlfwAnglePlatformTypeD3d9     uint32 = 0x00037004
	GlfwAnglePlatformTypeD3d11    uint32 = 0x00037005
	GlfwAnglePlatformTypeVulkan   uint32 = 0x00037007
	GlfwAnglePlatformTypeMetal    uint32 = 0x00037008
	GlfwWaylandPreferLibdecor     uint32 = 0x00038001
	GlfwWaylandDisableLibdecor    uint32 = 0x00038002
	GlfwAnyPosition               uint32 = 0x80000000
	GlfwArrowCursor               uint32 = 0x00036001
	GlfwIbeamCursor               uint32 = 0x00036002
	GlfwCrosshairCursor           uint32 = 0x00036003
	GlfwPointingHandCursor        uint32 = 0x00036004
	GlfwResizeEwCursor            uint32 = 0x00036005
	GlfwResizeNsCursor            uint32 = 0x00036006
	GlfwResizeNwseCursor          uint32 = 0x00036007
	GlfwResizeNeswCursor          uint32 = 0x00036008
	GlfwResizeAllCursor           uint32 = 0x00036009
	GlfwNotAllowedCursor          uint32 = 0x0003600A
	GlfwHresizeCursor             uint32 = GlfwResizeEwCursor
	GlfwVresizeCursor             uint32 = GlfwResizeNsCursor
	GlfwConnected                 uint32 = 0x00040001
	GlfwDisconnected              uint32 = 0x00040002
	GlfwJoystickHatButtons        uint32 = 0x00050001
	GlfwAnglePlatformType         uint32 = 0x00050002
	GlfwPlatform                  uint32 = 0x00050003
	GlfwCocoaChdirResources       uint32 = 0x00051001
	GlfwCocoaMenubar              uint32 = 0x00051002
	GlfwX11XcbVulkanSurface       uint32 = 0x00052001
	GlfwWaylandLibdecor           uint32 = 0x00053001
	GlfwAnyPlatform               uint32 = 0x00060000
	GlfwPlatformWin32             uint32 = 0x00060001
	GlfwPlatformCocoa             uint32 = 0x00060002
	GlfwPlatformWayland           uint32 = 0x00060003
	GlfwPlatformX11               uint32 = 0x00060004
	GlfwPlatformNull              uint32 = 0x00060005
	GlfwDontCare                  int32  = -1
)

func (g *Glfw) Init() int32 {
	r1, _, _ := getProc("glfwInit").Call()
	return int32(r1)
}

func (g *Glfw) Terminate() {
	getProc("glfwTerminate").Call()
}

func (g *Glfw) InitHint(Hint int32, Value int32) {
	getProc("glfwInitHint").Call(uintptr(Hint), uintptr(Value))
}

func (g *Glfw) InitAllocator(Allocator *GLFWallocator) {
	getProc("glfwInitAllocator").Call(uintptr(unsafe.Pointer(Allocator)))
}

func (g *Glfw) GetVersion(Major *int32, Minor *int32, Rev *int32) {
	getProc("glfwGetVersion").Call(uintptr(unsafe.Pointer(Major)), uintptr(unsafe.Pointer(Minor)), uintptr(unsafe.Pointer(Rev)))
}

func (g *Glfw) GetVersionString() *int8 {
	r1, _, _ := getProc("glfwGetVersionString").Call()
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetError(Description **int8) int32 {
	r1, _, _ := getProc("glfwGetError").Call(uintptr(unsafe.Pointer(Description)))
	return int32(r1)
}

func (g *Glfw) SetErrorCallback(Callback GLFWerrorfun) GLFWerrorfun {
	r1, _, _ := getProc("glfwSetErrorCallback").Call(func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWerrorfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) GetPlatform() int32 {
	r1, _, _ := getProc("glfwGetPlatform").Call()
	return int32(r1)
}

func (g *Glfw) PlatformSupported(Platform int32) int32 {
	r1, _, _ := getProc("glfwPlatformSupported").Call(uintptr(Platform))
	return int32(r1)
}

func (g *Glfw) GetMonitors(Count *int32) **GLFWmonitor {
	r1, _, _ := getProc("glfwGetMonitors").Call(uintptr(unsafe.Pointer(Count)))
	return (**GLFWmonitor)(unsafe.Pointer(r1))
}

func (g *Glfw) GetPrimaryMonitor() *GLFWmonitor {
	r1, _, _ := getProc("glfwGetPrimaryMonitor").Call()
	return (*GLFWmonitor)(unsafe.Pointer(r1))
}

func (g *Glfw) GetMonitorPos(Monitor *GLFWmonitor, Xpos *int32, Ypos *int32) {
	getProc("glfwGetMonitorPos").Call(uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(Xpos)), uintptr(unsafe.Pointer(Ypos)))
}

func (g *Glfw) GetMonitorWorkarea(Monitor *GLFWmonitor, Xpos *int32, Ypos *int32, Width *int32, Height *int32) {
	getProc("glfwGetMonitorWorkarea").Call(uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(Xpos)), uintptr(unsafe.Pointer(Ypos)), uintptr(unsafe.Pointer(Width)), uintptr(unsafe.Pointer(Height)))
}

func (g *Glfw) GetMonitorPhysicalSize(Monitor *GLFWmonitor, WidthMM *int32, HeightMM *int32) {
	getProc("glfwGetMonitorPhysicalSize").Call(uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(WidthMM)), uintptr(unsafe.Pointer(HeightMM)))
}

func (g *Glfw) GetMonitorContentScale(Monitor *GLFWmonitor, Xscale *float32, Yscale *float32) {
	getProc("glfwGetMonitorContentScale").Call(uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(Xscale)), uintptr(unsafe.Pointer(Yscale)))
}

func (g *Glfw) GetMonitorName(Monitor *GLFWmonitor) *int8 {
	r1, _, _ := getProc("glfwGetMonitorName").Call(uintptr(unsafe.Pointer(Monitor)))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) SetMonitorUserPointer(Monitor *GLFWmonitor, Pointer unsafe.Pointer) {
	getProc("glfwSetMonitorUserPointer").Call(uintptr(unsafe.Pointer(Monitor)), uintptr(Pointer))
}

func (g *Glfw) GetMonitorUserPointer(Monitor *GLFWmonitor) unsafe.Pointer {
	r1, _, _ := getProc("glfwGetMonitorUserPointer").Call(uintptr(unsafe.Pointer(Monitor)))
	return unsafe.Pointer(r1)
}

func (g *Glfw) SetMonitorCallback(Callback GLFWmonitorfun) GLFWmonitorfun {
	r1, _, _ := getProc("glfwSetMonitorCallback").Call(func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWmonitorfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) GetVideoModes(Monitor *GLFWmonitor, Count *int32) *GLFWvidmode {
	r1, _, _ := getProc("glfwGetVideoModes").Call(uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(Count)))
	return (*GLFWvidmode)(unsafe.Pointer(r1))
}

func (g *Glfw) GetVideoMode(Monitor *GLFWmonitor) *GLFWvidmode {
	r1, _, _ := getProc("glfwGetVideoMode").Call(uintptr(unsafe.Pointer(Monitor)))
	return (*GLFWvidmode)(unsafe.Pointer(r1))
}

func (g *Glfw) SetGamma(Monitor *GLFWmonitor, Gamma float32) {
	getProc("glfwSetGamma").Call(uintptr(unsafe.Pointer(Monitor)), uintptr(Gamma))
}

func (g *Glfw) GetGammaRamp(Monitor *GLFWmonitor) *GLFWgammaramp {
	r1, _, _ := getProc("glfwGetGammaRamp").Call(uintptr(unsafe.Pointer(Monitor)))
	return (*GLFWgammaramp)(unsafe.Pointer(r1))
}

func (g *Glfw) SetGammaRamp(Monitor *GLFWmonitor, Ramp *GLFWgammaramp) {
	getProc("glfwSetGammaRamp").Call(uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(Ramp)))
}

func (g *Glfw) DefaultWindowHints() {
	getProc("glfwDefaultWindowHints").Call()
}

func (g *Glfw) WindowHint(Hint int32, Value int32) {
	getProc("glfwWindowHint").Call(uintptr(Hint), uintptr(Value))
}

func (g *Glfw) WindowHintString(Hint int32, Value *int8) {
	getProc("glfwWindowHintString").Call(uintptr(Hint), uintptr(unsafe.Pointer(Value)))
}

func (g *Glfw) CreateWindow(Width int32, Height int32, Title *int8, Monitor *GLFWmonitor, Share *GLFWwindow) *GLFWwindow {
	r1, _, _ := getProc("glfwCreateWindow").Call(uintptr(Width), uintptr(Height), uintptr(unsafe.Pointer(Title)), uintptr(unsafe.Pointer(Monitor)), uintptr(unsafe.Pointer(Share)))
	return (*GLFWwindow)(unsafe.Pointer(r1))
}

func (g *Glfw) DestroyWindow(Window *GLFWwindow) {
	getProc("glfwDestroyWindow").Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) WindowShouldClose(Window *GLFWwindow) int32 {
	r1, _, _ := getProc("glfwWindowShouldClose").Call(uintptr(unsafe.Pointer(Window)))
	return int32(r1)
}

func (g *Glfw) SetWindowShouldClose(Window *GLFWwindow, Value int32) {
	getProc("glfwSetWindowShouldClose").Call(uintptr(unsafe.Pointer(Window)), uintptr(Value))
}

func (g *Glfw) GetWindowTitle(Window *GLFWwindow) *int8 {
	r1, _, _ := getProc("glfwGetWindowTitle").Call(uintptr(unsafe.Pointer(Window)))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) SetWindowTitle(Window *GLFWwindow, Title *int8) {
	getProc("glfwSetWindowTitle").Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Title)))
}

func (g *Glfw) SetWindowIcon(Window *GLFWwindow, Count int32, Images *GLFWimage) {
	getProc("glfwSetWindowIcon").Call(uintptr(unsafe.Pointer(Window)), uintptr(Count), uintptr(unsafe.Pointer(Images)))
}

func (g *Glfw) GetWindowPos(Window *GLFWwindow, Xpos *int32, Ypos *int32) {
	getProc("glfwGetWindowPos").Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Xpos)), uintptr(unsafe.Pointer(Ypos)))
}

func (g *Glfw) SetWindowPos(Window *GLFWwindow, Xpos int32, Ypos int32) {
	getProc("glfwSetWindowPos").Call(uintptr(unsafe.Pointer(Window)), uintptr(Xpos), uintptr(Ypos))
}

func (g *Glfw) GetWindowSize(Window *GLFWwindow, Width *int32, Height *int32) {
	getProc("glfwGetWindowSize").Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Width)), uintptr(unsafe.Pointer(Height)))
}

func (g *Glfw) SetWindowSizeLimits(Window *GLFWwindow, Minwidth int32, Minheight int32, Maxwidth int32, Maxheight int32) {
	getProc("glfwSetWindowSizeLimits").Call(uintptr(unsafe.Pointer(Window)), uintptr(Minwidth), uintptr(Minheight), uintptr(Maxwidth), uintptr(Maxheight))
}

func (g *Glfw) SetWindowAspectRatio(Window *GLFWwindow, Numer int32, Denom int32) {
	getProc("glfwSetWindowAspectRatio").Call(uintptr(unsafe.Pointer(Window)), uintptr(Numer), uintptr(Denom))
}

func (g *Glfw) SetWindowSize(Window *GLFWwindow, Width int32, Height int32) {
	getProc("glfwSetWindowSize").Call(uintptr(unsafe.Pointer(Window)), uintptr(Width), uintptr(Height))
}

func (g *Glfw) GetFramebufferSize(Window *GLFWwindow, Width *int32, Height *int32) {
	getProc("glfwGetFramebufferSize").Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Width)), uintptr(unsafe.Pointer(Height)))
}

func (g *Glfw) GetWindowFrameSize(Window *GLFWwindow, Left *int32, Top *int32, Right *int32, Bottom *int32) {
	getProc("glfwGetWindowFrameSize").Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Left)), uintptr(unsafe.Pointer(Top)), uintptr(unsafe.Pointer(Right)), uintptr(unsafe.Pointer(Bottom)))
}

func (g *Glfw) GetWindowContentScale(Window *GLFWwindow, Xscale *float32, Yscale *float32) {
	getProc("glfwGetWindowContentScale").Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Xscale)), uintptr(unsafe.Pointer(Yscale)))
}

func (g *Glfw) GetWindowOpacity(Window *GLFWwindow) float32 {
	r1, _, _ := getProc("glfwGetWindowOpacity").Call(uintptr(unsafe.Pointer(Window)))
	return *(*float32)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowOpacity(Window *GLFWwindow, Opacity float32) {
	getProc("glfwSetWindowOpacity").Call(uintptr(unsafe.Pointer(Window)), uintptr(Opacity))
}

func (g *Glfw) IconifyWindow(Window *GLFWwindow) {
	getProc("glfwIconifyWindow").Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) RestoreWindow(Window *GLFWwindow) {
	getProc("glfwRestoreWindow").Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) MaximizeWindow(Window *GLFWwindow) {
	getProc("glfwMaximizeWindow").Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) ShowWindow(Window *GLFWwindow) {
	getProc("glfwShowWindow").Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) HideWindow(Window *GLFWwindow) {
	getProc("glfwHideWindow").Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) FocusWindow(Window *GLFWwindow) {
	getProc("glfwFocusWindow").Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) RequestWindowAttention(Window *GLFWwindow) {
	getProc("glfwRequestWindowAttention").Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) GetWindowMonitor(Window *GLFWwindow) *GLFWmonitor {
	r1, _, _ := getProc("glfwGetWindowMonitor").Call(uintptr(unsafe.Pointer(Window)))
	return (*GLFWmonitor)(unsafe.Pointer(r1))
}

func (g *Glfw) SetWindowMonitor(Window *GLFWwindow, Monitor *GLFWmonitor, Xpos int32, Ypos int32, Width int32, Height int32, RefreshRate int32) {
	getProc("glfwSetWindowMonitor").Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Monitor)), uintptr(Xpos), uintptr(Ypos), uintptr(Width), uintptr(Height), uintptr(RefreshRate))
}

func (g *Glfw) GetWindowAttrib(Window *GLFWwindow, Attrib int32) int32 {
	r1, _, _ := getProc("glfwGetWindowAttrib").Call(uintptr(unsafe.Pointer(Window)), uintptr(Attrib))
	return int32(r1)
}

func (g *Glfw) SetWindowAttrib(Window *GLFWwindow, Attrib int32, Value int32) {
	getProc("glfwSetWindowAttrib").Call(uintptr(unsafe.Pointer(Window)), uintptr(Attrib), uintptr(Value))
}

func (g *Glfw) SetWindowUserPointer(Window *GLFWwindow, Pointer unsafe.Pointer) {
	getProc("glfwSetWindowUserPointer").Call(uintptr(unsafe.Pointer(Window)), uintptr(Pointer))
}

func (g *Glfw) GetWindowUserPointer(Window *GLFWwindow) unsafe.Pointer {
	r1, _, _ := getProc("glfwGetWindowUserPointer").Call(uintptr(unsafe.Pointer(Window)))
	return unsafe.Pointer(r1)
}

func (g *Glfw) SetWindowPosCallback(Window *GLFWwindow, Callback GLFWwindowposfun) GLFWwindowposfun {
	r1, _, _ := getProc("glfwSetWindowPosCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowposfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowSizeCallback(Window *GLFWwindow, Callback GLFWwindowsizefun) GLFWwindowsizefun {
	r1, _, _ := getProc("glfwSetWindowSizeCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowsizefun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowCloseCallback(Window *GLFWwindow, Callback GLFWwindowclosefun) GLFWwindowclosefun {
	r1, _, _ := getProc("glfwSetWindowCloseCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowclosefun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowRefreshCallback(Window *GLFWwindow, Callback GLFWwindowrefreshfun) GLFWwindowrefreshfun {
	r1, _, _ := getProc("glfwSetWindowRefreshCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowrefreshfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowFocusCallback(Window *GLFWwindow, Callback GLFWwindowfocusfun) GLFWwindowfocusfun {
	r1, _, _ := getProc("glfwSetWindowFocusCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowfocusfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowIconifyCallback(Window *GLFWwindow, Callback GLFWwindowiconifyfun) GLFWwindowiconifyfun {
	r1, _, _ := getProc("glfwSetWindowIconifyCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowiconifyfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowMaximizeCallback(Window *GLFWwindow, Callback GLFWwindowmaximizefun) GLFWwindowmaximizefun {
	r1, _, _ := getProc("glfwSetWindowMaximizeCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowmaximizefun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetFramebufferSizeCallback(Window *GLFWwindow, Callback GLFWframebuffersizefun) GLFWframebuffersizefun {
	r1, _, _ := getProc("glfwSetFramebufferSizeCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWframebuffersizefun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetWindowContentScaleCallback(Window *GLFWwindow, Callback GLFWwindowcontentscalefun) GLFWwindowcontentscalefun {
	r1, _, _ := getProc("glfwSetWindowContentScaleCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWwindowcontentscalefun)(unsafe.Pointer(&r1))
}

func (g *Glfw) PollEvents() {
	getProc("glfwPollEvents").Call()
}

func (g *Glfw) WaitEvents() {
	getProc("glfwWaitEvents").Call()
}

func (g *Glfw) WaitEventsTimeout(Timeout float64) {
	getProc("glfwWaitEventsTimeout").Call(uintptr(Timeout))
}

func (g *Glfw) PostEmptyEvent() {
	getProc("glfwPostEmptyEvent").Call()
}

func (g *Glfw) GetInputMode(Window *GLFWwindow, Mode int32) int32 {
	r1, _, _ := getProc("glfwGetInputMode").Call(uintptr(unsafe.Pointer(Window)), uintptr(Mode))
	return int32(r1)
}

func (g *Glfw) SetInputMode(Window *GLFWwindow, Mode int32, Value int32) {
	getProc("glfwSetInputMode").Call(uintptr(unsafe.Pointer(Window)), uintptr(Mode), uintptr(Value))
}

func (g *Glfw) RawMouseMotionSupported() int32 {
	r1, _, _ := getProc("glfwRawMouseMotionSupported").Call()
	return int32(r1)
}

func (g *Glfw) GetKeyName(Key int32, Scancode int32) *int8 {
	r1, _, _ := getProc("glfwGetKeyName").Call(uintptr(Key), uintptr(Scancode))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetKeyScancode(Key int32) int32 {
	r1, _, _ := getProc("glfwGetKeyScancode").Call(uintptr(Key))
	return int32(r1)
}

func (g *Glfw) GetKey(Window *GLFWwindow, Key int32) int32 {
	r1, _, _ := getProc("glfwGetKey").Call(uintptr(unsafe.Pointer(Window)), uintptr(Key))
	return int32(r1)
}

func (g *Glfw) GetMouseButton(Window *GLFWwindow, Button int32) int32 {
	r1, _, _ := getProc("glfwGetMouseButton").Call(uintptr(unsafe.Pointer(Window)), uintptr(Button))
	return int32(r1)
}

func (g *Glfw) GetCursorPos(Window *GLFWwindow, Xpos *float64, Ypos *float64) {
	getProc("glfwGetCursorPos").Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Xpos)), uintptr(unsafe.Pointer(Ypos)))
}

func (g *Glfw) SetCursorPos(Window *GLFWwindow, Xpos float64, Ypos float64) {
	getProc("glfwSetCursorPos").Call(uintptr(unsafe.Pointer(Window)), uintptr(Xpos), uintptr(Ypos))
}

func (g *Glfw) CreateCursor(Image *GLFWimage, Xhot int32, Yhot int32) *GLFWcursor {
	r1, _, _ := getProc("glfwCreateCursor").Call(uintptr(unsafe.Pointer(Image)), uintptr(Xhot), uintptr(Yhot))
	return (*GLFWcursor)(unsafe.Pointer(r1))
}

func (g *Glfw) CreateStandardCursor(Shape int32) *GLFWcursor {
	r1, _, _ := getProc("glfwCreateStandardCursor").Call(uintptr(Shape))
	return (*GLFWcursor)(unsafe.Pointer(r1))
}

func (g *Glfw) DestroyCursor(Cursor *GLFWcursor) {
	getProc("glfwDestroyCursor").Call(uintptr(unsafe.Pointer(Cursor)))
}

func (g *Glfw) SetCursor(Window *GLFWwindow, Cursor *GLFWcursor) {
	getProc("glfwSetCursor").Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(Cursor)))
}

func (g *Glfw) SetKeyCallback(Window *GLFWwindow, Callback GLFWkeyfun) GLFWkeyfun {
	r1, _, _ := getProc("glfwSetKeyCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWkeyfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetCharCallback(Window *GLFWwindow, Callback GLFWcharfun) GLFWcharfun {
	r1, _, _ := getProc("glfwSetCharCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWcharfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetCharModsCallback(Window *GLFWwindow, Callback GLFWcharmodsfun) GLFWcharmodsfun {
	r1, _, _ := getProc("glfwSetCharModsCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWcharmodsfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetMouseButtonCallback(Window *GLFWwindow, Callback GLFWmousebuttonfun) GLFWmousebuttonfun {
	r1, _, _ := getProc("glfwSetMouseButtonCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWmousebuttonfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetCursorPosCallback(Window *GLFWwindow, Callback GLFWcursorposfun) GLFWcursorposfun {
	r1, _, _ := getProc("glfwSetCursorPosCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWcursorposfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetCursorEnterCallback(Window *GLFWwindow, Callback GLFWcursorenterfun) GLFWcursorenterfun {
	r1, _, _ := getProc("glfwSetCursorEnterCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWcursorenterfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetScrollCallback(Window *GLFWwindow, Callback GLFWscrollfun) GLFWscrollfun {
	r1, _, _ := getProc("glfwSetScrollCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWscrollfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetDropCallback(Window *GLFWwindow, Callback GLFWdropfun) GLFWdropfun {
	r1, _, _ := getProc("glfwSetDropCallback").Call(uintptr(unsafe.Pointer(Window)), func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWdropfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) JoystickPresent(Jid int32) int32 {
	r1, _, _ := getProc("glfwJoystickPresent").Call(uintptr(Jid))
	return int32(r1)
}

func (g *Glfw) GetJoystickAxes(Jid int32, Count *int32) *float32 {
	r1, _, _ := getProc("glfwGetJoystickAxes").Call(uintptr(Jid), uintptr(unsafe.Pointer(Count)))
	return (*float32)(unsafe.Pointer(r1))
}

func (g *Glfw) GetJoystickButtons(Jid int32, Count *int32) *uint8 {
	r1, _, _ := getProc("glfwGetJoystickButtons").Call(uintptr(Jid), uintptr(unsafe.Pointer(Count)))
	return (*uint8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetJoystickHats(Jid int32, Count *int32) *uint8 {
	r1, _, _ := getProc("glfwGetJoystickHats").Call(uintptr(Jid), uintptr(unsafe.Pointer(Count)))
	return (*uint8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetJoystickName(Jid int32) *int8 {
	r1, _, _ := getProc("glfwGetJoystickName").Call(uintptr(Jid))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetJoystickGUID(Jid int32) *int8 {
	r1, _, _ := getProc("glfwGetJoystickGUID").Call(uintptr(Jid))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) SetJoystickUserPointer(Jid int32, Pointer unsafe.Pointer) {
	getProc("glfwSetJoystickUserPointer").Call(uintptr(Jid), uintptr(Pointer))
}

func (g *Glfw) GetJoystickUserPointer(Jid int32) unsafe.Pointer {
	r1, _, _ := getProc("glfwGetJoystickUserPointer").Call(uintptr(Jid))
	return unsafe.Pointer(r1)
}

func (g *Glfw) JoystickIsGamepad(Jid int32) int32 {
	r1, _, _ := getProc("glfwJoystickIsGamepad").Call(uintptr(Jid))
	return int32(r1)
}

func (g *Glfw) SetJoystickCallback(Callback GLFWjoystickfun) GLFWjoystickfun {
	r1, _, _ := getProc("glfwSetJoystickCallback").Call(func() uintptr {
		if Callback == nil {
			println("Callback is nil")
			return 0
		}
		return windows.NewCallback(Callback)
	}())
	return *(*GLFWjoystickfun)(unsafe.Pointer(&r1))
}

func (g *Glfw) UpdateGamepadMappings(String *int8) int32 {
	r1, _, _ := getProc("glfwUpdateGamepadMappings").Call(uintptr(unsafe.Pointer(String)))
	return int32(r1)
}

func (g *Glfw) GetGamepadName(Jid int32) *int8 {
	r1, _, _ := getProc("glfwGetGamepadName").Call(uintptr(Jid))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetGamepadState(Jid int32, State *GLFWgamepadstate) int32 {
	r1, _, _ := getProc("glfwGetGamepadState").Call(uintptr(Jid), uintptr(unsafe.Pointer(State)))
	return int32(r1)
}

func (g *Glfw) SetClipboardString(Window *GLFWwindow, String *int8) {
	getProc("glfwSetClipboardString").Call(uintptr(unsafe.Pointer(Window)), uintptr(unsafe.Pointer(String)))
}

func (g *Glfw) GetClipboardString(Window *GLFWwindow) *int8 {
	r1, _, _ := getProc("glfwGetClipboardString").Call(uintptr(unsafe.Pointer(Window)))
	return (*int8)(unsafe.Pointer(r1))
}

func (g *Glfw) GetTime() float64 {
	r1, _, _ := getProc("glfwGetTime").Call()
	return *(*float64)(unsafe.Pointer(&r1))
}

func (g *Glfw) SetTime(Time float64) {
	getProc("glfwSetTime").Call(uintptr(Time))
}

func (g *Glfw) GetTimerValue() uint64 {
	r1, _, _ := getProc("glfwGetTimerValue").Call()
	return *(*uint64)(unsafe.Pointer(&r1))
}

func (g *Glfw) GetTimerFrequency() uint64 {
	r1, _, _ := getProc("glfwGetTimerFrequency").Call()
	return *(*uint64)(unsafe.Pointer(&r1))
}

func (g *Glfw) MakeContextCurrent(Window *GLFWwindow) {
	getProc("glfwMakeContextCurrent").Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) GetCurrentContext() *GLFWwindow {
	r1, _, _ := getProc("glfwGetCurrentContext").Call()
	return (*GLFWwindow)(unsafe.Pointer(r1))
}

func (g *Glfw) SwapBuffers(Window *GLFWwindow) {
	getProc("glfwSwapBuffers").Call(uintptr(unsafe.Pointer(Window)))
}

func (g *Glfw) SwapInterval(Interval int32) {
	getProc("glfwSwapInterval").Call(uintptr(Interval))
}

func (g *Glfw) ExtensionSupported(Extension *int8) int32 {
	r1, _, _ := getProc("glfwExtensionSupported").Call(uintptr(unsafe.Pointer(Extension)))
	return int32(r1)
}

func (g *Glfw) GetProcAddress(Procname *int8) GLFWglproc {
	r1, _, _ := getProc("glfwGetProcAddress").Call(uintptr(unsafe.Pointer(Procname)))
	return *(*GLFWglproc)(unsafe.Pointer(&r1))
}

func (g *Glfw) VulkanSupported() int32 {
	r1, _, _ := getProc("glfwVulkanSupported").Call()
	return int32(r1)
}

func (g *Glfw) GetRequiredInstanceExtensions(Count *uint32) **int8 {
	r1, _, _ := getProc("glfwGetRequiredInstanceExtensions").Call(uintptr(unsafe.Pointer(Count)))
	return (**int8)(unsafe.Pointer(r1))
}
