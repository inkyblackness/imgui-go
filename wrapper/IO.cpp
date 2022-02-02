#include "ConfiguredImGui.h"

#include "IO.h"
#include "WrapperConverter.h"

#include <string>

IggIO iggGetCurrentIO()
{
   return reinterpret_cast<IggIO>(&ImGui::GetIO());
}

IggBool iggWantCaptureMouse(IggIO handle)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   return io->WantCaptureMouse ? 1 : 0;
}

IggBool iggWantCaptureMouseUnlessPopupClose(IggIO handle)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   return io->WantCaptureMouseUnlessPopupClose ? 1 : 0;
}

IggBool iggWantCaptureKeyboard(IggIO handle)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   return io->WantCaptureKeyboard ? 1 : 0;
}

IggBool iggWantTextInput(IggIO handle)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   return io->WantTextInput ? 1 : 0;
}

extern float iggFramerate(IggIO handle)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   return io->Framerate;
}

extern int iggMetricsRenderVertices(IggIO handle)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   return io->MetricsRenderVertices;
}

extern int iggMetricsRenderIndices(IggIO handle)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   return io->MetricsRenderIndices;
}

extern int iggMetricsRenderWindows(IggIO handle)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   return io->MetricsRenderWindows;
}

extern int iggMetricsActiveWindows(IggIO handle)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   return io->MetricsActiveWindows;
}

extern int iggMetricsActiveAllocations(IggIO handle)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   return io->MetricsActiveAllocations;
}

extern void iggMouseDelta(IggIO handle, IggVec2 *value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   exportValue(*value, io->MouseDelta);
}

extern void iggMouseWheel(IggIO handle, float *mouseWheelH, float *mouseWheel)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   *mouseWheelH = io->MouseWheelH;
   *mouseWheel = io->MouseWheel;
}

extern void iggDisplayFrameBufferScale(IggIO handle, IggVec2 *value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   exportValue(*value, io->DisplayFramebufferScale);
}

IggFontAtlas iggIoGetFonts(IggIO handle)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   return reinterpret_cast<IggFontAtlas>(io->Fonts);
}

void iggIoSetDisplaySize(IggIO handle, IggVec2 const *value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   importValue(io->DisplaySize, *value);
}

void iggIoSetDisplayFrameBufferScale(IggIO handle, IggVec2 const *value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   importValue(io->DisplayFramebufferScale, *value);
}

void iggIoGetMousePosition(IggIO handle, IggVec2 *value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   exportValue(*value, io->MousePos);
}

void iggIoSetMousePosition(IggIO handle, IggVec2 const *value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   importValue(io->MousePos, *value);
}

void iggIoSetMouseButtonDown(IggIO handle, int index, IggBool value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   io->MouseDown[index] = value != 0;
}

void iggIoAddMouseWheelDelta(IggIO handle, float horizontal, float vertical)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   io->MouseWheelH += horizontal;
   io->MouseWheel += vertical;
}

void iggIoSetDeltaTime(IggIO handle, float value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   io->DeltaTime = value;
}

void iggIoSetFontGlobalScale(IggIO handle, float value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   io->FontGlobalScale = value;
}

void iggIoKeyPress(IggIO handle, int key)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   io.KeysDown[key] = true;
}

void iggIoKeyRelease(IggIO handle, int key)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   io.KeysDown[key] = false;
}

void iggIoKeyMap(IggIO handle, int imguiKey, int nativeKey)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   io.KeyMap[imguiKey] = nativeKey;
}

void iggIoKeyCtrl(IggIO handle, int leftCtrl, int rightCtrl)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   io.KeyCtrl = io.KeysDown[leftCtrl] || io.KeysDown[rightCtrl];
}

IggBool iggIoKeyCtrlPressed(IggIO handle)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle); 
   return io.KeyCtrl ? 1 : 0;
}

void iggIoKeyShift(IggIO handle, int leftShift, int rightShift)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   io.KeyShift = io.KeysDown[leftShift] || io.KeysDown[rightShift];
}

IggBool iggIoKeyShiftPressed(IggIO handle)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle); 
   return io.KeyShift ? 1 : 0;
}

void iggIoKeyAlt(IggIO handle, int leftAlt, int rightAlt)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   io.KeyAlt = io.KeysDown[leftAlt] || io.KeysDown[rightAlt];
}

IggBool iggIoKeyAltPressed(IggIO handle)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle); 
   return io.KeyAlt ? 1 : 0;
}

void iggIoKeySuper(IggIO handle, int leftSuper, int rightSuper)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   io.KeySuper = io.KeysDown[leftSuper] || io.KeysDown[rightSuper];
}

IggBool iggIoKeySuperPressed(IggIO handle)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle); 
   return io.KeySuper? 1 : 0;
}

void iggIoAddInputCharactersUTF8(IggIO handle, char const *utf8Chars)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   io.AddInputCharactersUTF8(utf8Chars);
}

void iggIoSetIniFilename(IggIO handle, char const *value)
{
   static std::string bufferValue;
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   bufferValue = (value != nullptr) ? value : "";
   io.IniFilename = bufferValue.empty() ? nullptr : bufferValue.c_str();
}

void iggIoSetConfigFlags(IggIO handle, int flags)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   io.ConfigFlags = flags;
}

void iggIoSetBackendFlags(IggIO handle, int flags)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   io.BackendFlags = flags;
}

int iggIoGetBackendFlags(IggIO handle)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   return io.BackendFlags;
}


void iggIoSetMouseDrawCursor(IggIO handle, IggBool show)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   io.MouseDrawCursor = show != 0;
}

extern "C" void iggIoSetClipboardText(IggIO handle, char *text);
extern "C" char *iggIoGetClipboardText(IggIO handle);

static void iggIoSetClipboardTextWrapper(void *userData, char const *text)
{
   iggIoSetClipboardText(userData, const_cast<char *>(text));
}

static char const *iggIoGetClipboardTextWrapper(void *userData)
{
   return iggIoGetClipboardText(userData);
}

void iggIoRegisterClipboardFunctions(IggIO handle)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   io.ClipboardUserData = handle;
   io.GetClipboardTextFn = iggIoGetClipboardTextWrapper;
   io.SetClipboardTextFn = iggIoSetClipboardTextWrapper;
}

void iggIoClearClipboardFunctions(IggIO handle)
{
   ImGuiIO &io = *reinterpret_cast<ImGuiIO *>(handle);
   io.GetClipboardTextFn = nullptr;
   io.SetClipboardTextFn = nullptr;
   io.ClipboardUserData = nullptr;
}
