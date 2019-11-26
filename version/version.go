package version

import (
	"github.com/coreos/go-semver/semver"
	"github.com/sirupsen/logrus"
)

var (
	// Trade's version info
	vCtlMajor, vCtlMinor, vCtlPatch int64 = 0, 1, 0
	// GitHash Value will be set during build
	GitHash = "Not provided"
	// BuildTime Value will be set during build
	BuildTime = "Not provided"
)

// TradeVer version of trade
var TradeVer = semver.Version{
	Major: vCtlMajor,
	Minor: vCtlMinor,
	Patch: vCtlPatch,
}

// LogAppInfo 打印版本信息
func LogAppInfo() {
	logrus.Infof("\nAppVersion: %s\nApiVersion: %s\nGitHash: %s\nBuildTime: %s\n\n",
		TradeVer.Major, TradeVer.Minor, GitHash, BuildTime)
}
