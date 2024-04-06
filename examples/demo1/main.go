// Copyright(C) 2024 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2024/4/6

package main

import (
	"bytes"
	"image/png"
	"log"
	"os"

	"github.com/disintegration/imaging"

	"github.com/fsgo/photos/layout"
)

func main() {
	bf, _ := os.ReadFile("../mg.png")
	mg, err := png.Decode(bytes.NewBuffer(bf))
	if err != nil {
		log.Fatalln(err)
	}
	sc := imaging.Resize(mg, 295, 413, imaging.Lanczos)
	ly := &layout.Grid{
		Width:  1204,
		Height: 1794,
	}
	bf2, err := ly.Draw(sc)
	if err != nil {
		log.Fatalln(err)
	}
	os.WriteFile("tmp.png", bf2, 0644)
}
