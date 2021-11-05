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
	fmt.Printf("%s\n", "🛎Runtime info")
	fmt.Printf("%-20s : %s\n", xcolor.Green("name"), xcolor.Blue(_appName))
	fmt.Printf("%-20s : %s\n", xcolor.Green("host"), xcolor.Blue(hostName))
	fmt.Printf("%-20s : %s\n", xcolor.Green("start"), xcolor.Blue(startTime))
	fmt.Printf("%-20s : %s\n", xcolor.Green("go"), xcolor.Blue(goVersion))
	fmt.Printf("%-20s : %s\n", xcolor.Green("xlib"), xcolor.Red(xlibVersion))
	fmt.Printf("%s\n", "🛎Env info")
	fmt.Printf("%-20s : %s\n", xcolor.Green("mode"), xcolor.Blue(appMode))
	fmt.Printf("%-20s : %s\n", xcolor.Green("region"), xcolor.Blue(appRegion))
	fmt.Printf("%-20s : %s\n", xcolor.Green("zone"), xcolor.Blue(appZone))
	fmt.Printf("%-20s : %s\n", xcolor.Green("instance"), xcolor.Blue(appInstance))
	fmt.Printf("%-20s : %s\n", xcolor.Green("debug"), xcolor.Blue(appDebug))
	fmt.Printf("%-20s : %s\n", xcolor.Green("trace"), xcolor.Blue(appTraceIDName))
	fmt.Printf("%s\n", "🛎Build info")
	fmt.Printf("%-20s : %s\n", xcolor.Green("version"), xcolor.Blue(_buildAppVersion))
	fmt.Printf("%-20s : %s\n", xcolor.Green("user"), xcolor.Blue(_buildUser))
	fmt.Printf("%-20s : %s\n", xcolor.Green("host"), xcolor.Blue(_buildHost))
	fmt.Printf("%-20s : %s\n", xcolor.Green("time"), xcolor.Blue(_buildTime))
	fmt.Printf("%-20s : %s\n", xcolor.Green("status"), xcolor.Blue(_buildStatus))

	fmt.Printf("%-20s : %s\n", xcolor.Green("git"), xcolor.Blue(_buildGitVersion))
	fmt.Printf("%-20s : %s\n", xcolor.Green("branch"), xcolor.Blue(_buildGitBranch))
	fmt.Printf("%-20s : %s\n", xcolor.Green("commit"), xcolor.Blue(_buildGitLastCommitId))
}
