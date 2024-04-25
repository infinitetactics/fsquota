package fsquota

import "fmt"

// VersionMajor is the Major version
const VersionMajor = 0

// VersionMinor is the Minor version
const VersionMinor = 1

// VersionPatch is the Patch version
const VersionPatch = 9

// VersionString returns the complete version string
func VersionString() string {
	return fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
}
