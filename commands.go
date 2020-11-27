package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"piSecurityCam/camera"
	"piSecurityCam/server"
	"strings"
	"time"
)

func processCommand(c string) {
	switch strings.ToLower(c) {
	case "on":
		camera.TurnCameraOff()
	case "off":
		camera.TurnCameraOn()
	case "snap":
		sendPicture()
	case "status":
		sendSMS(camera.Status())
	case "help":
		sendSMS(help())
	case "": //getCommands returns empty string if there are no new sms commands
	default:
		sendSMS("Command not found." + help())
	}
}

func help() string {
	return `Commands:
		\non :Turns the camera on.
		\noff : Turns the camera off
		\nsnap : Takes a picture and sends it to you
		\nstatus : Sends the state of the camera.`
}

func sendPicture() {
	camera.TakePicture("snap")
	imageURL := mediaURL + server.UploadImage("snap", mediaURL)
	sendMMS("test2", imageURL)
}

// MessagesResponse used by getMessages to populate json fields
type MessagesResponse struct {
	Body      string    `json:"body"`
	FirstPage string    `json:"first_page_uri"`
	NextPage  string    `json:"next_page_uri"`
	Page      int       `json:"page"`
	PageSize  int       `json:"page_size"`
	Messages  []Message `json:"messages"`
}

// Message used by getMessages to populate json fields
type Message struct {
	Body        string `json:"body"`
	DateCreated string `json:"date_created"`
	DateSent    string `json:"date_sent"`
	From        int    `json:"from"`
}

var lastUpdateTime string
var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

// CheckForCommands processes the commands received by sms
func CheckForCommands() {
	processCommand(getCommand())
}

func getCommand() string {
	var response MessagesResponse
	getMessagesURL := twilioURL + "?" + getLastUpdateTime() + "&PageSize=20&From=" + personalNumber
	req, _ := http.NewRequest("GET", getMessagesURL, nil)
	req.SetBasicAuth(accountSid, authToken)
	resp, _ := httpClient.Do(req)

	b, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, &response)
	if len(response.Messages) == 0 {
		return ""
	}
	return response.Messages[0].Body
}

func getLastUpdateTime() string {
	defer resetLastUpdateTime()
	if len(lastUpdateTime) == 0 {
		resetLastUpdateTime()
	}
	return lastUpdateTime
}

func resetLastUpdateTime() {
	lastUpdateTime = time.Now().UTC().Format("DateSent>=2006-01-02T15:04:05")
}
