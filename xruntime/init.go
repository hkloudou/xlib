package xruntime

import (
	"fmt"

	"github.com/hkloudou/xlib/xcolor"
)

func init() {
	initRuntime() //读取运行环境
	initEnv()     //读取环境变量
}

func PrintInfo() {
	// fmt.Printf("%-20s : %s\n", xcolor.Green("MM"), xcolor.Yellow("I am MM,do you like me"))
	fmt.Printf("%s\n", xcolor.Yellow("🛎  Runtime info"))
	fmt.Printf("%-20s : %s\n", xcolor.Green("name"), xcolor.Blue(_appName))
	fmt.Printf("%-20s : %s\n", xcolor.Green("host"), xcolor.Blue(hostName))
	fmt.Printf("%-20s : %s\n", xcolor.Green("start"), xcolor.Blue(startTime))
	fmt.Printf("%-20s : %s\n", xcolor.Green("go"), xcolor.Red(goVersion))
	fmt.Printf("%-20s : %s\n", xcolor.Green("xlib"), xcolor.Red(xlibVersion))
	fmt.Printf("%s\n", xcolor.Yellow("🛎  Env info"))
	fmt.Printf("%-20s : %s\n", xcolor.Green("mode"), xcolor.Blue(appMode))
	fmt.Printf("%-20s : %s\n", xcolor.Green("region"), xcolor.Blue(appRegion))
	fmt.Printf("%-20s : %s\n", xcolor.Green("zone"), xcolor.Blue(appZone))
	fmt.Printf("%-20s : %s\n", xcolor.Green("instance"), xcolor.Blue(appInstance))
	fmt.Printf("%-20s : %s\n", xcolor.Green("debug"), xcolor.Red(appDebug))
	fmt.Printf("%-20s : %s\n", xcolor.Green("trace"), xcolor.Red(appTraceIDName))
	fmt.Printf("%s\n", xcolor.Yellow("🛎  Build info"))
	fmt.Printf("%-20s : %s\n", xcolor.Green("version"), xcolor.Blue(_buildVersion))
	fmt.Printf("%-20s : %s\n", xcolor.Green("app"), xcolor.Blue(_buildAppVersion))
	fmt.Printf("%-20s : %s\n", xcolor.Green("tag"), xcolor.Blue(_buildTag))
	fmt.Printf("%-20s : %s\n", xcolor.Green("user"), xcolor.Blue(_buildUser))
	fmt.Printf("%-20s : %s\n", xcolor.Green("host"), xcolor.Blue(_buildHost))
	fmt.Printf("%-20s : %s\n", xcolor.Green("time"), xcolor.Blue(fmt.Sprintf("%d", buildTime.Unix())))
	fmt.Printf("%-20s : %s\n", xcolor.Green("time"), xcolor.Blue(buildTime.Local().String()))
	fmt.Printf("%-20s : %s\n", xcolor.Green("status"), xcolor.Blue(_buildStatus))

	// fmt.Printf("%-20s : %s\n", xcolor.Green("git"), xcolor.Red(_buildGitVersion))
	// fmt.Printf("%-20s : %s\n", xcolor.Green("branch"), xcolor.Red(_buildGitBranch))
	// fmt.Printf("%-20s : %s\n", xcolor.Green("commit"), xcolor.Red(_buildGitLastCommitId))
}
