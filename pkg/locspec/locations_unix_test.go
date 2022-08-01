//go:build !windows

package locspec

import (
	"testing"
)

func TestSubstitutePathUnix(t *testing.T) {
	// Relative paths mapping
	assertSubstitutePathEqual(t, "/my/asb/folder/relative/path", SubstitutePath("relative/path", [][2]string{{".", "/my/asb/folder/"}}))
	assertSubstitutePathEqual(t, "/already/abs/path", SubstitutePath("/already/abs/path", [][2]string{{".", "/my/asb/folder/"}}))
	assertSubstitutePathEqual(t, "relative/path", SubstitutePath("/my/asb/folder/relative/path", [][2]string{{"/my/asb/folder/", "."}}))
	assertSubstitutePathEqual(t, "/another/folder/relative/path", SubstitutePath("/another/folder/relative/path", [][2]string{{"/my/asb/folder/", "."}}))

	// Absolute paths mapping
	assertSubstitutePathEqual(t, "/new/mapping/path", SubstitutePath("/original/path", [][2]string{{"/original", "/new/mapping"}}))
	assertSubstitutePathEqual(t, "/no/change/path", SubstitutePath("/no/change/path", [][2]string{{"/original", "/new/mapping"}}))

	// Mix absolute and relative mapping
	assertSubstitutePathEqual(t, "/new/mapping/path", SubstitutePath("/original/path", [][2]string{{"/original", "/new/mapping"}, {".", "/my/asb/folder/"}, {"/my/asb/folder/", "."}}))
	assertSubstitutePathEqual(t, "/my/asb/folder/path", SubstitutePath("path", [][2]string{{"/original", "/new/mapping"}, {".", "/my/asb/folder/"}, {"/my/asb/folder/", "."}}))
	assertSubstitutePathEqual(t, "path", SubstitutePath("/my/asb/folder/path", [][2]string{{"/original", "/new/mapping"}, {".", "/my/asb/folder/"}, {"/my/asb/folder/", "."}}))
}
