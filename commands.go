package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"piSecurityCam/camera"
	"strings"
	"time"
)

func processCommand(c string) {
	fmt.Println("processing command :" + c + ":")
	switch strings.ToLower(c) {
	case "on":
		camera.TurnCameraOn()
	case "off":
		camera.TurnCameraOff()
	case "snap":
		sendPicture()
	case "status":
		sendSMS(camera.Status())
	// case "":
	// 	sendSMS(help())
	case "": //getCommands returns empty string if there are no new sms commands
	default:
		sendSMS("Command not found.\n" + help())
	}
}

func help() string {
	return `Commands:
		on :Turns the camera on.
		off : Turns the camera off.
		snap : Take a picture.
		status : Get the camera status.`
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

var lastMsgReceivedTime string
var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

// CheckForCommands processes the commands received by sms
func CheckForCommands() {
	processCommand(getCommand())
}

func getCommand() string {
	var response MessagesResponse
	getMessagesURL := twilioURL + "?" + getLastUpdateTime() + "&PageSize=5&From=" + personalNumber
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
	defer updateLastMsgReceivedTime()
	if len(lastMsgReceivedTime) == 0 {
		updateLastMsgReceivedTime()
	}
	return lastMsgReceivedTime
}

func updateLastMsgReceivedTime() {
	lastMsgReceivedTime = time.Now().UTC().Format("DateSent>=2006-01-02T15:04:05")
}
