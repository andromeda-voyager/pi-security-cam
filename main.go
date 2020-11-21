package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/smtp"
	"os"
)

func createImage() *image.NRGBA {
	return image.NewNRGBA(image.Rect(0, 0, 100, 100))
}

func saveImage(fileName string, img *image.NRGBA) {
	f, err := os.Create(fileName + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func draw(x, y int, img *image.NRGBA) {
	var col color.Color
	col = color.RGBA{255, 0, 0, 255}
	img.Set(x, y, col)
}

func openImage(fileName string) image.Image {
	imgFile, _ := os.Open(fileName)
	defer imgFile.Close()
	img, _ := png.Decode(imgFile)
	return img
}

func compareImages(img1, img2 *image.NRGBA) {
	for y := 0; y < 100; y++ {
		for x := img1.Bounds().Min.X; x < img1.Bounds().Max.X; x++ {
			if (img1.At(x, y)) != img2.At(x, y) {
				fmt.Println("picture does not match at", x, y)
			}
		}
	}
}

func sendEmail(msg []byte) {
	hostname := "smtp.gmail.com"
	auth := smtp.PlainAuth("", account,
		password, hostname)

	err := smtp.SendMail(hostname+":587", auth, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}
}

// bounds := img.Bounds()
// newImage := image.NewNRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
// draw.Draw(newImage, newImage.Bounds(), src, bounds.Min, draw.Src)
func main() {
	//sendEmail([]byte("testing"))

	img1 := createImage()
	img2 := createImage()

	draw(0, 0, img1)
	draw(99, 99, img2)

	compareImages(img1, img2)

}
