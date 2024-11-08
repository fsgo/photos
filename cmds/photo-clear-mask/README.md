# 清理蒙层
用手机相机拍摄的试卷或者书页，在打印出来的时候后，页面整体会呈现灰蒙蒙的一片。
使用此程序，会将照片处理为黑白色，并将灰色背景替换为白色，并将字体的黑色等较深的颜色替换为纯黑色。
处理后打印出来黑更清晰。


## 安装
```bash
go install  github.com/fsgo/photos/cmds/photo-clear-mask@master
```

## 使用 
```bash
photo-clear-mask -help
```

```text
Usage of photo-clear-mask:
photo-clear-mask [file1] [file2] ...
-m uint8 用于清理的灰度中间值，>m 的会替换为白色，反之则替换为黑色 (default 120)
```

具体如：
```bash
photo-clear-mask 1.png 2.jpg
```

处理后输出：`1.png.cleared.png`  和 `2.png.cleared.png` 
