#include "ConfiguredImGui.h"

#include "Utils.h"
#include "WrapperConverter.h"

#define IMGUI_DEFINE_MATH_OPERATORS
#include "imgui_internal.h"


void iggBufferingBar(const char* label, float value, IggVec2 const *size_arg, IggVec4 const *fg_color_arg, IggVec4 const *bg_color_arg) {
    using namespace ImGui;

    ImGuiWindow* window = GetCurrentWindow();
    if (window->SkipItems)
        return;

    ImGuiContext& g = *GImGui;
    const ImGuiStyle& style = g.Style;
    const ImGuiID id = window->GetID(label);

    ImVec2 pos = window->DC.CursorPos;

    Vec2Wrapper sizeArg(size_arg);
    ImVec2 size = *sizeArg;
    size.x -= style.FramePadding.x * 2;

    const ImRect bb(pos, ImVec2(pos.x + size.x, pos.y + size.y));
    ItemSize(bb, style.FramePadding.y);
    if (!ItemAdd(bb, id))
        return;

    // Render
    const float circleStart = size.x * 0.7f;
    const float circleEnd = size.x;
    const float circleWidth = circleEnd - circleStart;

    Vec4Wrapper bgColorArg(bg_color_arg);
    const ImU32 bg_color = ImGui::GetColorU32(*bgColorArg);

    Vec4Wrapper fgColorArg(fg_color_arg);
    const ImU32 fg_color = ImGui::GetColorU32(*fgColorArg);

    window->DrawList->AddRectFilled(bb.Min, ImVec2(pos.x + circleStart, bb.Max.y), bg_color);
    window->DrawList->AddRectFilled(bb.Min, ImVec2(pos.x + circleStart*value, bb.Max.y), fg_color);

    const float t = g.Time;
    const float r = size.y / 2;
    const float speed = 1.5f;

    const float a = speed*0;
    const float b = speed*0.333f;
    const float c = speed*0.666f;

    const float o1 = (circleWidth+r) * (t+a - speed * (int)((t+a) / speed)) / speed;
    const float o2 = (circleWidth+r) * (t+b - speed * (int)((t+b) / speed)) / speed;
    const float o3 = (circleWidth+r) * (t+c - speed * (int)((t+c) / speed)) / speed;

    window->DrawList->AddCircleFilled(ImVec2(pos.x + circleEnd - o1, bb.Min.y + r), r, bg_color);
    window->DrawList->AddCircleFilled(ImVec2(pos.x + circleEnd - o2, bb.Min.y + r), r, bg_color);
    window->DrawList->AddCircleFilled(ImVec2(pos.x + circleEnd - o3, bb.Min.y + r), r, bg_color);
}

void iggLoadingIndicatorCircle(const char* label, float indicator_radius, int circle_count, float speed, IggVec4 const *fg_color_arg, IggVec4 const *bg_color_arg) {
    using namespace ImGui;

    ImGuiWindow* window = GetCurrentWindow();
    if (window->SkipItems) {
        return;
    }

    ImGuiContext& g = *GImGui;
    const ImGuiStyle& style = g.Style;
    const ImGuiID id = window->GetID(label);

    const ImVec2 pos = window->DC.CursorPos;
    const float circle_radius = indicator_radius / 10.0f;

    const ImRect bb(pos, ImVec2(pos.x + indicator_radius * 2.0f, pos.y + indicator_radius * 2.0f));
    ItemSize(bb, style.FramePadding.y);
    if (!ItemAdd(bb, id)) {
        return;
    }

    Vec4Wrapper fgColorArg(fg_color_arg);
    const ImVec4 fg_color = *fgColorArg;

    Vec4Wrapper bgColorArg(bg_color_arg);
    const ImVec4 bg_color = *bgColorArg;

    const float t = g.Time;
    const auto degree_offset = 2.0f * IM_PI / circle_count;
    for (int i = 0; i < circle_count; ++i) {
        const auto x = indicator_radius * ImSin(degree_offset * i);
        const auto y = indicator_radius * ImCos(degree_offset * i);
        auto growth = ImSin(t * speed - i * degree_offset);
        if (growth < 0) {
            growth = 0.0f;
        }

        ImVec4 color;
        color.x = fg_color.x * growth + bg_color.x * (1.0f - growth);
        color.y = fg_color.y * growth + bg_color.y * (1.0f - growth);
        color.z = fg_color.z * growth + bg_color.z * (1.0f - growth);
        color.w = 1.0f;

        window->DrawList->AddCircleFilled(ImVec2(pos.x + indicator_radius + x, pos.y + indicator_radius - y), circle_radius + growth * circle_radius, GetColorU32(color), 40);
    }
}

void iggSpinner(const char* label, float radius, int thickness, IggVec4 const *color_arg) {
    using namespace ImGui;
    // const ImU32 color = GetColorU32(ImGuiCol_ButtonHovered);

    Vec4Wrapper colorArg(color_arg);
    const ImU32 color = ImGui::GetColorU32(*colorArg);

    ImGuiWindow* window = GetCurrentWindow();
    if (window->SkipItems)
        return;

    ImGuiContext& g = *GImGui;
    const ImGuiStyle& style = g.Style;
    const ImGuiID id = window->GetID(label);

    ImVec2 pos = window->DC.CursorPos;

    const ImRect bb(pos, ImVec2(pos.x + radius * 2.0f, pos.y + radius * 2.0f));
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

IggBool iggToggleButton(const char* str_id, IggBool* selected) {
    BoolWrapper selectedArg(selected);

    ImVec2 p = ImGui::GetCursorScreenPos();
    ImDrawList* draw_list = ImGui::GetWindowDrawList();

    float height = ImGui::GetFrameHeight();
    float width = height * 1.75f;
    float radius = height * 0.50f;

    bool clicked = false;
    ImGui::InvisibleButton(str_id, ImVec2(width, height));
    if (ImGui::IsItemClicked()) {
        *selectedArg = !*selectedArg;
        clicked = true;
    }

    float t = *selectedArg ? 1.0f : 0.0f;

    ImGuiContext& g = *GImGui;
    float ANIM_TIME = 0.10f;
    if (g.LastActiveId == g.CurrentWindow->GetID(str_id)) {
        float t_anim = ImSaturate(g.LastActiveIdTimer / ANIM_TIME);
        t = *selectedArg ? (t_anim) : (1.0f - t_anim);
    }

    ImU32 col_bg;
    if (ImGui::IsItemHovered())
        col_bg = ImGui::GetColorU32(ImLerp(ImVec4(0.12f, 0.20f, 0.28f, 1.0f), ImVec4(0.07f, 0.45f, 0.37f, 1.0f), t));
    else
        col_bg = ImGui::GetColorU32(ImLerp(ImVec4(0.20f, 0.25, 0.29f, 1.0f), ImVec4(0.08f, 0.55f, 0.45f, 1.0f), t));

    draw_list->AddRectFilled(p, ImVec2(p.x + width, p.y + height), col_bg, height);
    draw_list->AddCircleFilled(ImVec2(p.x + radius + t * (width - radius * 2.0f), p.y + radius), radius - 3.0f, ImGui::GetColorU32(ImLerp(ImVec4(0.72f, 0.77f, 0.83f, 1.0f), ImVec4(0.95f, 0.96f, 0.98f, 1.0f), t)), 20);

    return clicked ? 1 : 0;
}
