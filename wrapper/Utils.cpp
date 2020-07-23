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

IggBool iggSplitter(IggBool split_vertically, float thickness, float *size1, float *size2) {
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

// imgui combo filter v1.0, by @r-lyeh

/*  Demo: */
/*
    {
        // requisite: hints must be alphabetically sorted beforehand
        const char *hints[] = {
            "AnimGraphNode_CopyBone",
            "ce skipaa",
            "ce skipscreen",
            "ce skipsplash",
            "ce skipsplashscreen",
            "client_unit.cpp",
            "letrograd",
            "level",
            "leveler",
            "MacroCallback.cpp",
            "Miskatonic university",
            "MockAI.h",
            "MockGameplayTasks.h",
            "MovieSceneColorTrack.cpp",
            "r.maxfps",
            "r.maxsteadyfps",
            "reboot",
            "rescale",
            "reset",
            "resource",
            "restart",
            "retrocomputer",
            "retrograd",
            "return",
            "slomo 10",
            "SVisualLoggerLogsList.h",
            "The Black Knight",
        };
        static ComboFilterState s = {0};
        static char buf[128] = "type text here...";
        if( ComboFilter("my combofilter", buf, IM_ARRAYSIZE(buf), hints, IM_ARRAYSIZE(hints), s) ) {
            puts( buf );
        }
    }
*/

// bool ComboFilter__DrawPopup( IggComboFilterState *state, int START, const char **ENTRIES, int ENTRY_COUNT ) {
//     using namespace ImGui;
//     bool clicked = 0;

//     // Grab the position for the popup
//     ImVec2 pos = GetItemRectMin(); pos.y += GetItemRectSize().y;
//     ImVec2 size = ImVec2( GetItemRectSize().x-60, GetFrameHeightWithSpacing() * 4 );

//     PushStyleVar( ImGuiStyleVar_WindowRounding, 0 );

//     ImGuiWindowFlags flags = 
//         ImGuiWindowFlags_NoTitleBar          | 
//         ImGuiWindowFlags_NoResize            |
//         ImGuiWindowFlags_NoMove              |
//         ImGuiWindowFlags_HorizontalScrollbar |
//         ImGuiWindowFlags_NoSavedSettings     |
//         0; //ImGuiWindowFlags_ShowBorders;

//     SetNextWindowFocus();

//     SetNextWindowPos ( pos );
//     SetNextWindowSize( size );
//     Begin("##combo_filter", nullptr, flags );

//     PushAllowKeyboardFocus( false );

//     for( int i = 0; i < ENTRY_COUNT; i++ ) {
//         // Track if we're drawing the active index so we
//         // can scroll to it if it has changed
//         bool isIndexActive = state->activeIdx == i;

//         if( isIndexActive ) {
//             // Draw the currently 'active' item differently
//             // ( used appropriate colors for your own style )
//             PushStyleColor( ImGuiCol_Border, ImVec4( 1, 1, 0, 1 ) );
//         }

//         PushID( i );
//         if( Selectable( ENTRIES[i], isIndexActive ) ) {
//             // And item was clicked, notify the input
//             // callback so that it can modify the input text
//             state->activeIdx = i;
//             clicked = 1;
//         }
//         if( IsItemFocused() && IsKeyPressed(GetIO().KeyMap[ImGuiKey_Enter]) ) {
//             // Allow ENTER key to select current highlighted item (w/ keyboard navigation)
//             state->activeIdx = i;
//             clicked = 1;
//         }
//         PopID();

//         if( isIndexActive ) {
//             if( state->selectionChanged != 0 ) {
//                 // Make sure we bring the currently 'active' item into view.
//                 SetScrollHereY();
//                 state->selectionChanged = 0;
//             }

//             PopStyleColor(1);
//         }
//     }

//     PopAllowKeyboardFocus();
//     End();
//     PopStyleVar(1);

//     return clicked;
// }

// IggBool iggComboFilter( const char *label, char *text, int textlen, const char **hints, int num_hints, IggComboFilterState *state ) {
//     struct fuzzy {
//         static int score( const char *str1, const char *str2 ) {
//             int score = 0, consecutive = 0, maxerrors = 0;
//             while( *str1 && *str2 ) {
//                 int is_leading = (*str1 & 64) && !(str1[1] & 64);
//                 if( (*str1 & ~32) == (*str2 & ~32) ) {
//                     int had_separator = (str1[-1] <= 32);
//                     int x = had_separator || is_leading ? 10 : consecutive * 5;
//                     consecutive = 1;
//                     score += x;
//                     ++str2;
//                 } else {
//                     int x = -1, y = is_leading * -3;
//                     consecutive = 0;
//                     score += x;
//                     maxerrors += y;
//                 }
//                 ++str1;
//             }
//             return score + (maxerrors < -9 ? -9 : maxerrors);
//         }
//         static int search( const char *str, int num, const char *words[] ) {
//             int scoremax = 0;
//             int best = -1;
//             for( int i = 0; i < num; ++i ) {
//                 int score = fuzzy::score( words[i], str );
//                 int record = ( score >= scoremax );
//                 int draw = ( score == scoremax );
//                 if( record ) {
//                     scoremax = score;
//                     if( !draw ) best = i;
//                     else best = best >= 0 && strlen(words[best]) < strlen(words[i]) ? best : i;
//                 }
//             }
//             return best;
//         }
//     };
//     using namespace ImGui;
//     bool done = InputText(label, text, textlen, ImGuiInputTextFlags_AutoSelectAll | ImGuiInputTextFlags_EnterReturnsTrue );
//     bool hot = state->activeIdx >= 0 && strcmp(text, hints[state->activeIdx]);
//     if( hot ) {
//         int new_idx = fuzzy::search( text, num_hints, hints );
//         int idx = new_idx >= 0 ? new_idx : state->activeIdx;
//         state->selectionChanged = state->activeIdx != idx ? 1 : 0;
//         state->activeIdx = idx;
//         if( done || ComboFilter__DrawPopup( state, idx, hints, num_hints ) ) {
//             int i = state->activeIdx;
//             if( i >= 0 ) {
//                 strcpy(text, hints[i]);
//                 done = true;
//             }
//         }
//     }
//     return done;
// }

// IggBool iggComboFilter(const char *label, char *buffer, int bufferlen, const char **hints, int num_hints, IggComboFilterState *state, int callback_key) {
//     using namespace ImGui;
//     ImGuiComboFlags flags = 0;

//     state->selectionChanged = 0;

//     // Always consume the SetNextWindowSizeConstraint() call in our early return paths
//     ImGuiContext& g = *GImGui;

//     ImGuiWindow* window = GetCurrentWindow();
//     if (window->SkipItems)
//         return false;
    
//     const ImGuiID id = window->GetID(label);
//     bool popup_open = IsPopupOpen(id, ImGuiPopupFlags_None);
//     bool popupNeedBeOpen = strcmp(buffer, hints[state->activeIdx]);
//     bool popupJustOpened = false;
    
//     IM_ASSERT((flags & (ImGuiComboFlags_NoArrowButton | ImGuiComboFlags_NoPreview)) != (ImGuiComboFlags_NoArrowButton | ImGuiComboFlags_NoPreview)); // Can't use both flags together

//     const ImGuiStyle& style = g.Style;
    
//     const float arrow_size = (flags & ImGuiComboFlags_NoArrowButton) ? 0.0f : GetFrameHeight();
//     const ImVec2 label_size = CalcTextSize(label, NULL, true);
//     const float expected_w = CalcItemWidth();
//     const float w = (flags & ImGuiComboFlags_NoPreview) ? arrow_size : expected_w;
//     const ImRect frame_bb(window->DC.CursorPos, ImVec2(window->DC.CursorPos.x + w, window->DC.CursorPos.y + label_size.y + style.FramePadding.y*2.0f));
//     const ImRect total_bb(frame_bb.Min, ImVec2((label_size.x > 0.0f ? style.ItemInnerSpacing.x + label_size.x : 0.0f) + frame_bb.Max.x, frame_bb.Max.y));
//     const float value_x2 = ImMax(frame_bb.Min.x, frame_bb.Max.x - arrow_size);
//     ItemSize(total_bb, style.FramePadding.y);
//     if (!ItemAdd(total_bb, id, &frame_bb))
//         return false;
    
    
//     bool hovered, held;
//     bool pressed = ButtonBehavior(frame_bb, id, &hovered, &held);
    
//     if(!popup_open) {
//         const ImU32 frame_col = GetColorU32(hovered ? ImGuiCol_FrameBgHovered : ImGuiCol_FrameBg);
//         RenderNavHighlight(frame_bb, id);
//         if (!(flags & ImGuiComboFlags_NoPreview))
//             window->DrawList->AddRectFilled(frame_bb.Min, ImVec2(value_x2, frame_bb.Max.y), frame_col, style.FrameRounding, (flags & ImGuiComboFlags_NoArrowButton) ? ImDrawCornerFlags_All : ImDrawCornerFlags_Left);
//     }
//     if (!(flags & ImGuiComboFlags_NoArrowButton))
//     {
//         ImU32 bg_col = GetColorU32((popup_open || hovered) ? ImGuiCol_ButtonHovered : ImGuiCol_Button);
//         ImU32 text_col = GetColorU32(ImGuiCol_Text);
//         window->DrawList->AddRectFilled(ImVec2(value_x2, frame_bb.Min.y), frame_bb.Max, bg_col, style.FrameRounding, (w <= arrow_size) ? ImDrawCornerFlags_All : ImDrawCornerFlags_Right);
//         if (value_x2 + arrow_size - style.FramePadding.x <= frame_bb.Max.x)
//             RenderArrow(window->DrawList, ImVec2(value_x2 + style.FramePadding.y, frame_bb.Min.y + style.FramePadding.y), text_col, ImGuiDir_Down, 1.0f);
//     }
//     if(!popup_open) {

//         RenderFrameBorder(frame_bb.Min, frame_bb.Max, style.FrameRounding);
//         if (buffer != NULL && !(flags & ImGuiComboFlags_NoPreview))
            
//             RenderTextClipped(ImVec2(frame_bb.Min.x + style.FramePadding.x, frame_bb.Min.y + style.FramePadding.y), ImVec2(value_x2, frame_bb.Max.y), buffer, NULL, NULL, ImVec2(0.0f,0.0f));

//         if ((pressed || g.NavActivateId == id || popupNeedBeOpen) && !popup_open)
//         {
//             if (window->DC.NavLayerCurrent == 0)
//                 window->NavLastIds[0] = id;
//             OpenPopupEx(id);
//             popup_open = true;
//             popupJustOpened = true;
//         }
//     }
    
//     if (label_size.x > 0)
//     RenderText(ImVec2(frame_bb.Max.x + style.ItemInnerSpacing.x, frame_bb.Min.y + style.FramePadding.y), label);
    
//     if (!popup_open) {
//         return false;
//     }
    
//     const float totalWMinusArrow = w - arrow_size;
//     struct ImGuiSizeCallbackWrapper {
//         static void sizeCallback(ImGuiSizeCallbackData* data)
//         {
//             float* totalWMinusArrow = (float*)(data->UserData);
//             data->DesiredSize = ImVec2(*totalWMinusArrow, 200.f);
//         }
//     };
//     SetNextWindowSizeConstraints(ImVec2(0 ,0), ImVec2(totalWMinusArrow, 150.f), ImGuiSizeCallbackWrapper::sizeCallback, (void*)&totalWMinusArrow);

//     char name[16];
//     ImFormatString(name, IM_ARRAYSIZE(name), "##Combo_%02d", g.BeginPopupStack.Size); // Recycle windows based on depth

//     // Peak into expected window size so we can position it
//     if (ImGuiWindow* popup_window = FindWindowByName(name))
//         if (popup_window->WasActive)
//         {
//             ImVec2 size_expected = CalcWindowExpectedSize(popup_window);
//             if (flags & ImGuiComboFlags_PopupAlignLeft)
//                 popup_window->AutoPosLastDirection = ImGuiDir_Left;
//             ImRect r_outer = GetWindowAllowedExtentRect(popup_window);
//             ImVec2 pos = FindBestWindowPosForPopupEx(frame_bb.GetBL(), size_expected, &popup_window->AutoPosLastDirection, r_outer, frame_bb, ImGuiPopupPositionPolicy_ComboBox);
            
//             pos.y -= label_size.y + style.FramePadding.y*2.0f;
            
//             SetNextWindowPos(pos);
//         }

//     // Horizontally align ourselves with the framed text
//     ImGuiWindowFlags window_flags = ImGuiWindowFlags_AlwaysAutoResize | ImGuiWindowFlags_Popup | ImGuiWindowFlags_NoTitleBar | ImGuiWindowFlags_NoResize | ImGuiWindowFlags_NoSavedSettings;
// //    PushStyleVar(ImGuiStyleVar_WindowPadding, ImVec2(style.FramePadding.x, style.WindowPadding.y));
//     bool ret = Begin(name, NULL, window_flags);

//     ImGui::PushItemWidth(ImGui::GetWindowWidth());
//     ImGui::SetCursorPos(ImVec2(0.f, window->DC.CurrLineTextBaseOffset));
//     if(popupJustOpened) {
//         ImGui::SetKeyboardFocusHere(0);
//     }
//     bool done = InputTextEx("", NULL, buffer, bufferlen, ImVec2(0, 0), ImGuiInputTextFlags_CallbackResize | ImGuiInputTextFlags_AutoSelectAll | ImGuiInputTextFlags_EnterReturnsTrue, iggInputTextCallbackWrapper, reinterpret_cast<void *>(callback_key));
//     ImGui::PopItemWidth();

//     if(state->activeIdx < 0) {
//         IM_ASSERT(false); //Undefined behaviour
//         return false;
//     }


//     if (!ret)
//     {
//         ImGui::EndChild();
//         ImGui::PopItemWidth();
//         EndPopup();
//         IM_ASSERT(0);   // This should never happen as we tested for IsPopupOpen() above
//         return false;
//     }


//     ImGuiWindowFlags window_flags2 =  0; //ImGuiWindowFlags_HorizontalScrollbar
//     ImGui::BeginChild("ChildL", ImVec2(ImGui::GetContentRegionAvail().x, ImGui::GetContentRegionAvail().y), false, window_flags2);




//     struct fuzzy {
//         static int score( const char *str1, const char *str2 ) {
//             int score = 0, consecutive = 0, maxerrors = 0;
//             while( *str1 && *str2 ) {
//                 int is_leading = (*str1 & 64) && !(str1[1] & 64);
//                 if( (*str1 & ~32) == (*str2 & ~32) ) {
//                     int had_separator = (str1[-1] <= 32);
//                     int x = had_separator || is_leading ? 10 : consecutive * 5;
//                     consecutive = 1;
//                     score += x;
//                     ++str2;
//                 } else {
//                     int x = -1, y = is_leading * -3;
//                     consecutive = 0;
//                     score += x;
//                     maxerrors += y;
//                 }
//                 ++str1;
//             }
//             return score + (maxerrors < -9 ? -9 : maxerrors);
//         }
//         static int search( const char *str, int num, const char *words[] ) {
//             int scoremax = 0;
//             int best = -1;
//             for( int i = 0; i < num; ++i ) {
//                 int score = fuzzy::score( words[i], str );
//                 int record = ( score >= scoremax );
//                 int draw = ( score == scoremax );
//                 if( record ) {
//                     scoremax = score;
//                     if( !draw ) best = i;
//                     else best = best >= 0 && strlen(words[best]) < strlen(words[i]) ? best : i;
//                 }
//             }
//             return best;
//         }
//     };
    
//     int new_idx = fuzzy::search( buffer, num_hints, hints );
//     int idx = new_idx >= 0 ? new_idx : state->activeIdx;
//     state->selectionChanged = state->activeIdx != idx ? 1 : 0;
//     bool selectionChangedLocal = state->selectionChanged;
//     state->activeIdx = idx;
    
//     if(done) {
//         CloseCurrentPopup();
//     }
//     for (int n = 0; n < num_hints; n++) {;
//         bool is_selected = n == state->activeIdx;
//         if (is_selected && (IsWindowAppearing() || selectionChangedLocal)) {
//              SetScrollHereY();
// //            ImGui::SetItemDefaultFocus();
//         }
//         if (ImGui::Selectable(hints[n], is_selected)) {
//             state->selectionChanged = state->activeIdx != n ? 1 : 0;
//             state->activeIdx = n;
//             strcpy(buffer, hints[n]);
//             CloseCurrentPopup();
//         }
//     }
//     ImGui::EndChild();
//     EndPopup();

//     return state->selectionChanged != 0 && !strcmp(hints[state->activeIdx], buffer);
// }


static int iggInputTextCallbackWrapper(ImGuiInputTextCallbackData *data);

// Create text input in place of another active widget (e.g. used when doing a CTRL+Click on drag/slider widgets)
// FIXME: Facilitate using this in variety of other situations.
bool TempInputText__(const ImRect& bb, ImGuiID id, const char* label, char* buf, int buf_size, void* user_data) {
    using namespace ImGui;
    // On the first frame, g.TempInputTextId == 0, then on subsequent frames it becomes == id.
    // We clear ActiveID on the first frame to allow the InputText() taking it back.
    ImGuiContext& g = *GImGui;
    const bool init = (g.TempInputId != id);
    if (init)
        ClearActiveID();

    g.CurrentWindow->DC.CursorPos = bb.Min;
    bool value_changed = InputTextEx(label, NULL, buf, buf_size, bb.GetSize(), ImGuiInputTextFlags_CallbackResize|ImGuiInputTextFlags_EnterReturnsTrue, iggInputTextCallbackWrapper, user_data) || IsItemDeactivatedAfterEdit();
    if (init)
    {
        // First frame we started displaying the InputText widget, we expect it to take the active id.
        IM_ASSERT(g.ActiveId == id);
        g.TempInputId = g.ActiveId;
    }
    return value_changed;
}

IggBool iggSelectableInput(const char *label, char *text, int text_size, int callback_key) {
    using namespace ImGui;

    ImGuiContext& g = *GImGui;
    ImGuiWindow* window = g.CurrentWindow;
    ImVec2 pos_before = window->DC.CursorPos;

    PushID(label);
    PushStyleVar(ImGuiStyleVar_ItemSpacing, ImVec2(g.Style.ItemSpacing.x, g.Style.FramePadding.y * 2.0f));
    bool selectable_ret = Selectable("##Selectable", false, ImGuiSelectableFlags_AllowDoubleClick | ImGuiSelectableFlags_AllowItemOverlap);
    PopStyleVar();

    ImGuiID id = window->GetID("##Input");
    bool temp_input_is_active = TempInputIsActive(id);
    bool temp_input_start = selectable_ret ? IsMouseDoubleClicked(0) : false;

    if (temp_input_start) SetActiveID(id, window);

    bool value_changed = false;
    if (temp_input_is_active || temp_input_start) {
        ImVec2 pos_after = window->DC.CursorPos;
        window->DC.CursorPos = pos_before;
        value_changed = TempInputText__(window->DC.LastItemRect, id, "##Input", text, text_size, reinterpret_cast<void *>(callback_key));
        window->DC.CursorPos = pos_after;
    } else {
        window->DrawList->AddText(pos_before, GetColorU32(ImGuiCol_Text), text);
    }

    PopID();
    return value_changed;
}

