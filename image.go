package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

// bounds := img.Bounds()
// newImage := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
// draw.Draw(newImage, newImage.Bounds(), src, bounds.Min, draw.Src)

func createImage() *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, 100, 100))
}

func saveImage(fileName string, img *image.RGBA) {
	f, err := os.Create(fileName + ".jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	option := &jpeg.Options{Quality: 75}
	jpeg.Encode(f, img, option)
}

func draw(x, y int, img *image.RGBA) {
	var col color.Color
	col = color.RGBA{255, 0, 0, 255}
	img.Set(x, y, col)
}

func loadImage(fileName string) image.Image {
	imgFile, _ := os.Open(fileName)
	defer imgFile.Close()
	img, _ := jpeg.Decode(imgFile)
	return img
}

func findImageDifferences(img1, img2 *image.RGBA) {
	for y := 0; y < 100; y++ {
		for x := img1.Bounds().Min.X; x < img1.Bounds().Max.X; x++ {
			if (img1.At(x, y)) != img2.At(x, y) {
				fmt.Println("picture does not match at", x, y)
			}
		}
	}
}

func averageImages(img1, img2 *image.RGBA) {
	bounds := img1.Bounds()
	avgImg := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	for i := 0; i < len(img1.Pix); i++ {
		avgImg.Pix[i] = uint8((uint16(img1.Pix[i]) + uint16(img2.Pix[i])) / 2)
	}
	saveImage("avg.jpg", avgImg)
}

func drawSquare(img *image.RGBA, c color.Color, x, y, l int) {
	initX := x
	for ; y < l; y++ {
		for x := initX; x < l; x++ {
			img.Set(x, y, c)
		}
	}
}

func squareDifference(x, y uint8) uint64 {
	diff := uint64(x) - uint64(y)
	return diff * diff
}

func imageDiff(img1, img2 *image.RGBA) (int64, error) {
	if img1.Bounds() != img2.Bounds() {
		return 0, errors.New("images are not the same size")
	}

	diffSum := int64(0)
	for i := 0; i < len(img1.Pix); i++ {
		diffSum += int64(squareDifference(img1.Pix[i], img2.Pix[i]))
	}

	return int64(math.Sqrt(float64(diffSum))), nil
}

func compareImages() {
	img1 := createImage()
	img2 := createImage()

	draw(0, 0, img1)
	draw(99, 99, img2)

	var c color.Color
	c = color.RGBA{255, 0, 0, 255}
	drawSquare(img1, c, 0, 0, 30)
	saveImage("image1", img1)
	saveImage("image2", img2)

	fmt.Println(imageDiff(img1, img2))
}
