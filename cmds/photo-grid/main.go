// Copyright(C) 2024 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2024/4/6

package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/disintegration/imaging"

	"github.com/fsgo/photos"
	"github.com/fsgo/photos/layout"
)

var avatarPath = flag.String("a", "", "avatar image path")
var avatarSize = flag.String("t", layout.Avatar1IN, "avatar size, allow:\n"+strings.Join(layout.Avatar.NameDescs(), "\n")+"\n")
var outFilePath = flag.String("of", "./out.png", "canvas output path")
var outSize = flag.String("os", layout.Canvas6IN4R, "canvas size, allow:\n"+strings.Join(layout.Canvas.NameDescs(), "\n")+"\n")
var padding = flag.Int("p", 10, "padding size")

func main() {
	flag.Parse()
	if *avatarPath == "" {
		log.Fatalln("avatar image path '-a' is required")
	}
	avatar, err := photos.Load(*avatarPath)
	if err != nil {
		log.Fatalln("load avatar failed:", err)
	}

	as := layout.Avatar.MustFind(*avatarSize)
	sc := imaging.Resize(avatar, as.Width, as.Height, imaging.Lanczos)

	cs := layout.Canvas.MustFind(*outSize)
	ly := &layout.Grid{
		Width:   cs.Width,
		Height:  cs.Height,
		Padding: *padding,
	}

	bo, err := ly.Draw(sc)
	if err != nil {
		log.Fatalln("draw image failed:", err)
	}
	err = os.WriteFile(*outFilePath, bo, 0655)
	if err != nil {
		log.Fatalln("write image failed:", err)
	}
	log.Println("write image to ", *outFilePath)
}
