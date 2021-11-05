package xruntime

import (
	"os"
)

var (
	//以下参数基于ENV去配置
	appMode        string
	appRegion      string
	appZone        string
	appInstance    string // 通常是实例的机器名
	appDebug       string
	appTraceIDName string
)

func initEnv() {
	appMode = os.Getenv(envAppMode)
	appRegion = os.Getenv(envAppRegion)
	appZone = os.Getenv(envAppZone)
	appInstance = os.Getenv(envAppInstance)
	if appInstance == "" {
		appInstance = HostName()
	}
	appDebug = os.Getenv(envDebug)
	if appTraceIDName == "" {
		appTraceIDName = "x-trace-id"
	}
}

// AppMode 获取应用运行的环境
func AppMode() string {
	return appMode
}

// AppRegion 获取APP运行的地区
func AppRegion() string {
	return appRegion
}

// AppZone 获取应用运行的可用区
func AppZone() string {
	return appZone
}

// AppInstance 获取应用实例，通常是实例的机器名
func AppInstance() string {
	return appInstance
}

// IsDevelopmentMode 判断是否是生产模式
func IsDevelopmentMode() bool {
	return appDebug == "true"
}

// AppTraceIDName 获取链路名称
func AppTraceIDName() string {
	return appTraceIDName
}
