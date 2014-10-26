package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	assert.False(t, Exists("/tmp/gheogheofjojeokfej"))
	assert.True(t, Exists("/"))
}

func TestThisDir(t *testing.T) {
	assert.Equal(t, "/Users/fr/projects/artemedia/system/misc", ThisDir())
}

func TestThisDirJoin(t *testing.T) {
	assert.Equal(t, "/Users/fr/projects/artemedia/system/misc", ThisDirJoin())
	assert.Equal(t, "/Users/fr/projects/artemedia/system/misc/misc.go", ThisDirJoin("misc.go"))
	assert.Equal(t, "/Users/fr/projects/artemedia/system/types", ThisDirJoin("..", "types"))
	assert.Equal(t, "/Users/fr/projects/artemedia/system/types", ThisDirJoin("../types"))
	assert.Equal(t, "/Users/fr/projects/artemedia", ThisDirJoin("..", ".."))
}

func TestPathJoin(t *testing.T) {
	assert.Equal(t, "a/b/c", PathJoin("a/b", "c"))
}
