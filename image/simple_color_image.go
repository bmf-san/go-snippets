package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	x, y := 0, 0
	width, height := 400, 400
	quality := 100

	img := image.NewRGBA(image.Rect(x, y, width, height))
	for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
		for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
			img.Set(j, i, color.RGBA{255, 255, 255, 255})
		}
	}

	file, err := os.Create("sample.jpg")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	if err = jpeg.Encode(file, img, &jpeg.Options{quality}); err != nil {
		log.Println(err)
	}
}
