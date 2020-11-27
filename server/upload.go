package server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"piSecurityCam/camera"
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
func UploadImage(i, url string) string {
	imageName := randString(15)
	image := camera.LoadImage(i)
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, image, nil)
	b := buf.Bytes()
	imageStr := base64.StdEncoding.EncodeToString(b)
	imageUpload := &ImageUpload{imageStr, imageName, sign(imageName)}
	fmt.Println("in: " + imageName)
	j, _ := json.Marshal(imageUpload)
	resp, err := httpClient.Post(url, "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Printf("Start command error: %v", err)
	}
	respB, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respB))
	return string(respB)
}

//loaded from file implementation
// imageName := randString(15)
// 	file, _ := os.Open(f)
// 	reader := bufio.NewReader(file)
// 	b, _ := ioutil.ReadAll(reader)
// 	imageStr := base64.StdEncoding.EncodeToString(b)
// 	imageUpload := &ImageUpload{imageStr, imageName, sign(imageName)}

// func uploadImage2(image *image.RGBA) {
// 	msg := randString(15)
// 	buf := new(bytes.Buffer)
// 	err := jpeg.Encode(buf, image, nil)
// 	b := buf.Bytes()
// 	imageStr := base64.StdEncoding.EncodeToString(b)
// 	imageUpload := &ImageUpload{imageStr, msg, sign(msg)}
// 	j, _ := json.Marshal(imageUpload)
// 	resp, _ := httpClient.Post("http://localhost:3000/upload-image", "application/json", bytes.NewBuffer(j))
// 	respB, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println(string(respB))
// }
