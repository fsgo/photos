//  Copyright(C) 2024 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2024-09-01

package layout

import (
	"image"
	"image/color"
)

func DrawBorder(sc *image.NRGBA) {
	b := sc.Bounds()
	gray := color.Gray{Y: 128}
	for _, x := range genRG(0, b.Dx(), 10) {
		sc.Set(x, 1, gray)
		sc.Set(x, b.Dy()-1, gray)
	}
	for _, y := range genRG(9, b.Dy(), 10) {
		sc.Set(0, y, gray)
		sc.Set(b.Dx()-1, y, gray)
	}
}

func genRG(start int, end int, step int) []int {
	result := make([]int, 0, (end-start)/step)
	for i := start; i < end; i += step {
		result = append(result, i)
	}
	if result[len(result)-1] != end-1 {
		result = append(result, end-1)
	}
	return result
}
