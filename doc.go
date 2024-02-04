// Get carbon efficiency data from WattTime.
//
// https://watttime.org/
//
// https://docs.watttime.org/
package powertime

import (
	"strconv"
	"strings"
)

// Semantic version https://semver.org/
type SemVer struct {
	Major      int
	Minor      int
	Patch      int
	PreRelease string
}

func Version() SemVer {
	return SemVer{Major: 0, Minor: 0, Patch: 0, PreRelease: "alpha.0"}
}

func VersionString() string {
	version := Version()
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(version.Major))
	builder.WriteRune('.')
	builder.WriteString(strconv.Itoa(version.Minor))
	builder.WriteRune('.')
	builder.WriteString(strconv.Itoa(version.Patch))
	if len(version.PreRelease) > 0 {
		builder.WriteRune('-')
		builder.WriteString(version.PreRelease)
	}
	return builder.String()
}
