// Copyright(C) 2024 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2024/4/6

package layout

import "fmt"

// Size 基于 pdi=300 所定义的
type Size struct {
	Name   string
	Width  int
	Height int
	Desc   string
}

type Sizes []Size

func (sz Sizes) Find(name string) (Size, bool) {
	for _, item := range sz {
		if item.Name == name {
			return item, true
		}
	}
	return Size{}, false
}

func (sz Sizes) MustFind(name string) Size {
	for _, item := range sz {
		if item.Name == name {
			return item
		}
	}
	panic(fmt.Sprintf("cannot find size of %s", name))
}

func (sz Sizes) NameDescs() []string {
	names := make([]string, 0, len(sz))
	for _, item := range sz {
		desc := item.Name + " :\t" + item.Desc
		names = append(names, desc)
	}
	return names
}

const (
	CanvasA0 = "A0"
	CanvasA1 = "A1"
	CanvasA2 = "A2"
	CanvasA3 = "A3"
	CanvasA4 = "A4"
	CanvasA5 = "A5"
	CanvasA6 = "A6"

	Canvas6IN4R  = "6in-4R"
	Canvas5IN3R  = "5in-3R"
	Canvas7IN5R  = "7in-5R"
	Canvas8IN6R  = "8in-6R"
	Canvas10IN8R = "10in-8R"
)

// Canvas 画布
var Canvas = Sizes{
	{
		Name:   CanvasA0,
		Width:  9933,
		Height: 14043,
	},
	{
		Name:   CanvasA1,
		Width:  7016,
		Height: 9933,
	},
	{
		Name:   CanvasA2,
		Width:  4960,
		Height: 7016,
		Desc:   "42.00cm * 59.40",
	},
	{
		Name:   CanvasA3,
		Width:  3508,
		Height: 4960,
		Desc:   "29.70cm * 42.00cm",
	},
	{
		Name:   CanvasA4,
		Width:  2480,
		Height: 3508,
		Desc:   "21.00cm * 29.70cm",
	},
	{
		Name:   CanvasA5,
		Width:  1748,
		Height: 2480,
		Desc:   "14.80cm * 21.00cm",
	},
	{
		Name:   CanvasA6,
		Width:  1240,
		Height: 1748,
		Desc:   "10.50cm * 14.80cm",
	},
	{
		Name:   Canvas6IN4R,
		Width:  1204,
		Height: 1794,
		Desc:   "10.20cm * 15.20cm",
	},
	{
		Name:   Canvas5IN3R,
		Width:  1499,
		Height: 1050,
		Desc:   "12.70cm * 8.90cm",
	},
	{
		Name:   Canvas7IN5R,
		Width:  2101,
		Height: 1499,
		Desc:   "17.80cm * 12.70cm",
	},
	{
		Name:   Canvas8IN6R,
		Width:  2396,
		Height: 1794,
		Desc:   "20.30cm * 15.20cm",
	},
	{
		Name:   Canvas10IN8R,
		Width:  2998,
		Height: 2396,
		Desc:   "25.40cm * 20.30cm",
	},
}

const (
	Avatar1IN      = "1in"
	Avatar1INBig   = "1in-big"
	Avatar1INSmall = "1in-small"
	Avatar2IN      = "2in"
	Avatar2INBig   = "2in-big"
	Avatar2INSmall = "2in-small"
)

var Avatar = Sizes{
	{
		Name:   Avatar1IN,
		Width:  295,
		Height: 413,
		Desc:   "2.50cm * 3.50cm",
	},
	{
		Name:   Avatar1INBig,
		Width:  390,
		Height: 567,
		Desc:   "大一寸 3.30cm * 4.80cm",
	},
	{
		Name:   Avatar1INSmall,
		Width:  260,
		Height: 378,
		Desc:   "小一寸 2.20cm * 3.20cm",
	},
	{
		Name:   Avatar2IN,
		Width:  413,
		Height: 579,
		Desc:   "3.50cm * 4.90cm",
	},
	{
		Name:   Avatar2INBig,
		Width:  413,
		Height: 626,
		Desc:   "大二寸 3.50cm * 5.30cm",
	},
	{
		Name:   Avatar2INSmall,
		Width:  413,
		Height: 531,
		Desc:   "小二寸 3.50cm * 4.50cm",
	},
}
