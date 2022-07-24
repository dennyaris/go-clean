package config

import (
	"path/filepath"
	"runtime"
)

var (
	//current file full path from routine
	_, b, _, _ = runtime.Caller(0)

	projectPath = filepath.Join(filepath.Dir(b), "../")
)
