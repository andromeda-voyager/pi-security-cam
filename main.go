package main

import "piSecurityCam/message"

func main() {

	// for {
	// 	time.Sleep(5 * time.Second)
	// 	CheckForCommands()
	// 	if camera.IsMotionDetected() { // no motion is detected if camera is off
	// 		fmt.Println("Motion detected. Sending image.")
	// 		message.SendPicture()
	// 	}
	// }

	message.SendSMS("hi")
	// camera.TakePicture("test")
	// server.UploadImage("test", mediaURL)
}
