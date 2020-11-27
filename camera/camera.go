package camera

import (
	"fmt"
	"image"
	"log"
	"os/exec"
)

const maxDifference = 20000

var referenceImg *image.RGBA
var isCameraOn = false

// TakePicture Takes a pictures with the camera and saves it to a file
func TakePicture(fileName string) {
	//raspistill -t 2000 -o image.jpg -w 640 -h 480
	//	cmd := exec.Command("touch", "test.jpg")
	cmd := exec.Command("fswebcam", "--skip", "50", "-r", "640x480", "--no-banner", imagesDirectory+fileName+".jpg")
	err := cmd.Start()
	if err != nil {
		log.Printf("Start command error: %v", err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command error: %v", err)
	}
}

// IsMotionDetected Detects motion based on differences in images
func IsMotionDetected() bool {
	if !isCameraOn {
		return false
	}
	TakePicture("capture")
	img := LoadImage("capture")
	imageDifference, err := getDiffValue(img, referenceImg)
	if err != nil {
		fmt.Println("An error occured getting image differences.")
		return false
	}
	if imageDifference > maxDifference {
		fmt.Println(imageDifference)
		return true
	}
	return false
}

// Status returns a string that says what the state of the camera is (on/off)
func Status() string {
	if isCameraOn {
		return "Camera is on."
	}
	return "Camera is off."
}

// TurnCameraOn Sets the state of the camera to on and images can now be taken
func TurnCameraOn() {
	isCameraOn = true
	TakePicture("reference")
	referenceImg = LoadImage("reference")
}

// TurnCameraOff Sets the state of the camera to off and images are no longer taken
func TurnCameraOff() {
	isCameraOn = false
}