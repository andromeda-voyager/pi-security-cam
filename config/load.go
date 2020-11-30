package config

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func init() {
	read()
}

func setValue(setting, value string) {
	switch setting {
	case "email_account":
		emailAccount = value
	case "twilio_account_sid":
		accountSid = value
	case "twilio_auth_token":
		authToken = value
	case "twilio_number":
		twilioNumber = value
	case "personal_number":
		personalNumber = value
	case "upload_url":
		uploadURL = value
	case "server_url":
		serverURL = value
	}
}

// Read the config file
func read() {

	configFile, err := os.Open("settings.ini") // For read access.
	if err != nil {
		log.Fatal("Unable to open settings.ini", err)
	}

	bufferedReader := bufio.NewReader(configFile)
	var lineNumber int
	for {
		fileLine, err := bufferedReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		lineNumber++
		if !(strings.HasPrefix(fileLine, "#") || strings.TrimSpace(fileLine) == "") {
			field := strings.Split(fileLine, "=")
			if len(field) < 2 {
				log.Fatalf("Configuration file is corrupt on line %v", lineNumber)
			}

			setting := strings.TrimSpace(field[0])
			value := strings.TrimSpace(field[1])
			setValue(setting, value)
		}
	}
}
