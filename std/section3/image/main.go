package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {
	createImage()
	pixels()
	colors()
	readImage()
}

func createImage() {
	w, h := 300, 150
	react := image.Rect(0, 0, w, h)
	img := image.NewRGBA(react)

	f, err := os.Create("blank.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	png.Encode(f, img)
}

func pixels() {
	w, h := 10, 10
	react := image.Rect(0, 0, w, h)
	img := image.NewRGBA(react)
	fmt.Printf("Pixels: %v\n", img.Pix)
	fmt.Println("Pixels counts:", len(img.Pix))

	// set pixel at (1,0) to a sky blue color
	img.Pix[4] = 100 // every pixel has 4 bytes; R
	img.Pix[5] = 200 // every pixel has 4 bytes; G
	img.Pix[6] = 255 // every pixel has 4 bytes; B
	img.Pix[7] = 255 // every pixel has 4 bytes; A
	fmt.Printf("Pixels: %v\n", img.Pix)

	f, err := os.Create("pixels.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func colors() {
	w, h := 300, 200
	rect := image.Rect(0, 0, w, h)
	img := image.NewRGBA(rect)

	fmt.Printf("Bounds: %#v\nSize: %#v\n", img.Bounds(), img.Bounds().Size())
	size := img.Bounds().Size()

	blue := color.RGBA{
		R: uint8(100),
		G: uint8(200),
		B: uint8(255),
		A: uint8(255),
	}

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			img.Set(x, y, blue)

			// rand := color.RGBA{
			// 	R: uint8(200 * x / size.X),
			// 	G: uint8(200 * y / size.Y),
			// 	B: uint8(200),
			// 	A: uint8(255),
			// }
			// img.Set(x, y, rand)
		}
	}

	f, err := os.Create("colors.png")
	if err != nil {
		log.Fatalln(err)
	}

	jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
	f.Close()
}

func readImage() {
	f, err := os.Open("pixels.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	img, format, err := image.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Image Format: %s\n", format)
	fmt.Printf("Image Color Model: %#v\n", img.ColorModel())
	fmt.Printf("Image Bounds: %#v\n", img.Bounds())
	fmt.Printf("Image Color At (1,0): %#v\n", img.At(1, 0))
}
