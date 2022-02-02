#include "ConfiguredImGui.h"

#include "Settings.h"

void iggLoadIniSettingsFromDisk(char const *ini_filename)
{
   ImGui::LoadIniSettingsFromDisk(ini_filename);
}

void iggLoadIniSettingsFromMemory(char const *ini_data)
{
   ImGui::LoadIniSettingsFromMemory(ini_data, 0);
}

void iggSaveIniSettingsToDisk(char const *ini_filename)
{
    ImGui::SaveIniSettingsToDisk(ini_filename);
}

char const *iggSaveIniSettingsToMemory()
{
    return ImGui::SaveIniSettingsToMemory();
}