package control

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pi-security-cam/camera"
	"pi-security-cam/config"
	"pi-security-cam/message"
	"strings"
	"time"
)

func processCommand(c string) {
	switch strings.ToLower(c) {
	case "on":
		camera.TurnCameraOn()
	case "off":
		camera.TurnCameraOff()
	case "snap":
		message.SendPicture()
	case "status":
		message.SendSMS(camera.Status())
	case "": //getCommands returns empty string if there are no new sms commands
	default:
		message.SendSMS("Command not found.\n" + help())
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
	getMessagesURL := config.TwilioURL() + "?DateSent>=" + getLastUpdateTime() + "&PageSize=5&From=" + config.PersonalNumber()
	req, _ := http.NewRequest("GET", getMessagesURL, nil)
	req.SetBasicAuth(config.AccountSid(), config.AuthToken())
	resp, _ := httpClient.Do(req)

	b, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, &response)
	if len(response.Messages) == 0 {
		return ""
	}
	return response.Messages[0].Body
}

func getLastUpdateTime() string {
	defer setLastUpdateTime()
	if len(lastUpdateTime) == 0 {
		setLastUpdateTime()
	}
	return lastUpdateTime
}

func setLastUpdateTime() {
	lastUpdateTime = time.Now().UTC().Format("2006-01-02T15:04:05")
}
