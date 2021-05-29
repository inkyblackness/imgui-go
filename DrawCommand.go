package imgui

// #include "wrapper/DrawCommand.h"
import "C"

// DrawCommand describes one GPU call (or a callback).
type DrawCommand uintptr

func (cmd DrawCommand) handle() C.IggDrawCmd {
	return C.IggDrawCmd(cmd)
}

// ElementCount is the number of indices (multiple of 3) to be rendered as triangles.
// Vertices are stored in the callee DrawList's VertexBuffer, indices in IndexBuffer.
func (cmd DrawCommand) ElementCount() int {
	var count C.uint
	C.iggDrawCommandGetElementCount(cmd.handle(), &count)
	return int(count)
}

// IndexOffset is the start offset in index buffer.
// Always equal to sum of ElemCount drawn so far.
func (cmd DrawCommand) IndexOffset() int {
	var count C.uint
	C.iggDrawCommandGetIndexOffset(cmd.handle(), &count)
	return int(count)
}

// VertexOffset is the start offset in vertex buffer.
// ImGuiBackendFlags_RendererHasVtxOffset: false always 0,
// otherwise may be >0 to support meshes larger than 64K vertices with 16-bit indices.
func (cmd DrawCommand) VertexOffset() int {
	var count C.uint
	C.iggDrawCommandGetVertexOffset(cmd.handle(), &count)
	return int(count)
}

// ClipRect defines the clipping rectangle (x1, y1, x2, y2).
func (cmd DrawCommand) ClipRect() (rect Vec4) {
	rectArg, rectFin := rect.wrapped()
	defer rectFin()
	C.iggDrawCommandGetClipRect(cmd.handle(), rectArg)
	return
}

// TextureID is the user-provided texture ID.
// Set by user in FontAtlas.SetTextureID() for fonts or passed to Image*() functions.
// Ignore if never using images or multiple fonts atlas.
func (cmd DrawCommand) TextureID() TextureID {
	var id C.IggTextureID
	C.iggDrawCommandGetTextureID(cmd.handle(), &id)
	return TextureID(id)
}

// HasUserCallback returns true if this handle command should be deferred.
func (cmd DrawCommand) HasUserCallback() bool {
	return C.iggDrawCommandHasUserCallback(cmd.handle()) != 0
}

// CallUserCallback calls the user callback instead of rendering the vertices.
// ClipRect and TextureID will be set normally.
func (cmd DrawCommand) CallUserCallback(list DrawList) {
	C.iggDrawCommandCallUserCallback(cmd.handle(), list.handle())
}
