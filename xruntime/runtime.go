package xruntime

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/hkloudou/xlib/xtime"
)

var (
	//这些是从Runtime里去读取
	_appName    string //appname可以通过buildInfo来指定
	hostName    string
	startTime   string
	goVersion   string
	xlibVersion string
)

func initRuntime() {
	if _appName == "" {
		_appName = os.Getenv(envAppName)
		if _appName == "" {
			_appName = filepath.Base(os.Args[0])
		}
	}

	name, err := os.Hostname()
	if err != nil {
		name = "unknown"
	}
	hostName = name
	startTime = time.Now().In(xtime.TZ8).String()
	goVersion = runtime.Version()
	xlibVersion = GetPkgVersion("github.com/hkloudou/xlib")
}

// AppName gets application name.
func AppName() string {
	return _appName
}

// HostName get host name
func HostName() string {
	return hostName
}

// StartTime get start time
func StartTime() string {
	return startTime
}

// GoVersion get go version
func GoVersion() string {
	return goVersion
}

// XlibVersion get xlib version
func XlibVersion() string {
	return xlibVersion
}
