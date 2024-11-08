# 照片排版
可用于将较小的登记照（如 1 寸的），在相纸（如普通照片纸 6in-4R ）上排版，输出为排版好的图片文件。

## 安装
```bash
go install  github.com/fsgo/photos/cmds/photo-grid@master
```


## 使用
```bash
photo-grid -help
```

```
Usage of photo-grid:
  -b    border
  -g int
        avatar gap size (default 20)
  -in string
        avatar image path
  -is string
        avatar size, allow:
        1in :   2.50cm * 3.50cm
        1in-big :       大一寸 3.30cm * 4.80cm
        1in-small :     小一寸 2.20cm * 3.20cm
        2in :   3.50cm * 4.90cm
        2in-big :       大二寸 3.50cm * 5.30cm
        2in-small :     小二寸 3.50cm * 4.50cm
         (default "1in")
  -os string
        canvas size, allow:
        A0 :
        A1 :
        A2 :    42.00cm * 59.40
        A3 :    29.70cm * 42.00cm
        A4 :    21.00cm * 29.70cm
        A5 :    14.80cm * 21.00cm
        A6 :    10.50cm * 14.80cm
        6in-4R :        10.20cm * 15.20cm
        5in-3R :        12.70cm * 8.90cm
        7in-5R :        17.80cm * 12.70cm
        8in-6R :        20.30cm * 15.20cm
        10in-8R :       25.40cm * 20.30cm
         (default "6in-4R")
  -out string
        output image path (default "./out.png")
  -p int
        padding size (default 20)
```