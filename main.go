package main

import (
	"fmt"
	"piSecurityCam/camera"
	"piSecurityCam/control"
	"piSecurityCam/message"
	"time"
)

func main() {

	for {
		time.Sleep(5 * time.Second)
		control.CheckForCommands()
		if camera.IsMotionDetected() { // no motion is detected if camera is off
			fmt.Println("Motion detected. Sending image.")
			message.SendPicture()
		}
	}
}
