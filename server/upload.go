package server

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pi-security-cam/camera"
	"pi-security-cam/config"
	"time"
)

//ImageUpload struct used to upload image to server
type ImageUpload struct {
	Image     string
	Message   string
	Signature string
}

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

// UploadImage uploads an image to the provided url
func UploadImage(i string) string {
	imageName := randString(15)
	image := camera.LoadImage(i)
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, image, nil)
	b := buf.Bytes()
	imageStr := base64.StdEncoding.EncodeToString(b)
	imageUpload := &ImageUpload{imageStr, imageName, sign(imageName)}
	j, _ := json.Marshal(imageUpload)
	resp, err := httpClient.Post(config.UploadURL(), "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Printf("Start command error: %v", err)
	}
	respB, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respB))
	return string(respB)
}

// UploadImageFromFile loads an image saved locally and uploads it to the provided url
func UploadImageFromFile(fileName string) string {
	imageName := randString(15)
	file, _ := os.Open(fileName + ".jpg")
	reader := bufio.NewReader(file)
	b, _ := ioutil.ReadAll(reader)
	imageStr := base64.StdEncoding.EncodeToString(b)
	imageUpload := &ImageUpload{imageStr, imageName, sign(imageName)}
	j, _ := json.Marshal(imageUpload)
	resp, err := httpClient.Post(config.UploadURL(), "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Printf("Start command error: %v", err)
	}
	respB, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respB))
	return string(respB)
}
