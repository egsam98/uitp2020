package file

import (
	"path/filepath"
	"runtime"
)

func Root() string {
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Dir(b)
	for i := 0; i < 2; i++ {
		root = filepath.Dir(root)
	}
	return root
}
