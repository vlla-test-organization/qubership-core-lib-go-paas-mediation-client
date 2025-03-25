package entity

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func WriteContext() (n string, err error) {
	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}
	res := fmt.Sprintf("%s:%d %s", filepath.Base(file), line, fnName)
	return res, nil
}
