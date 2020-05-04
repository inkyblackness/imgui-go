#include "imgui.h"

#include "Utils.h"

#define IMGUI_DEFINE_MATH_OPERATORS
#include "imgui_internal.h"

void iggSpinner(const char* label, float radius, int thickness) {    
    using namespace ImGui;
    const ImU32 color = GetColorU32(ImGuiCol_ButtonHovered);

    ImGuiWindow* window = GetCurrentWindow();
    if (window->SkipItems)
        return;

    ImGuiContext& g = *GImGui;
    const ImGuiStyle& style = g.Style;
    const ImGuiID id = window->GetID(label);

    ImVec2 pos = window->DC.CursorPos;
    ImVec2 size((radius )*2, (radius + style.FramePadding.y)*2);

    const ImRect bb(pos, ImVec2(pos.x + size.x, pos.y + size.y));
    ItemSize(bb, style.FramePadding.y);
    if (!ItemAdd(bb, id))
        return;

    // Render
    window->DrawList->PathClear();

    int num_segments = 30;
    int start = fabs(ImSin(g.Time*1.8f)*(num_segments-5));

    const float a_min = IM_PI*2.0f * ((float)start) / (float)num_segments;
    const float a_max = IM_PI*2.0f * ((float)num_segments-3) / (float)num_segments;

    const ImVec2 centre = ImVec2(pos.x+radius, pos.y+radius+style.FramePadding.y);

    for (int i = 0; i < num_segments; i++) {
        const float a = a_min + ((float)i / (float)num_segments) * (a_max - a_min);
        window->DrawList->PathLineTo(ImVec2(centre.x + ImCos(a+g.Time*8) * radius,
                                            centre.y + ImSin(a+g.Time*8) * radius));
    }

    window->DrawList->PathStroke(color, false, thickness);
}

IggBool iggSplitter(IggBool split_vertically, float thickness, float *size1, float *size2)
{
   using namespace ImGui;
   ImGuiContext& g = *GImGui;
   ImGuiWindow* window = g.CurrentWindow;
   ImGuiID id = window->GetID("##Splitter");
   ImRect bb;
   bb.Min = window->DC.CursorPos + (split_vertically != 0 ? ImVec2(*size1, 0.0f) : ImVec2(0.0f, *size1));
   bb.Max = bb.Min + CalcItemSize(split_vertically != 0 ? ImVec2(thickness, -1.f) : ImVec2(-1.f, thickness), 0.0f, 0.0f);
//IMGUI_API bool          SplitterBehavior(const ImRect& bb, ImGuiID id, ImGuiAxis axis, float* size1, float* size2, float min_size1, float min_size2, float hover_extend = 0.0f, float hover_visibility_delay = 0.0f);
    
   return SplitterBehavior(bb, id, split_vertically != 0 ? ImGuiAxis_X : ImGuiAxis_Y, size1, size2, 8.f, 8.f);
}