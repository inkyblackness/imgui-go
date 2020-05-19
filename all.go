package imgui

// #cgo CPPFLAGS: -I./imgui-1.76
// #cgo CXXFLAGS: -std=c++11
// #cgo CXXFLAGS: -Wno-subobject-linkage
import "C"

// Note: imgui_freetype.cpp compilation gives these warnings, hence it's disabled in CXXFLAGS
//
// In file included from all.cpp:10:
// .\imgui-1.76/misc/freetype/imgui_freetype.cpp:294:8: warning: 'ImFontBuildSrcGlyphFT' has a field 'ImFontBuildSrcGlyphFT::Info' whose type uses the anonymous namespace [-Wsubobject-linkage]
//   294 | struct ImFontBuildSrcGlyphFT
//       |        ^~~~~~~~~~~~~~~~~~~~~
// .\imgui-1.76/misc/freetype/imgui_freetype.cpp:301:8: warning: 'ImFontBuildSrcDataFT' has a field 'ImFontBuildSrcDataFT::Font' whose type uses the anonymous namespace [-Wsubobject-linkage]
//   301 | struct ImFontBuildSrcDataFT
//       |        ^~~~~~~~~~~~~~~~~~~~
