package main

import (
	"log"
	"os/exec"
)

func takePicture(fileName string) {
	//raspistill -t 2000 -o image.jpg -w 640 -h 480
	//	cmd := exec.Command("touch", "test.jpg")
	cmd := exec.Command("fswebcam", "--skip", "50", "-r", "640x480", "--no-banner", fileName+".jpg")
	err := cmd.Start()
	if err != nil {
		log.Printf("Start command error: %v", err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command error: %v", err)
	}
}

func turnOnCamera() {
	// take a picture every five seconds
}

func turnOffCamera() {

}
