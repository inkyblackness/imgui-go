package main

import (
	"time"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/inkyblackness/imgui-go"
)

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, 1)

	window, err := glfw.CreateWindow(1280, 720, "ImGui-Go GLFW+OpenGL3 example", nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	gl.Init()

	context := imgui.CreateContext(nil)
	defer context.Destroy()

	/*
		imgui.CurrentStyle().ScaleAllSizes(2.0)
		imgui.CurrentIO().SetFontGlobalScale(2.0)
	*/

	impl := imguiGlfw3Init(window)
	defer impl.Shutdown()

	var clearColor imgui.Vec4

	for !window.ShouldClose() {
		glfw.PollEvents()
		impl.NewFrame()

		imgui.ShowDemoWindow(nil)

		displayWidth, displayHeight := window.GetFramebufferSize()
		gl.Viewport(0, 0, int32(displayWidth), int32(displayHeight))
		gl.ClearColor(clearColor.X, clearColor.Y, clearColor.Z, clearColor.W)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		imgui.Render()
		impl.Render(imgui.RenderedDrawData())

		window.SwapBuffers()
		<-time.After(time.Millisecond * 25)
	}
}
