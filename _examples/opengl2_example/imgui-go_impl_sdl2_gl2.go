package main

import (
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/inkyblackness/imgui-go"
	"github.com/veandco/go-sdl2/sdl"
)

type ImplSdl2Gl2 struct {
	renderer *rendererGl2
	platform *platformSdl2
}

func NewImplSdl2Gl2(window *sdl.Window) *ImplSdl2Gl2 {
	return &ImplSdl2Gl2{
		renderer: newRendererGl2(window),
		platform: newPlatformSdl2(window),
	}
}

func (impl *ImplSdl2Gl2) NewFrame() {
	impl.platform.newFrame()
	impl.renderer.newFrame()
}

func (impl *ImplSdl2Gl2) Render(drawData imgui.DrawData) {
	impl.renderer.render(drawData)
}

func (impl *ImplSdl2Gl2) ProcessEvent(event sdl.Event) bool {
	return impl.platform.processEvent(event)
}

func (impl *ImplSdl2Gl2) Shutdown() {
	impl.renderer.shutdown()
}

type platformSdl2 struct {
	window      *sdl.Window
	time        uint64
	buttonsDown [3]bool
}

func newPlatformSdl2(window *sdl.Window) *platformSdl2 {
	keys := map[int]int{
		imgui.KeyTab:        sdl.SCANCODE_TAB,
		imgui.KeyLeftArrow:  sdl.SCANCODE_LEFT,
		imgui.KeyRightArrow: sdl.SCANCODE_RIGHT,
		imgui.KeyUpArrow:    sdl.SCANCODE_UP,
		imgui.KeyDownArrow:  sdl.SCANCODE_DOWN,
		imgui.KeyPageUp:     sdl.SCANCODE_PAGEUP,
		imgui.KeyPageDown:   sdl.SCANCODE_PAGEDOWN,
		imgui.KeyHome:       sdl.SCANCODE_HOME,
		imgui.KeyEnd:        sdl.SCANCODE_END,
		imgui.KeyInsert:     sdl.SCANCODE_INSERT,
		imgui.KeyDelete:     sdl.SCANCODE_DELETE,
		imgui.KeyBackspace:  sdl.SCANCODE_BACKSPACE,
		imgui.KeySpace:      sdl.SCANCODE_BACKSPACE,
		imgui.KeyEnter:      sdl.SCANCODE_RETURN,
		imgui.KeyEscape:     sdl.SCANCODE_ESCAPE,
		imgui.KeyA:          sdl.SCANCODE_A,
		imgui.KeyC:          sdl.SCANCODE_C,
		imgui.KeyV:          sdl.SCANCODE_V,
		imgui.KeyX:          sdl.SCANCODE_X,
		imgui.KeyY:          sdl.SCANCODE_Y,
		imgui.KeyZ:          sdl.SCANCODE_Z,
	}

	// Keyboard mapping. ImGui will use those indices to peek into the io.KeysDown[] array.
	io := imgui.CurrentIO()
	for imguiKey, nativeKey := range keys {
		io.KeyMap(imguiKey, nativeKey)
	}

	return &platformSdl2{window: window}
}

func (plat *platformSdl2) newFrame() {
	io := imgui.CurrentIO()

	// Setup display size (every frame to accommodate for window resizing)
	windowWidth, windowHeight := plat.window.GetSize()
	io.SetDisplaySize(imgui.Vec2{X: float32(windowWidth), Y: float32(windowHeight)})

	// Setup time step (we don't use SDL_GetTicks() because it is using millisecond resolution)
	frequency := sdl.GetPerformanceFrequency()
	currentTime := sdl.GetPerformanceCounter()
	if plat.time > 0 {
		io.SetDeltaTime(float32(currentTime-plat.time) / float32(frequency))
	} else {
		io.SetDeltaTime(1.0 / 60.0)
	}
	plat.time = currentTime

	// If a mouse press event came, always pass it as "mouse held this frame", so we don't miss click-release events that are shorter than 1 frame.
	x, y, state := sdl.GetMouseState()
	for i, button := range []uint32{sdl.BUTTON_LEFT, sdl.BUTTON_RIGHT, sdl.BUTTON_MIDDLE} {
		io.SetMouseButtonDown(i, plat.buttonsDown[i] || (state&sdl.Button(button)) != 0)
		plat.buttonsDown[i] = false
	}

	io.SetMousePosition(imgui.Vec2{X: float32(x), Y: float32(y)})
}

// You can read the io.WantCaptureMouse, io.WantCaptureKeyboard flags to tell if dear imgui wants to use your inputs.
// - When io.WantCaptureMouse is true, do not dispatch mouse input data to your main application.
// - When io.WantCaptureKeyboard is true, do not dispatch keyboard input data to your main application.
// Generally you may always pass all inputs to dear imgui, and hide them from your application based on those two flags.
// If you have multiple SDL events and some of them are not meant to be used by dear imgui, you may need to filter events based on their windowID field.
func (plat *platformSdl2) processEvent(event sdl.Event) bool {
	switch io := imgui.CurrentIO(); event.GetType() {
	case sdl.MOUSEWHEEL:
		wheelEvent := event.(*sdl.MouseWheelEvent)
		var deltaX, deltaY float32
		if wheelEvent.X > 0 {
			deltaX++
		} else if wheelEvent.X < 0 {
			deltaX--
		}
		if wheelEvent.Y > 0 {
			deltaY++
		} else if wheelEvent.Y < 0 {
			deltaY--
		}
		return true
	case sdl.MOUSEBUTTONDOWN:
		buttonEvent := event.(*sdl.MouseButtonEvent)
		switch buttonEvent.Button {
		case sdl.BUTTON_LEFT:
			plat.buttonsDown[0] = true
			break
		case sdl.BUTTON_RIGHT:
			plat.buttonsDown[1] = true
			break
		case sdl.BUTTON_MIDDLE:
			plat.buttonsDown[2] = true
			break
		}
		return true
	case sdl.TEXTINPUT:
		inputEvent := event.(*sdl.TextInputEvent)
		io.AddInputCharacters(string(inputEvent.Text[:]))
		return true
	case sdl.KEYDOWN:
		keyEvent := event.(*sdl.KeyboardEvent)
		io.KeyPress(int(keyEvent.Keysym.Scancode))
		modState := int(sdl.GetModState())
		io.KeyShift(modState&sdl.KMOD_LSHIFT, modState&sdl.KMOD_RSHIFT)
		io.KeyCtrl(modState&sdl.KMOD_LCTRL, modState&sdl.KMOD_RCTRL)
		io.KeyAlt(modState&sdl.KMOD_LALT, modState&sdl.KMOD_RALT)
	case sdl.KEYUP:
		keyEvent := event.(*sdl.KeyboardEvent)
		io.KeyRelease(int(keyEvent.Keysym.Scancode))
		modState := int(sdl.GetModState())
		io.KeyShift(modState&sdl.KMOD_LSHIFT, modState&sdl.KMOD_RSHIFT)
		io.KeyCtrl(modState&sdl.KMOD_LCTRL, modState&sdl.KMOD_RCTRL)
		io.KeyAlt(modState&sdl.KMOD_LALT, modState&sdl.KMOD_RALT)
		return true
	}

	return false
}

type rendererGl2 struct {
	window      *sdl.Window
	context     *sdl.GLContext
	fontTexture uint32
}

func newRendererGl2(window *sdl.Window) *rendererGl2 {
	return &rendererGl2{window: window}
}

func (rend *rendererGl2) newFrame() {
	if rend.fontTexture == 0 {
		rend.createFontsTexture()
	}

	imgui.NewFrame()
}

// OpenGL2 Render function.
// (this used to be set in io.RenderDrawListsFn and called by ImGui::Render(), but you can now call this directly from your main loop)
// Note that this implementation is little overcomplicated because we are saving/setting up/restoring every OpenGL state explicitly, in order to be able to run within any OpenGL engine that doesn't do so.
func (rend *rendererGl2) render(drawData imgui.DrawData) {
	displayWidth, displayHeight := rend.window.GetSize()
	fbWidth, fbHeight := rend.window.GLGetDrawableSize()
	drawData.ScaleClipRects(imgui.Vec2{X: float32(fbWidth) / float32(displayWidth), Y: float32(fbHeight) / float32(displayHeight)})

	// We are using the OpenGL fixed pipeline to make the example code simpler to read!
	// Setup render state: alpha-blending enabled, no face culling, no depth testing, scissor enabled, vertex/texcoord/color pointers, polygon fill.
	var lastTexture int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &lastTexture)
	var lastPolygonMode [2]int32
	gl.GetIntegerv(gl.POLYGON_MODE, &lastPolygonMode[0])
	var lastViewport [4]int32
	gl.GetIntegerv(gl.VIEWPORT, &lastViewport[0])
	var lastScissorBox [4]int32
	gl.GetIntegerv(gl.SCISSOR_BOX, &lastScissorBox[0])
	gl.PushAttrib(gl.ENABLE_BIT | gl.COLOR_BUFFER_BIT | gl.TRANSFORM_BIT)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Disable(gl.CULL_FACE)
	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.LIGHTING)
	gl.Disable(gl.COLOR_MATERIAL)
	gl.Enable(gl.SCISSOR_TEST)
	gl.EnableClientState(gl.VERTEX_ARRAY)
	gl.EnableClientState(gl.TEXTURE_COORD_ARRAY)
	gl.EnableClientState(gl.COLOR_ARRAY)
	gl.Enable(gl.TEXTURE_2D)
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

	// You may want this if using this code in an OpenGL 3+ context where shaders may be bound
	// gl.UseProgram(0)

	// Setup viewport, orthographic projection matrix
	// Our visible imgui space lies from draw_data->DisplayPos (top left) to draw_data->DisplayPos+data_data->DisplaySize (bottom right). DisplayMin is typically (0,0) for single viewport apps.
	gl.Viewport(0, 0, int32(fbWidth), int32(fbHeight))
	gl.MatrixMode(gl.PROJECTION)
	gl.PushMatrix()
	gl.LoadIdentity()
	gl.Ortho(0, float64(displayWidth), float64(displayHeight), 0, -1, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.PushMatrix()
	gl.LoadIdentity()

	vertexSize, vertexOffsetPos, vertexOffsetUv, vertexOffsetCol := imgui.VertexBufferLayout()
	indexSize := imgui.IndexBufferLayout()

	drawType := gl.UNSIGNED_SHORT
	if indexSize == 4 {
		drawType = gl.UNSIGNED_INT
	}

	// Render command lists
	for _, commandList := range drawData.CommandLists() {
		vertexBuffer, _ := commandList.VertexBuffer()
		indexBuffer, _ := commandList.IndexBuffer()
		indexBufferOffset := uintptr(indexBuffer)

		gl.VertexPointer(2, gl.FLOAT, int32(vertexSize), unsafe.Pointer(uintptr(vertexBuffer)+uintptr(vertexOffsetPos)))
		gl.TexCoordPointer(2, gl.FLOAT, int32(vertexSize), unsafe.Pointer(uintptr(vertexBuffer)+uintptr(vertexOffsetUv)))
		gl.ColorPointer(4, gl.UNSIGNED_BYTE, int32(vertexSize), unsafe.Pointer(uintptr(vertexBuffer)+uintptr(vertexOffsetCol)))

		for _, command := range commandList.Commands() {
			if command.HasUserCallback() {
				command.CallUserCallback(commandList)
			} else {
				clipRect := command.ClipRect()
				gl.Scissor(int32(clipRect.X), int32(fbHeight)-int32(clipRect.W), int32(clipRect.Z-clipRect.X), int32(clipRect.W-clipRect.Y))
				gl.BindTexture(gl.TEXTURE_2D, uint32(command.TextureID()))
				gl.DrawElements(gl.TRIANGLES, int32(command.ElementCount()), uint32(drawType), unsafe.Pointer(indexBufferOffset))
			}

			indexBufferOffset += uintptr(command.ElementCount() * indexSize)
		}
	}

	// Restore modified state
	gl.DisableClientState(gl.COLOR_ARRAY)
	gl.DisableClientState(gl.TEXTURE_COORD_ARRAY)
	gl.DisableClientState(gl.VERTEX_ARRAY)
	gl.BindTexture(gl.TEXTURE_2D, uint32(lastTexture))
	gl.MatrixMode(gl.MODELVIEW)
	gl.PopMatrix()
	gl.MatrixMode(gl.PROJECTION)
	gl.PopMatrix()
	gl.PopAttrib()
	gl.PolygonMode(gl.FRONT, uint32(lastPolygonMode[0]))
	gl.PolygonMode(gl.BACK, uint32(lastPolygonMode[1]))
	gl.Viewport(lastViewport[0], lastViewport[1], lastViewport[2], lastViewport[3])
	gl.Scissor(lastScissorBox[0], lastScissorBox[1], lastScissorBox[2], lastScissorBox[3])
}

func (rend *rendererGl2) shutdown() {
	rend.destroyFontsTexture()
}

func (rend *rendererGl2) createFontsTexture() {
	// Build texture atlas
	io := imgui.CurrentIO()
	image := io.Fonts().TextureDataRGBA32()

	// Upload texture to graphics system
	var lastTexture int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &lastTexture)
	gl.GenTextures(1, &rend.fontTexture)
	gl.BindTexture(gl.TEXTURE_2D, rend.fontTexture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.PixelStorei(gl.UNPACK_ROW_LENGTH, 0)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(image.Width), int32(image.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, image.Pixels)

	// Store our identifier
	io.Fonts().SetTextureID(imgui.TextureID(rend.fontTexture))

	// Restore state
	gl.BindTexture(gl.TEXTURE_2D, uint32(lastTexture))
}

func (rend *rendererGl2) destroyFontsTexture() {
	if rend.fontTexture != 0 {
		gl.DeleteTextures(1, &rend.fontTexture)
		imgui.CurrentIO().Fonts().SetTextureID(0)
		rend.fontTexture = 0
	}
}
