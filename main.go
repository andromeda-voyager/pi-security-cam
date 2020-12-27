package main

import (
	"fmt"
	"pi-security-cam/camera"
	"pi-security-cam/control"
	"pi-security-cam/message"
	"runtime"
	"time"
)

func main() {

	fmt.Print("Pi Security Cam Started. Running on ", runtime.GOOS, " ", runtime.GOARCH, ".")
	for {
		time.Sleep(5 * time.Second)
		control.CheckForCommands()
		if camera.IsMotionDetected() { // no motion is detected if camera is off
			fmt.Println("Motion detected. Sending image.")
			message.SendPicture()
		}
	}
}
