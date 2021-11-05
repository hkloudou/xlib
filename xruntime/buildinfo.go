package xruntime

// build info
var (
	_buildAppVersion      string
	_buildUser            string
	_buildHost            string
	_buildStatus          string
	_buildTime            string
	_buildGitVersion      string // 构建git 版本
	_buildGitBranch       string // 构建git branch
	_buildGitLastCommitId string
)

// BuildAppVersion get buildAppVersion
func BuildAppVersion() string {
	return _buildAppVersion
}

// BuildUser get buildUser
func BuildUser() string {
	return _buildUser
}

// BuildHost get buildHost
func BuildHost() string {
	return _buildHost
}

// BuildStatus get buildStatus
func BuildStatus() string {
	return _buildStatus
}

// BuildTime get buildTime
func BuildTime() string {
	return _buildTime
}

// BuildGitVersion get buildGitVersion
func BuildGitVersion() string {
	return _buildGitVersion
}

// BuildGitBranch get buildTime
func BuildGitBranch() string {
	return _buildGitBranch
}

// BuildGitLastCommitId get buildTime
func BuildGitLastCommitId() string {
	return _buildGitLastCommitId
}

// func PrintBuildInfo() {
// 	// fmt.Printf("%-20s : %s\n", fmt.Sprintf("xlib(%s)", xlibVersion), xcolor.Yellow("🧰🔨build info"))
// 	fmt.Println("🛎🛎🛎🛎🛎")
// 	fmt.Println(xcolor.Yellow("build info"))
// 	//xlibVersion
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("AppName"), xcolor.Blue(_appName))
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("AppHost"), xcolor.Blue(HostName()))
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("Region"), xcolor.Blue(AppRegion()))
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("Zone"), xcolor.Blue(AppZone()))
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("mmVersion"), xcolor.Red(mmVersion))
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("go ver"), xcolor.Blue(goVersion))
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("xlib"), xcolor.Blue(xlibVersion))
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("version"), xcolor.Blue(_buildAppVersion))
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("user"), xcolor.Blue(_buildUser))
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("host"), xcolor.Blue(_buildHost))
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("time"), xcolor.Blue(_buildTime))
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("status"), xcolor.Blue(_buildStatus))

// 	fmt.Printf("%-20s : %s\n", xcolor.Green("git"), xcolor.Blue(_buildGitVersion))
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("branch"), xcolor.Blue(_buildGitBranch))
// 	fmt.Printf("%-20s : %s\n", xcolor.Green("commit"), xcolor.Blue(_buildGitLastCommitId))
// }
