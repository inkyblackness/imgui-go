package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/inkyblackness/imgui-go"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	runtime.LockOSThread()

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("ImGui-Go SDL2+OpenGL2 example", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 1280, 720, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	contextGl, err := window.GLCreateContext()
	if err != nil {
		panic(err)
	}

	contextImgui := imgui.CreateContext(nil)
	defer contextImgui.Destroy()

	impl := NewImplSdl2Gl2(window)
	defer impl.Shutdown()

	var (
		showDemoWindow    bool
		showAnotherWindow bool
		counter           int
		clearColor        imgui.Vec4
	)

	for running := true; running; {
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

		gl.ClearColor(clearColor.X, clearColor.Y, clearColor.Z, clearColor.W)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		imgui.Render()
		window.GLMakeCurrent(contextGl)
		impl.Render(imgui.RenderedDrawData())
		window.GLSwap()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			if impl.ProcessEvent(event) {
				continue
			}

			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		<-time.After(time.Millisecond * 25)
	}
}
