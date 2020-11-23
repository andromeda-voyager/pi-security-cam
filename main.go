package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

func createImage() *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, 100, 100))
}

func saveImage(fileName string, img *image.RGBA) {
	f, err := os.Create(fileName + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func draw(x, y int, img *image.RGBA) {
	var col color.Color
	col = color.RGBA{255, 0, 0, 255}
	img.Set(x, y, col)
}

func loadImage(fileName string) image.Image {
	imgFile, _ := os.Open(fileName)
	defer imgFile.Close()
	img, _ := png.Decode(imgFile)
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

func drawSquare(img *image.RGBA, c color.Color, x, y, l int) {
	initX := x
	for ; y < l; y++ {
		for x := initX; x < l; x++ {
			img.Set(x, y, c)
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

type message struct {
	size    int64
	gmailID string
	date    string // retrieved from message header
	snippet string
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

// bounds := img.Bounds()
// newImage := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
// draw.Draw(newImage, newImage.Bounds(), src, bounds.Min, draw.Src)
func main() {
	//sendEmail([]byte("testing"))

	//client := &http.Client{}
	//msgData := url.Values{}
	//msgData.Set("Body", "Hello there!")

	//msgDataReader := *strings.NewReader(msgData.Encode())
	//req, err := http.NewRequest("POST", "http://localhost:3000/upload-image", &msgDataReader)
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// req.Header.Add("Content-Type", "text/plain")
	msgbody := strings.NewReader("save my image")
	resp, _ := http.Post("http://localhost:3000/upload-image", "text/plain", msgbody)

	//f, _ := client.Do(req)

	//msg, _ := ioutil.ReadAll(f.Body)
	msg, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(msg))
}
