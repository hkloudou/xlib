package xruntime

import "runtime/debug"

func GetPkgVersion(name string) string {
	tmp := ""
	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, value := range info.Deps {
			if value.Path == name {
				tmp = value.Version
			}
		}
	}
	return tmp
}
