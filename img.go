package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	green := color.RGBA{0, 255, 0, 255}
	red := color.RGBA{255, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	Img := image.NewRGBA(image.Rect(0, 0, 200, 200))

	draw.Draw(Img, Img.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	rectFile, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer rectFile.Close()
	png.Encode(rectFile, Img)

	draw.Draw(Img, Img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	for x := 20; x < 380; x++ {
		y := x/2 + 15
		Img.Set(x, y, red)
		Img.Set(x-20, y, red)
		Img.Set(x-40, y, red)
	}
	lineFile, err := os.Create("line.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer lineFile.Close()
	png.Encode(lineFile, Img)

}
