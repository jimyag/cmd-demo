package version

import "runtime/debug"

var version = "UNSTABLE"

func Version() string {
	if version != "UNSTABLE" {
		return version
	}
	if v, ok := buildInfoVersion(); ok {
		return v
	}
	return version
}

func buildInfoVersion() (string, bool) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", false
	}
	return info.Main.Version, true
}
