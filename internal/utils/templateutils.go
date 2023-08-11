package utils

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

var TemplateFuncMap = template.FuncMap{
	// String inspection and manipulation
	"contains":    strings.Contains,
	"hasPrefix":   strings.HasPrefix,
	"hasSuffix":   strings.HasSuffix,
	"join":        strings.Join,
	"replace":     strings.Replace,
	"replaceAll":  strings.ReplaceAll,
	"split":       strings.Split,
	"splitAfter":  strings.SplitAfter,
	"splitAfterN": strings.SplitAfterN,
	"trim":        strings.Trim,
	"trimLeft":    strings.TrimLeft,
	"trimPrefix":  strings.TrimPrefix,
	"trimRight":   strings.TrimRight,
	"trimSpace":   strings.TrimSpace,
	"trimSuffix":  strings.TrimSuffix,

	// Regular expression matching
	"matchString": regexp.MatchString,
	"quoteMeta":   regexp.QuoteMeta,

	// Filepath manipulation
	"base":  filepath.Base,
	"clean": filepath.Clean,
	"dir":   filepath.Dir,

	// Basic access to reading environment variables
	"expandEnv": os.ExpandEnv,
	"getenv":    os.Getenv,

	// Naming Help
	"camel":      strcase.ToCamel,
	"lowerCamel": strcase.ToLowerCamel,
	"snake":      strcase.ToSnake,
	"kebab":      strcase.ToKebab,
	"lower":      strings.ToLower,

	//Custom
	"add":           add,
	"sub":           sub,
	"sliceContains": sliceContains,
}

func sub(a, b int) int {
	return a - b
}

func add(a, b int) int {
	return a + b
}

func sliceContains(s string, ar []string) bool {
	for _, i := range ar {
		if s == i {
			return true
		}
	}
	return false
}
