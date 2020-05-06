package extractor

import (
	"path/filepath"
	"strings"
)

func IsImageExt(ext string) bool {
	return strings.EqualFold(ext, ".jpg") ||
		strings.EqualFold(ext, ".jpeg") ||
		strings.EqualFold(ext, ".png") ||
		strings.EqualFold(ext, ".gif") ||
		strings.EqualFold(ext, ".svg") ||
		strings.EqualFold(ext, ".bmp") ||
		strings.EqualFold(ext, ".webp") ||
		strings.EqualFold(ext, ".bmp")
}

func IsFontExt(ext string) bool {
	return strings.EqualFold(ext, ".eot") ||
		strings.EqualFold(ext, ".otf") ||
		strings.EqualFold(ext, ".ttf") ||
		strings.EqualFold(ext, ".woff") ||
		strings.EqualFold(ext, ".woff2")
}

func IsOtherExt(ext string) bool {
	return strings.EqualFold(ext, ".so")
}

func SkipExtension(filename string) bool {
	ext := filepath.Ext(filename)
	return IsOtherExt(ext) || IsImageExt(ext) || IsFontExt(ext)
}
