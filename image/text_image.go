package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func main() {
	baseFile, err := os.Open("./image/base.jpg")
	if err != nil {
		log.Println(err)
	}
	defer baseFile.Close()
	baseImage, _, err := image.Decode(baseFile)
	if err != nil {
		log.Println(err)
	}

	fontFile, err := ioutil.ReadFile("./font/Roboto-Regular.ttf")
	if err != nil {
		log.Println(err)
	}
	parsedFont, err := truetype.Parse(fontFile)
	if err != nil {
		log.Println(err)
	}

	r := baseImage.Bounds()
	rgbaImage := image.NewRGBA(image.Rect(0, 0, r.Dx(), r.Dy()))
	draw.Draw(rgbaImage, rgbaImage.Bounds(), baseImage, r.Min, draw.Src)
	drawer := font.Drawer{
		Dst: rgbaImage,
		Src: image.Black,
	}
	drawer.Face = truetype.NewFace(parsedFont, &truetype.Options{
		Size: 20,
		DPI:  350,
	})
	drawText := "Hello World"
	drawer.Dot = fixed.Point26_6{
		X: (fixed.I(r.Dx()) - drawer.MeasureString(drawText)) / 2,
		Y: fixed.I(r.Dy() / 2),
	}

	file, err := os.Create("sample_text.jpg")
	if err != nil {
		log.Println(err)
	}
	drawer.DrawString(drawText)
	if err = jpeg.Encode(file, drawer.Dst, &jpeg.Options{Quality: 100}); err != nil {
		log.Println(err)
	}
}
