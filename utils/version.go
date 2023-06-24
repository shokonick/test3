package utils

import (
	"runtime/debug"
)

var Branch = "Alpha"

// Version returns the git commit hash of the current build. Taken from https://git.vern.cc/pjals/bestofbot.
func Version() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value[:8]
			}
		}
	}
	return "unknown, please build with Go 1.13+ or use Git"
}
