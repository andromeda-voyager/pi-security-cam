package main

import (
	"fmt"
	"piSecurityCam/camera"
	"time"
)

func main() {
	camera.TurnCameraOn()
	for {
		time.Sleep(5 * time.Second)
		if camera.IsMotionDetected() { // no motion is detected if camera is off
			fmt.Println("motion detected")
			sendPicture()
		}
	}

	//camera.TakePicture("new")

}
