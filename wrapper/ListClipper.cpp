#include "ConfiguredImGui.h"

#include "ListClipper.h"
#include "WrapperConverter.h"

static ImGuiListClipper* importValue(IggListClipper const &in);
static void exportValue(IggListClipper &out, ImGuiListClipper const * in);

typedef TypeWrapper<ImGuiListClipper, IggListClipper> ListClipperWrapper;

IggBool iggListClipperStep(IggListClipper *clipper)
{
   ImGuiListClipper* imguiClipper = importValue(*clipper);
   IggBool returnValue = imguiClipper->Step() ? 1 : 0;
   exportValue(*clipper, imguiClipper);
   return returnValue;
}

void iggListClipperBegin(IggListClipper *clipper, int items_count, float items_height)
{
   ImGuiListClipper* imguiClipper = importValue(*clipper);
   imguiClipper->Begin(items_count, items_height);
   exportValue(*clipper, imguiClipper);
}

void iggListClipperEnd(IggListClipper *clipper)
{
   ImGuiListClipper* imguiClipper = importValue(*clipper);
   imguiClipper->End();
   ImGuiListClipper* f = (ImGuiListClipper*)clipper->imguiListClipper;
   clipper->imguiListClipper = NULL;
   exportValue(*clipper, imguiClipper);
   free(f);
}

static ImGuiListClipper* importValue(IggListClipper const &in)
{
   ImGuiListClipper* out = (ImGuiListClipper*)in.imguiListClipper;
   if (out == NULL) {
      out = new ImGuiListClipper();
      memset(out, 0, sizeof(*out));
   }
   out->DisplayStart = in.DisplayStart;
   out->DisplayEnd = in.DisplayEnd;
   return out;
}

static void exportValue(IggListClipper &out, ImGuiListClipper const * in)
{
   out.DisplayStart = in->DisplayStart;
   out.DisplayEnd = in->DisplayEnd;
   out.imguiListClipper = (void*)in;
}
