package main

import (
	"math"
	"unsafe"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/inkyblackness/imgui-go"
)

type imguiGlfw3 struct {
	window           *glfw.Window
	time             float64
	mouseJustPressed [3]bool

	glslVersion            string
	fontTexture            uint32
	shaderHandle           uint32
	vertHandle             uint32
	fragHandle             uint32
	attribLocationTex      int32
	attribLocationProjMtx  int32
	attribLocationPosition int32
	attribLocationUV       int32
	attribLocationColor    int32
	vboHandle              uint32
	elementsHandle         uint32
}

func imguiGlfw3Init(window *glfw.Window) *imguiGlfw3 {
	impl := &imguiGlfw3{
		window:      window,
		glslVersion: "#version 150",
	}

	impl.installCallbacks()

	return impl
}

func (impl *imguiGlfw3) NewFrame() {
	if impl.fontTexture == 0 {
		impl.createDeviceObjects()
	}

	io := imgui.CurrentIO()
	// Setup display size (every frame to accommodate for window resizing)
	windowWidth, windowHeight := impl.window.GetSize()
	io.SetDisplaySize(imgui.Vec2{X: float32(windowWidth), Y: float32(windowHeight)})

	// Setup time step
	currentTime := glfw.GetTime()
	if impl.time > 0 {
		io.SetDeltaTime(float32(currentTime - impl.time))
	}
	impl.time = currentTime

	// Setup inputs
	if impl.window.GetAttrib(glfw.Focused) != 0 {
		x, y := impl.window.GetCursorPos()
		io.SetMousePosition(imgui.Vec2{X: float32(x), Y: float32(y)})
	} else {
		io.SetMousePosition(imgui.Vec2{X: -math.MaxFloat32, Y: -math.MaxFloat32})
	}

	for i := 0; i < len(impl.mouseJustPressed); i++ {
		down := impl.mouseJustPressed[i] || (impl.window.GetMouseButton(buttonIDByIndex[i]) == glfw.Press)
		io.SetMouseButtonDown(i, down)
		impl.mouseJustPressed[i] = false
	}

	imgui.NewFrame()
}

func (impl *imguiGlfw3) createDeviceObjects() {
	// Backup GL state
	var lastTexture int32
	var lastArrayBuffer int32
	var lastVertexArray int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &lastTexture)
	gl.GetIntegerv(gl.ARRAY_BUFFER_BINDING, &lastArrayBuffer)
	gl.GetIntegerv(gl.VERTEX_ARRAY_BINDING, &lastVertexArray)

	vertexShader := impl.glslVersion + `
uniform mat4 ProjMtx;
in vec2 Position;
in vec2 UV;
in vec4 Color;
out vec2 Frag_UV;
out vec4 Frag_Color;
void main()
{
	Frag_UV = UV;
	Frag_Color = Color;
	gl_Position = ProjMtx * vec4(Position.xy,0,1);
}
`
	fragmentShader := impl.glslVersion + `
uniform sampler2D Texture;
in vec2 Frag_UV;
in vec4 Frag_Color;
out vec4 Out_Color;
void main()
{
	Out_Color = Frag_Color * texture( Texture, Frag_UV.st).r;
}
`
	impl.shaderHandle = gl.CreateProgram()
	impl.vertHandle = gl.CreateShader(gl.VERTEX_SHADER)
	impl.fragHandle = gl.CreateShader(gl.FRAGMENT_SHADER)

	glShaderSource := func(handle uint32, source string) {
		csource, free := gl.Strs(source + "\x00")
		defer free()

		gl.ShaderSource(handle, 1, csource, nil)
	}

	glShaderSource(impl.vertHandle, vertexShader)
	glShaderSource(impl.fragHandle, fragmentShader)
	gl.CompileShader(impl.vertHandle)
	gl.CompileShader(impl.fragHandle)
	gl.AttachShader(impl.shaderHandle, impl.vertHandle)
	gl.AttachShader(impl.shaderHandle, impl.fragHandle)
	gl.LinkProgram(impl.shaderHandle)

	impl.attribLocationTex = gl.GetUniformLocation(impl.shaderHandle, gl.Str("Texture"+"\x00"))
	impl.attribLocationProjMtx = gl.GetUniformLocation(impl.shaderHandle, gl.Str("ProjMtx"+"\x00"))
	impl.attribLocationPosition = gl.GetAttribLocation(impl.shaderHandle, gl.Str("Position"+"\x00"))
	impl.attribLocationUV = gl.GetAttribLocation(impl.shaderHandle, gl.Str("UV"+"\x00"))
	impl.attribLocationColor = gl.GetAttribLocation(impl.shaderHandle, gl.Str("Color"+"\x00"))

	gl.GenBuffers(1, &impl.vboHandle)
	gl.GenBuffers(1, &impl.elementsHandle)

	impl.createFontsTexture()

	// Restore modified GL state
	gl.BindTexture(gl.TEXTURE_2D, uint32(lastTexture))
	gl.BindBuffer(gl.ARRAY_BUFFER, uint32(lastArrayBuffer))
	gl.BindVertexArray(uint32(lastVertexArray))
}

func (impl *imguiGlfw3) createFontsTexture() {
	// Build texture atlas
	io := imgui.CurrentIO()
	image := io.Fonts().TextureDataAlpha8()

	// Upload texture to graphics system
	var lastTexture int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &lastTexture)
	gl.GenTextures(1, &impl.fontTexture)
	gl.BindTexture(gl.TEXTURE_2D, impl.fontTexture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.PixelStorei(gl.UNPACK_ROW_LENGTH, 0)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RED, int32(image.Width), int32(image.Height),
		0, gl.RED, gl.UNSIGNED_BYTE, image.Pixels)

	// Store our identifier
	io.Fonts().SetTextureID(imgui.TextureID(impl.fontTexture))

	// Restore state
	gl.BindTexture(gl.TEXTURE_2D, uint32(lastTexture))
}

func (impl *imguiGlfw3) Render(drawData imgui.DrawData) {
	// Avoid rendering when minimized, scale coordinates for retina displays (screen coordinates != framebuffer coordinates)
	displayWidth, displayHeight := impl.window.GetSize()
	fbWidth, fbHeight := impl.window.GetFramebufferSize()
	if (fbWidth == 0) || (fbHeight == 0) {
		return
	}
	drawData.ScaleClipRects(imgui.Vec2{
		X: float32(fbWidth) / float32(displayWidth),
		Y: float32(fbHeight) / float32(displayHeight),
	})

	// Backup GL state
	var lastActiveTexture int32
	gl.GetIntegerv(gl.ACTIVE_TEXTURE, &lastActiveTexture)
	gl.ActiveTexture(gl.TEXTURE0)
	var lastProgram int32
	gl.GetIntegerv(gl.CURRENT_PROGRAM, &lastProgram)
	var lastTexture int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &lastTexture)
	var lastSampler int32
	gl.GetIntegerv(gl.SAMPLER_BINDING, &lastSampler)
	var lastArrayBuffer int32
	gl.GetIntegerv(gl.ARRAY_BUFFER_BINDING, &lastArrayBuffer)
	var lastElementArrayBuffer int32
	gl.GetIntegerv(gl.ELEMENT_ARRAY_BUFFER_BINDING, &lastElementArrayBuffer)
	var lastVertexArray int32
	gl.GetIntegerv(gl.VERTEX_ARRAY_BINDING, &lastVertexArray)
	var lastPolygonMode [2]int32
	gl.GetIntegerv(gl.POLYGON_MODE, &lastPolygonMode[0])
	var lastViewport [4]int32
	gl.GetIntegerv(gl.VIEWPORT, &lastViewport[0])
	var lastScissorBox [4]int32
	gl.GetIntegerv(gl.SCISSOR_BOX, &lastScissorBox[0])
	var lastBlendSrcRgb int32
	gl.GetIntegerv(gl.BLEND_SRC_RGB, &lastBlendSrcRgb)
	var lastBlendDstRgb int32
	gl.GetIntegerv(gl.BLEND_DST_RGB, &lastBlendDstRgb)
	var lastBlendSrcAlpha int32
	gl.GetIntegerv(gl.BLEND_SRC_ALPHA, &lastBlendSrcAlpha)
	var lastBlendDstAlpha int32
	gl.GetIntegerv(gl.BLEND_DST_ALPHA, &lastBlendDstAlpha)
	var lastBlendEquationRgb int32
	gl.GetIntegerv(gl.BLEND_EQUATION_RGB, &lastBlendEquationRgb)
	var lastBlendEquationAlpha int32
	gl.GetIntegerv(gl.BLEND_EQUATION_ALPHA, &lastBlendEquationAlpha)
	lastEnableBlend := gl.IsEnabled(gl.BLEND)
	lastEnableCullFace := gl.IsEnabled(gl.CULL_FACE)
	lastEnableDepthTest := gl.IsEnabled(gl.DEPTH_TEST)
	lastEnableScissorTest := gl.IsEnabled(gl.SCISSOR_TEST)

	// Setup render state: alpha-blending enabled, no face culling, no depth testing, scissor enabled, polygon fill
	gl.Enable(gl.BLEND)
	gl.BlendEquation(gl.FUNC_ADD)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Disable(gl.CULL_FACE)
	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.SCISSOR_TEST)
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

	// Setup viewport, orthographic projection matrix
	gl.Viewport(0, 0, int32(fbWidth), int32(fbHeight))
	orthoProjection := [4][4]float32{
		{2.0 / float32(displayWidth), 0.0, 0.0, 0.0},
		{0.0, 2.0 / float32(-displayHeight), 0.0, 0.0},
		{0.0, 0.0, -1.0, 0.0},
		{-1.0, 1.0, 0.0, 1.0},
	}
	gl.UseProgram(impl.shaderHandle)
	gl.Uniform1i(impl.attribLocationTex, 0)
	gl.UniformMatrix4fv(impl.attribLocationProjMtx, 1, false, &orthoProjection[0][0])
	gl.BindSampler(0, 0) // Rely on combined texture/sampler state.

	// Recreate the VAO every time
	// (This is to easily allow multiple GL contexts. VAO are not shared among GL contexts, and we don't track creation/deletion of windows so we don't have an obvious key to use to cache them.)
	var vaoHandle uint32
	gl.GenVertexArrays(1, &vaoHandle)
	gl.BindVertexArray(vaoHandle)
	gl.BindBuffer(gl.ARRAY_BUFFER, impl.vboHandle)
	gl.EnableVertexAttribArray(uint32(impl.attribLocationPosition))
	gl.EnableVertexAttribArray(uint32(impl.attribLocationUV))
	gl.EnableVertexAttribArray(uint32(impl.attribLocationColor))
	vertexSize, vertexOffsetPos, vertexOffsetUv, vertexOffsetCol := imgui.VertexBufferLayout()
	gl.VertexAttribPointer(uint32(impl.attribLocationPosition), 2, gl.FLOAT, false, int32(vertexSize), unsafe.Pointer(uintptr(vertexOffsetPos)))
	gl.VertexAttribPointer(uint32(impl.attribLocationUV), 2, gl.FLOAT, false, int32(vertexSize), unsafe.Pointer(uintptr(vertexOffsetUv)))
	gl.VertexAttribPointer(uint32(impl.attribLocationColor), 4, gl.UNSIGNED_BYTE, true, int32(vertexSize), unsafe.Pointer(uintptr(vertexOffsetCol)))
	indexSize := imgui.IndexBufferLayout()
	drawType := gl.UNSIGNED_SHORT
	if indexSize == 4 {
		drawType = gl.UNSIGNED_INT
	}

	// Draw
	for _, list := range drawData.CommandLists() {
		var indexBufferOffset uintptr

		vertexBuffer, vertexBufferSize := list.VertexBuffer()
		gl.BindBuffer(gl.ARRAY_BUFFER, impl.vboHandle)
		gl.BufferData(gl.ARRAY_BUFFER, vertexBufferSize, vertexBuffer, gl.STREAM_DRAW)

		indexBuffer, indexBufferSize := list.IndexBuffer()
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, impl.elementsHandle)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, indexBufferSize, indexBuffer, gl.STREAM_DRAW)

		for _, cmd := range list.Commands() {
			if cmd.HasUserCallback() {
				cmd.CallUserCallback(list)
			} else {
				gl.BindTexture(gl.TEXTURE_2D, uint32(cmd.TextureID()))
				clipRect := cmd.ClipRect()
				gl.Scissor(int32(clipRect.X), int32(fbHeight)-int32(clipRect.W), int32(clipRect.Z-clipRect.X), int32(clipRect.W-clipRect.Y))
				gl.DrawElements(gl.TRIANGLES, int32(cmd.ElementCount()), uint32(drawType), unsafe.Pointer(indexBufferOffset))
			}
			indexBufferOffset += uintptr(cmd.ElementCount() * indexSize)
		}
	}
	gl.DeleteVertexArrays(1, &vaoHandle)

	// Restore modified GL state
	gl.UseProgram(uint32(lastProgram))
	gl.BindTexture(gl.TEXTURE_2D, uint32(lastTexture))
	gl.BindSampler(0, uint32(lastSampler))
	gl.ActiveTexture(uint32(lastActiveTexture))
	gl.BindVertexArray(uint32(lastVertexArray))
	gl.BindBuffer(gl.ARRAY_BUFFER, uint32(lastArrayBuffer))
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, uint32(lastElementArrayBuffer))
	gl.BlendEquationSeparate(uint32(lastBlendEquationRgb), uint32(lastBlendEquationAlpha))
	gl.BlendFuncSeparate(uint32(lastBlendSrcRgb), uint32(lastBlendDstRgb), uint32(lastBlendSrcAlpha), uint32(lastBlendDstAlpha))
	if lastEnableBlend {
		gl.Enable(gl.BLEND)
	} else {
		gl.Disable(gl.BLEND)
	}
	if lastEnableCullFace {
		gl.Enable(gl.CULL_FACE)
	} else {
		gl.Disable(gl.CULL_FACE)
	}
	if lastEnableDepthTest {
		gl.Enable(gl.DEPTH_TEST)
	} else {
		gl.Disable(gl.DEPTH_TEST)
	}
	if lastEnableScissorTest {
		gl.Enable(gl.SCISSOR_TEST)
	} else {
		gl.Disable(gl.SCISSOR_TEST)
	}
	gl.PolygonMode(gl.FRONT_AND_BACK, uint32(lastPolygonMode[0]))
	gl.Viewport(lastViewport[0], lastViewport[1], lastViewport[2], lastViewport[3])
	gl.Scissor(lastScissorBox[0], lastScissorBox[1], lastScissorBox[2], lastScissorBox[3])
}

func (impl *imguiGlfw3) Shutdown() {
	impl.invalidateDeviceObjects()
}

func (impl *imguiGlfw3) invalidateDeviceObjects() {
	if impl.vboHandle != 0 {
		gl.DeleteBuffers(1, &impl.vboHandle)
	}
	impl.vboHandle = 0
	if impl.elementsHandle != 0 {
		gl.DeleteBuffers(1, &impl.elementsHandle)
	}
	impl.elementsHandle = 0

	if (impl.shaderHandle != 0) && (impl.vertHandle != 0) {
		gl.DetachShader(impl.shaderHandle, impl.vertHandle)
	}
	if impl.vertHandle != 0 {
		gl.DeleteShader(impl.vertHandle)
	}
	impl.vertHandle = 0

	if (impl.shaderHandle != 0) && (impl.fragHandle != 0) {
		gl.DetachShader(impl.shaderHandle, impl.fragHandle)
	}
	if impl.fragHandle != 0 {
		gl.DeleteShader(impl.fragHandle)
	}
	impl.fragHandle = 0

	if impl.shaderHandle != 0 {
		gl.DeleteProgram(impl.shaderHandle)
	}
	impl.shaderHandle = 0

	if impl.fontTexture != 0 {
		gl.DeleteTextures(1, &impl.fontTexture)
		imgui.CurrentIO().Fonts().SetTextureID(0)
		impl.fontTexture = 0
	}
}

func (impl *imguiGlfw3) installCallbacks() {
	impl.window.SetMouseButtonCallback(impl.mouseButtonChange)
	impl.window.SetScrollCallback(impl.mouseScrollChange)
}

var buttonIndexByID = map[glfw.MouseButton]int{
	glfw.MouseButton1: 0,
	glfw.MouseButton2: 1,
	glfw.MouseButton3: 2,
}

var buttonIDByIndex = map[int]glfw.MouseButton{
	0: glfw.MouseButton1,
	1: glfw.MouseButton2,
	2: glfw.MouseButton3,
}

func (impl *imguiGlfw3) mouseButtonChange(window *glfw.Window, rawButton glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	buttonIndex, known := buttonIndexByID[rawButton]

	if known && (action == glfw.Press) {
		impl.mouseJustPressed[buttonIndex] = true
	}
}

func (impl *imguiGlfw3) mouseScrollChange(window *glfw.Window, x, y float64) {
	io := imgui.CurrentIO()
	io.AddMouseWheelDelta(float32(x), float32(y))
}
