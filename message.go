package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"net/url"
	"strings"
	"time"
)

func sendEmail(msg []byte) {
	hostname := "smtp.gmail.com"
	auth := smtp.PlainAuth("", account,
		password, hostname)

	err := smtp.SendMail(hostname+":587", auth, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func sendSMS(message string) {

	msgData := url.Values{}
	msgData.Set("To", personalNumber)
	msgData.Set("From", twilioNumber)
	msgData.Set("Body", message)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", twilioURL, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
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

func sendMMS(message string, imageURL string) {

	msgData := url.Values{}
	msgData.Set("To", personalNumber)
	msgData.Set("From", twilioNumber)
	msgData.Set("Body", message)
	msgData.Set("MediaUrl", imageURL)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", twilioURL, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
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

func getMessages() {

	var response MessagesResponse
	timeNow := time.Now().UTC()
	//dateSent := timeNow.Format("DateSent>=2006-01-02T15:04:05")
	dateSent := timeNow.Format("DateSent>=2006-01-02")

	getMessagesURL := twilioURL + "?" + dateSent + "&PageSize=20&From=" + personalNumber

	req, _ := http.NewRequest("GET", getMessagesURL, nil)
	req.SetBasicAuth(accountSid, authToken)
	resp, _ := httpClient.Do(req)

	b, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, &response)

	fmt.Println(string(b))
	c := response.Messages[0].Body
	processCommand(c)
}
