package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

type setting struct {
	printText string
	iniText   string
}

var settings = []setting{
	setting{"email account", "email_account"},
	setting{"password", "password"},
	setting{"Twilio account sid", "twilio_account_sid"},
	setting{"Twilio auth token", "twilio_auth_token"},
	setting{"Twilio number", "twilio_number"},
	setting{"personal number", "personal_number"},
	setting{"upload url", "upload_url"},
	setting{"server url", "server_url"},
	setting{"max image difference", "max_image_difference"},
}

func writeSettingsFile() {

	fmt.Println("Creating settings.ini file. Please enter a value for each setting.")
	var fileText string
	scanner := bufio.NewScanner(os.Stdin)
	for _, s := range settings {
		fmt.Print(s.printText + ": ")
		scanner.Scan()
		fileText += s.iniText + " = " + scanner.Text() + "\n"
	}
	err := ioutil.WriteFile("settings.ini", []byte(fileText), 0600)
	if err != nil {
		fmt.Println("Unable to create file")
	}
}
