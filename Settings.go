package imgui

// #include "wrapper/Settings.h"
import "C"

// LoadIniSettingsFromDisk loads ini settings from disk.
func LoadIniSettingsFromDisk(fileName string) {
	fileNameArg, fileNameFin := wrapString(fileName)
	defer fileNameFin()
	C.iggLoadIniSettingsFromDisk(fileNameArg)
}

// LoadIniSettingsFromMemory loads ini settings from memory.
func LoadIniSettingsFromMemory(data string) {
	dataArg, dataFin := wrapString(data)
	defer dataFin()
	C.iggLoadIniSettingsFromMemory(dataArg)
}

// SaveIniSettingsToDisk saves ini settings to disk.
func SaveIniSettingsToDisk(fileName string) {
	fileNameArg, fileNameFin := wrapString(fileName)
	defer fileNameFin()
	C.iggSaveIniSettingsToDisk(fileNameArg)
}

// SaveIniSettingsToMemory saves ini settings to memory.
func SaveIniSettingsToMemory() string {
	return C.GoString(C.iggSaveIniSettingsToMemory())
}
