package main

import (
	"flag"
	"fmt"
	"os"
	"pi-security-cam/camera"
	"pi-security-cam/control"
	"pi-security-cam/message"
	"time"
)

func main() {

	isDev := flag.Bool("raspi", false, "bool value for running on raspberrypi")
	flag.Parse()
	if *isDev {
		os.Setenv("device", "raspi")
	} else {
		os.Setenv("device", "unknown")
	}
	fmt.Println("Pi Security Cam Started. Running on", os.Getenv("device"))
	for {
		time.Sleep(5 * time.Second)
		control.CheckForCommands()
		if camera.IsMotionDetected() { // no motion is detected if camera is off
			fmt.Println("Motion detected. Sending image.")
			message.SendPicture()
		}
	}
}
