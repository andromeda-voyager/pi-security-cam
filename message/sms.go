package message

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"piSecurityCam/camera"
	"piSecurityCam/config"
	"piSecurityCam/server"
	"strings"
)

// SendPicture uploads an image to server and sends an api request to twilio to send an mms with that picture
func SendPicture() {
	camera.TakePicture("snap")
	imageURL := config.ServerURL() + server.UploadImage("snap")
	SendMMS("Security Breach! Motion detected.", imageURL)
}

// SendSMS sends an api request to twilio to send an sms message
func SendSMS(message string) {

	msgData := url.Values{}
	msgData.Set("To", config.PersonalNumber())
	msgData.Set("From", config.TwilioNumber())
	msgData.Set("Body", message)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, err := http.NewRequest("POST", config.TwilioURL(), &msgDataReader)
	if err != nil {
		fmt.Printf("error creating request")
	}
	req.SetBasicAuth(config.AccountSid(), config.AuthToken())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error sending request")
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}

// SendMMS sends an api request to twilio to send an mms message
func SendMMS(message string, imageURL string) {

	msgData := url.Values{}
	msgData.Set("To", config.PersonalNumber())
	msgData.Set("From", config.TwilioNumber())
	msgData.Set("Body", message)
	msgData.Set("MediaUrl", imageURL)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", config.TwilioURL(), &msgDataReader)
	req.SetBasicAuth(config.AccountSid(), config.AuthToken())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}

	} else {
		fmt.Println(resp.Status)
	}
}
