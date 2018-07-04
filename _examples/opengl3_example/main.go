package main

import (
	"runtime"
	"time"

	"fmt"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/inkyblackness/imgui-go"
)

func main() {
	runtime.LockOSThread()

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

	showDemoWindow := false
	showAnotherWindow := false
	counter := 0
	var clearColor imgui.Vec4

	for !window.ShouldClose() {
		glfw.PollEvents()
		impl.NewFrame()

		// 1. Show a simple window.
		// Tip: if we don't call ImGui::Begin()/ImGui::End() the widgets automatically appears in a window called "Debug".
		{
			imgui.Text("Hello, world!")

			imgui.Checkbox("Demo Window", &showDemoWindow)
			imgui.Checkbox("Another Window", &showAnotherWindow)

			if imgui.Button("Button") {
				counter++
			}
			imgui.SameLine()
			imgui.Text(fmt.Sprintf("counter = %d", counter))

		}

		// 2. Show another simple window. In most cases you will use an explicit Begin/End pair to name your windows.
		if showAnotherWindow {
			imgui.BeginV("Another Window", &showAnotherWindow, 0)
			imgui.Text("Hello from another window!")
			if imgui.Button("Close Me") {
				showAnotherWindow = false
			}
			imgui.End()
		}

		// 3. Show the ImGui demo window. Most of the sample code is in imgui.ShowDemoWindow().
		// Read its code to learn more about Dear ImGui!
		if showDemoWindow {
			imgui.ShowDemoWindow(&showDemoWindow)
		}

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
