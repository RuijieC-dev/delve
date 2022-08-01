//go:build windows

package locspec

import (
	"testing"
)

func TestSubstitutePathWindows(t *testing.T) {
	// Relative paths mapping
	assertSubstitutePathEqual(t, "c:\\my\\asb\\folder\\relative\\path", SubstitutePath("relative\\path", [][2]string{{".", "c:\\my\\asb\\folder\\"}}))
	assertSubstitutePathEqual(t, "f:\\already\\abs\\path", SubstitutePath("F:\\already\\abs\\path", [][2]string{{".", "c:\\my\\asb\\folder\\"}}))
	assertSubstitutePathEqual(t, "relative\\path", SubstitutePath("C:\\my\\asb\\folder\\relative\\path", [][2]string{{"c:\\my\\asb\\folder\\", "."}}))
	assertSubstitutePathEqual(t, "f:\\another\\folder\\relative\\path", SubstitutePath("F:\\another\\folder\\relative\\path", [][2]string{{"c:\\my\\asb\\folder\\", "."}}))

	// Absolute paths mapping
	assertSubstitutePathEqual(t, "c:\\new\\mapping\\path", SubstitutePath("D:\\original\\path", [][2]string{{"d:\\original", "c:\\new\\mapping"}}))
	assertSubstitutePathEqual(t, "f:\\no\\change\\path", SubstitutePath("F:\\no\\change\\path", [][2]string{{"d:\\original", "c:\\new\\mapping"}}))

	// Mix absolute and relative mapping
	assertSubstitutePathEqual(t, "c:\\new\\mapping\\path", SubstitutePath("D:\\original\\path", [][2]string{{"d:\\original", "c:\\new\\mapping"}, {".", "c:\\my\\asb\\folder\\"}, {"c:\\my\\asb\\folder\\", "."}}))
	assertSubstitutePathEqual(t, "c:\\my\\asb\\folder\\path\\", SubstitutePath("path\\", [][2]string{{"d:\\original", "c:\\new\\mapping"}, {".", "c:\\my\\asb\\folder\\"}, {"c:\\my\\asb\\folder\\", "."}}))
	assertSubstitutePathEqual(t, "path", SubstitutePath("C:\\my\\asb\\folder\\path", [][2]string{{"d:\\original", "c:\\new\\mapping"}, {".", "c:\\my\\asb\\folder\\"}, {"c:\\my\\asb\\folder\\", "."}}))
}
