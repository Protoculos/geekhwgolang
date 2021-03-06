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
	red := color.RGBA{200, 30, 30, 255}
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	draw.Draw(img, img.Bounds(), &image.Uniform{green}, image.ZP,
		draw.Src)
	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()

	draw.Draw(img, img.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
	for i := 0; i < 200; i += 10 {
		for x := 0; x < 200; x++ {
			y := 10 + i
			img.Set(x, y, red)
		}
		for y := 0; y < 200; y++ {
			x := 10 + i
			img.Set(x, y, red)
		}
	}

	png.Encode(file, img)
}
