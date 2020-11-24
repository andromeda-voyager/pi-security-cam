package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//ImageUpload struct used to upload image to server
type ImageUpload struct {
	Image     string
	Message   string
	Signature string
}

func uploadImage() {

	msg := randString(15)
	file, _ := os.Open("image1.jpg")
	reader := bufio.NewReader(file)
	b, _ := ioutil.ReadAll(reader)
	imageStr := base64.StdEncoding.EncodeToString(b)
	imageUpload := &ImageUpload{imageStr, msg, sign(msg)}
	j, _ := json.Marshal(imageUpload)
	resp, _ := http.Post("http://localhost:3000/upload-image", "application/json", bytes.NewBuffer(j))
	respB, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respB))
}
