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

func getMessages() {
	timeNow := time.Now()
	fmt.Println(timeNow.Format("DateSent>=2006-01-02 15:04:05-0700"))
	getMessagesURL := twilioURL + "?DateSent%3C=2019-03-01&PageSize=20&From=" + personalNumber
	resp, _ := httpClient.Get(getMessagesURL)
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}
