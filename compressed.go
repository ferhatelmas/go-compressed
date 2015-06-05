// Package compressed provides a collection of compressed file extensions
// and a utility function to check if given path is a compressed one
package compressed

import (
	"path"
	"strings"
)

var extensions = []string{
	"7z",
	"apk",
	"bz2",
	"cab",
	"dmg",
	"egg",
	"epub",
	"gz",
	"jar",
	"lz",
	"lzma",
	"lzo",
	"pea",
	"rar",
	"rz",
	"s7z",
	"tbz2",
	"tgz",
	"tlz",
	"war",
	"whl",
	"xpi",
	"xz",
	"z",
	"zip",
	"zipx",
}

// Extensions is the extensions for different compressed file types
var Extensions map[string]bool

func init() {
	Extensions = map[string]bool{}
	for _, ext := range extensions {
		Extensions[ext] = true
	}
}

// Is checks if a path is a compressed file
func Is(p string) bool {
	ext := strings.TrimLeft(path.Ext(p), ".")
	return Extensions[strings.ToLower(ext)]
}
