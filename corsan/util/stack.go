package util

import (
	"fmt"
	"regexp"
	"runtime"
)

var moduleRegexStr = `^.*\/(?P<moduleFn>.*)$`

var ModuleRegexStr = regexp.MustCompile(moduleRegexStr)

func ShortName(fullName string) string {
	names := ModuleRegexStr.SubexpNames()
	result := ModuleRegexStr.FindAllStringSubmatch(fullName, -1)
	m := map[string]string{}
	for i, n := range result[0] {
		m[names[i]] = n
	}
	return m["moduleFn"]
}

func TraceStr(depth int) string {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	fnx := runtime.FuncForPC(pc[depth])
	_, line := fnx.FileLine(pc[depth])
	module := ShortName(fnx.Name())
	return fmt.Sprintf("%s:%d", module, line)
}
