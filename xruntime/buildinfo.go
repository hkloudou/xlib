package xruntime

// build info
var (
	_buildVersion    string
	_buildAppVersion string
	_buildStatus     string
	_buildTag        string
	_buildUser       string
	_buildHost       string
	_buildTime       string

	// _buildGitVersion      string // 构建git 版本
	// _buildGitBranch       string // 构建git branch
	// _buildGitLastCommitId string
)

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
func BuildTime() string {
	return _buildTime
}

// // BuildGitVersion get buildGitVersion
// func BuildGitVersion() string {
// 	return _buildGitVersion
// }

// // BuildGitBranch get buildTime
// func BuildGitBranch() string {
// 	return _buildGitBranch
// }

// // BuildGitLastCommitId get buildTime
// func BuildGitLastCommitId() string {
// 	return _buildGitLastCommitId
// }
