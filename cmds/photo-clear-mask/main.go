//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-11-08

package main

import (
	"bytes"
	"flag"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/fsgo/photos"
	"github.com/xanygo/anygo/xcmd/xflag"
	"github.com/xanygo/anygo/ximage"
	"image/color"
)

var mediateColor = xflag.Uint8("m", 120, "用于清理的灰度中间值，>m 的会替换为白色，反之则替换为黑色")

func main() {
	flag.Parse()
	for _, fp := range flag.Args() {
		doClear(fp)
	}
}

func doClear(fp string) {
	inputImg, err := photos.Load(fp)
	if err != nil {
		log.Fatalln("load image failed:", err)
	}

	aim := ximage.ToGray(inputImg)

	outImg := image.NewGray(aim.Bounds())
	for y := inputImg.Bounds().Min.Y; y < inputImg.Bounds().Max.Y; y++ {
		for x := inputImg.Bounds().Min.X; x < inputImg.Bounds().Max.X; x++ {
			c := aim.GrayAt(x, y)
			if c.Y > *mediateColor {
				outImg.Set(x, y, color.White)
			} else {
				outImg.Set(x, y, color.Black)
			}
		}
	}
	bf := &bytes.Buffer{}
	err = png.Encode(bf, outImg)
	if err != nil {
		log.Fatalln("encode image failed:", err)
	}
	outPath := fp + ".cleared.png"
	err = os.WriteFile(outPath, bf.Bytes(), 0666)
	if err != nil {
		log.Fatalln("write file failed:", err)
	}
	log.Println("cleared:", fp, " out:", outPath)
}
