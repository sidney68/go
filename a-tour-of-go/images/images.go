package main

import (
    "golang.org/x/tour/pic"
    "image"
    "image/color"
)

type Image struct {
    width, height int
}

func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
    v := uint8(x ^ y)
    return color.RGBA{R: v, G: v, B: 255, A: 255}
}

func main() {
    m := Image{256, 64}
    pic.ShowImage(m)
}
