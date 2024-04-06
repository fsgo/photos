// Copyright(C) 2024 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2024/4/6

package layout

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
)

type Grid struct {
	Width   int
	Height  int
	Padding int // 照片周边留白
	Gap     int // 图片间隙
}

func (g *Grid) getPadding() int {
	if g.Padding > 0 {
		return g.Padding
	}
	return 10
}

func (g *Grid) getGap() int {
	if g.Gap > 0 {
		return g.Gap
	}
	return 5
}

func (g *Grid) getWH(sw int, sh int) (int, int) {
	w := g.Width - 2*g.getPadding()
	h := g.Height - 2*g.getPadding()
	num1 := w / sw * h / sh
	num2 := h / sw * w / sh
	if num1 >= num2 {
		return g.Width, g.Height
	}
	return g.Height, g.Width
}

func (g *Grid) Draw(sc image.Image) ([]byte, error) {
	avatarWidth := sc.Bounds().Dx() + g.getGap()
	avatarHeight := sc.Bounds().Dy() + g.getGap()

	width, height := g.getWH(avatarWidth, avatarHeight)

	im := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(im, im.Bounds(), image.White, image.Point{}, draw.Src)

	startX := (width-(width/avatarWidth)*avatarWidth)/2 + g.getPadding()
	startY := (height-(height/avatarHeight)*avatarHeight)/2 + g.getPadding()

	for y := startY; y < height; y += avatarHeight {
		for x := startX; x < width; x += avatarWidth {
			if x+avatarWidth > width || y+avatarHeight > height {
				continue
			}
			draw.Draw(im, image.Rect(x, y, width, height), sc, image.Point{}, draw.Over)
		}
	}

	bf := &bytes.Buffer{}
	err := png.Encode(bf, im)
	return bf.Bytes(), err
}
