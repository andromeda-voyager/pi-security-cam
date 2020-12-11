package camera

import (
	"fmt"
	"image"
	"log"
	"os/exec"
	"piSecurityCam/config"
	"time"
)

var referenceImg *image.RGBA
var isCameraOn bool
var updateInterval = 45 * time.Second
var lastReferenceUpdate time.Time

func init() {
	isCameraOn = false
}

func updateReferenceImage() {
	timeNow := time.Now().UTC()
	if timeNow.Sub(lastReferenceUpdate) > updateInterval {
		TakePicture("reference")
		img := LoadImage("reference")
		averageImages(referenceImg, img)
		fmt.Println("reference image updated")
		lastReferenceUpdate = timeNow
	}
}

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
	updateReferenceImage()
	TakePicture("capture")
	img := LoadImage("capture")
	imageDifference, err := getDiffValue(img, referenceImg)
	fmt.Println("image diff:", imageDifference)
	if err != nil {
		fmt.Println("An error occured getting image differences.")
		return false
	}
	if imageDifference > config.MaxImageDifference() {
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
	lastReferenceUpdate = time.Now().UTC()
}

// TurnCameraOff Sets the state of the camera to off and images are no longer taken
func TurnCameraOff() {
	isCameraOn = false
}
