// Copyright(C) 2024 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2024/4/6

package photos

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
)

func Load(fp string) (image.Image, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	ext := strings.ToLower(filepath.Ext(fp))
	switch ext {
	case ".png":
		return png.Decode(file)
	case ".jpeg", ".jpg":
		return jpeg.Decode(file)
	case ".bmp":
		return bmp.Decode(file)
	case ".gif":
		return gif.Decode(file)
	default:
		return nil, fmt.Errorf("unsupported file type: %s", ext)
	}
}
