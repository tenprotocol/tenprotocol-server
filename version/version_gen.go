// Code generated by winch. DO NOT EDIT.
package version

import (
	"fmt"
	"runtime"
)

const (
	Name        = "tenprotocol-server"
	Description = "Reference implementation of the Threat Exposure Notification Protocol (10p)."
	ReleaseName = "healthy hyena"
	Version     = "0.1.0"
	Prerelease  = ""
)

// String returns the complete version string, including prerelease
func String() string {
	s := fmt.Sprintf("%s %s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
	if Prerelease != "" {
		return fmt.Sprintf("%s-%s %s", Version, Prerelease, s)
	}
	return fmt.Sprintf("%s %s", Version, s)
}

