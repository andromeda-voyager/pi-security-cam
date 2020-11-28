package main

import (
	"fmt"
	"piSecurityCam/camera"
	"time"
)

func main() {

	for {
		time.Sleep(5 * time.Second)
		CheckForCommands()
		if camera.IsMotionDetected() { // no motion is detected if camera is off
			fmt.Println("motion detected")
			sendPicture()
		}
	}
	// camera.TakePicture("test")
	// server.UploadImage("test", mediaURL)
}
