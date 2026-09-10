package updater

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNewerVersion(t *testing.T) {
	assert.True(t, isNewerVersion("1.16.0", "1.15.2"))
	assert.True(t, isNewerVersion("1.10.0", "1.9.0"))
	assert.False(t, isNewerVersion("1.15.2", "1.16.0"))
	assert.False(t, isNewerVersion("1.16.0", "1.16.0"))
	assert.True(t, isNewerVersion("1.16", "1.15.9"))
	assert.False(t, isNewerVersion("1.16.0", "1.16"))
	assert.True(t, isNewerVersion("2.0.0", "1.99.99"))
}

func TestIsNewerVersionWithPreReleaseSuffix(t *testing.T) {
	// suffix is stripped: 1.16.0-beta → 1.16.0 for comparison

	// remote is newer
	assert.True(t, isNewerVersion("1.16.0-beta", "1.15.2"))
	// same base version, pre-release on remote → not newer
	assert.False(t, isNewerVersion("1.16.0-beta", "1.16.0"))
	// same base version, pre-release on current → not newer (fork based on same stable)
	assert.False(t, isNewerVersion("1.16.0", "1.16.0-beta"))
	// same base version, both pre-releases
	assert.False(t, isNewerVersion("1.16.0-beta", "1.16.0-beta"))
	// remote is newer base, current is pre-release of older base
	assert.True(t, isNewerVersion("1.17.0", "1.16.2-fork"))
	// current is ahead, suffix stripped → no prompt
	assert.False(t, isNewerVersion("1.16.2", "1.17.0-fork"))
	// same base, semver alpha suffix is also stripped
	assert.False(t, isNewerVersion("1.16.2", "1.16.2-alpha.1"))
	// different base, pre-release on older remote is not newer
	assert.False(t, isNewerVersion("1.16.2-alpha", "1.17.0"))
	// remote stable is newer than current pre-release base
	assert.True(t, isNewerVersion("1.16.2", "1.16.1-beta"))
}

func TestIsNewerVersionDeveloperBranch(t *testing.T) {
	// Same dev build — no update
	assert.False(t, isNewerVersion("main-dev", "main-dev"))
	// Remote is a proper release, current is a dev build — should update
	assert.True(t, isNewerVersion("1.16.0", "main-dev"))
	// Remote is a dev version, current is a release — string fallback
	assert.True(t, isNewerVersion("feature/foo-dev", "1.16.0"))
}
