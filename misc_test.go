package misc

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	assert.False(t, Exists("/tmp/gheogheofjojeokfej"))
	assert.True(t, Exists("/"))
}

func thisDir() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		Fatal("runtime.Caller() was not ok")
	}
	return filepath.Dir(filename)
}

func TestThisDir(t *testing.T) {
	assert.Equal(t, thisDir(), ThisDir())
}

func TestThisDirJoin(t *testing.T) {
	assert.Equal(t, thisDir(), ThisDirJoin())
	assert.Equal(t, filepath.Join(thisDir(), "misc.go"), ThisDirJoin("misc.go"))
	assert.Equal(t, filepath.Join(thisDir(), "..", "types"), ThisDirJoin("..", "types"))
	assert.Equal(t, filepath.Join(thisDir(), "../types"), ThisDirJoin("../types"))
	assert.Equal(t, filepath.Join(thisDir(), "..", ".."), ThisDirJoin("..", ".."))
}

func TestPathJoin(t *testing.T) {
	assert.Equal(t, filepath.Join("a/b", "c"), PathJoin("a/b", "c"))
}
