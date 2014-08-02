package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct{
     x int
    y int
    v uint8
}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(i.x, i.y, 100, 100)
}

func (i Image) At(x, y int) color.Color {
    return color.RGBA{uint8(i.v), uint8(i.v), 255, 255}
}

func main() {
    m := Image{}
    pic.ShowImage(m)
}
