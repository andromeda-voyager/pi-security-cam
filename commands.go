package main

func processCommand(c string) {
	switch c {
	case "on":
		turnOnCamera()
	case "off":
		turnOffCamera()
	case "snap":
		sendPicture()
	}
}

func sendPicture() {
	takePicture("snap")
	uploadImage("snap.jpg")
}
