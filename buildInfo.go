package xlib

import (
	"fmt"

	"github.com/hkloudou/xlib/xcolor"
)

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

func PrintBuildInfo() {
	fmt.Printf("%s\n", xcolor.Yellow("🧰🔨build info"))
	fmt.Printf("%-20s : %s\n", xcolor.Green("buildAppVersion"), xcolor.Blue(_buildAppVersion))
	fmt.Printf("%-20s : %s\n", xcolor.Green("BuildUser"), xcolor.Blue(_buildUser))
	fmt.Printf("%-20s : %s\n", xcolor.Green("BuildHost"), xcolor.Blue(_buildHost))
	fmt.Printf("%-20s : %s\n", xcolor.Green("BuildTime"), xcolor.Blue(_buildTime))
	fmt.Printf("%-20s : %s\n", xcolor.Green("BuildStatus"), xcolor.Blue(_buildStatus))
}
