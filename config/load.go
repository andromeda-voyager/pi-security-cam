package config

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {
	load()
}

func setValue(setting, value string) {
	var err error
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
	case "max_image_difference":
		maxImageDifference, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			fmt.Println("Unable to load the max_image_difference value in settings.ini. Please change this value to an integer value.")
			os.Exit(0)
		}
	}
}

// Read the settings.ini file and set the values in config.go
func load() {

	configFile, err := os.Open("settings.ini")
	for err != nil {
		fmt.Println("Settings.ini not found.")
		writeSettingsFile()
		configFile, err = os.Open("settings.ini")
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
