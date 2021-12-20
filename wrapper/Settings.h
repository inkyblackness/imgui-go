#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggLoadIniSettingsFromDisk(char const *ini_filename);
extern void iggLoadIniSettingsFromMemory(char const *ini_data);
extern void iggSaveIniSettingsToDisk(char const *ini_filename);
extern char const *iggSaveIniSettingsToMemory();

#ifdef __cplusplus
}
#endif