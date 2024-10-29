package main

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/nitrix/glfw-go"
	"github.com/nitrix/imgui-go"
	bglfw "github.com/nitrix/imgui-go/backends/glfw"
	bopengl3 "github.com/nitrix/imgui-go/backends/opengl3"
	"github.com/nitrix/imguizmo-go"
)

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Visible, glfw.False)

	window, err := glfw.CreateWindow(1280, 720, "Example", nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	glfw.SwapInterval(1)

	ctx := imgui.CreateContext(nil)
	if ctx == nil {
		panic("CreateContext failed")
	}
	defer imgui.DestroyContext(ctx)

	ctx.IO.IniFilename = nil

	bglfw.Init(window)
	defer bglfw.Shutdown()

	bopengl3.Init(window)
	defer bopengl3.Shutdown()

	window.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		gl.Viewport(0, 0, int32(width), int32(height))
	})

	window.Show()

	var matrix = mgl32.Ident4()
	var deltaMatrix = mgl32.Ident4()

	for !window.ShouldClose() {
		glfw.PollEvents()

		gl.ClearColor(0.1, 0.1, 0.1, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		bglfw.NewFrame()
		bopengl3.NewFrame()
		imgui.NewFrame()
		imguizmo.BeginFrame()

		fbWidth, fbHeight := window.GetFramebufferSize()
		view := mgl32.LookAtV(mgl32.Vec3{0, 0, 2}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
		projection := mgl32.Perspective(mgl32.DegToRad(90), float32(fbWidth)/float32(fbHeight), 0.1, 100.0)

		imguizmo.Enable(true)
		imguizmo.SetOrthographic(false)
		imguizmo.AllowAxisFlip(true)

		mainViewport := imgui.GetMainViewport()

		// viewportSize := mgl32.Vec2{mainViewport.Size.X, mainViewport.Size.Y}
		// viewportPos := mgl32.Vec2{mainViewport.Pos.X, mainViewport.Pos.Y}
		// imgui.SetNextWindowSize(mainViewport.Size, 0)
		// imgui.SetNextWindowPos(mainViewport.Pos, 0, mgl32.Vec2{0, 0})

		// imgui.SetNextWindowSize(mgl32.Vec2{float32(1280), float32(720)}, 0)
		// imgui.SetNextWindowPos(mgl32.Vec2{0, 0}, 0, mgl32.Vec2{0, 0})

		imguizmo.SetDrawlist(imgui.GetForegroundDrawList_ViewportPtr(mainViewport))

		// winWidth, winHeight := imgui.GetWindowWidth(), imgui.GetWindowHeight()
		var winPos mgl32.Vec2
		imgui.GetWindowPos(&winPos)
		imguizmo.SetRect(winPos.X(), winPos.Y(), float32(fbWidth), float32(fbHeight))

		if imguizmo.Manipulate(view, projection, imguizmo.Translate, imguizmo.World, &matrix, &deltaMatrix) {
			fmt.Println("Manipulated", matrix)
		}

		imgui.Render()

		bopengl3.RenderDrawData(imgui.GetDrawData())

		window.SwapBuffers()
	}
}
