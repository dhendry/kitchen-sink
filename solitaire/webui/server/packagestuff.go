package server

import (
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	pc, file, line, _ := runtime.Caller(0)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	log.WithFields(logrus.Fields{
		"file": file,
		"fileName": fileName,
		"line": line,
		"packageName": packageName,
		"parts": parts,
	}).Info()
	//logrus.WithField("foo", "bar")
}
