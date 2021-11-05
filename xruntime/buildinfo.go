package xruntime

import (
	"strconv"
	"time"
)

// build info
var (
	_buildVersion    string
	_buildAppVersion string
	_buildStatus     string
	_buildTag        string
	_buildUser       string
	_buildHost       string
	_buildTime       string
	buildTime        time.Time

	// _buildGitVersion      string // 构建git 版本
	// _buildGitBranch       string // 构建git branch
	// _buildGitLastCommitId string
)

func init() {
	if number, err := strconv.Atoi(_buildTime); err == nil {
		buildTime = time.Unix(int64(number), 0)
	}
}

// BuildVersion get buildVersion
func BuildVersion() string {
	return _buildVersion
}

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

// BuildTag get buildTag
func BuildTag() string {
	return _buildTag
}

// BuildTime get buildTime
func BuildTime() time.Time {
	return buildTime
}
