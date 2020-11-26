package main

func main() {
	// uploadImage()
	// takePicture("pic2")
	// timeNow := time.Now().UTC()
	// fmt.Println(timeNow.Format("DateSent>=2006-01-02T15:04:05"))

	// takePicture("pic1")
	// takePicture("pic2")

	// img1 := loadImage("pic1.jpg")
	// img2 := loadImage("pic2.jpg")
	// bounds := img1.Bounds()
	// newImage := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	// idraw.Draw(newImage, newImage.Bounds(), img1, bounds.Min, idraw.Src)
	// newImage2 := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	// idraw.Draw(newImage2, newImage.Bounds(), img2, bounds.Min, idraw.Src)

	// fmt.Println(imageDiff(newImage, newImage2))
	// averageImages(newImage, newImage2)
	//rand.Seed(time.Now().UTC().UnixNano())
	takePicture("snap")
	//fmt.Println(mediaURL + uploadImage("snap.jpg"))
	// fmt.Println(uploadImage("snap.jpg"))
	// sendSMS("test")
	imageURL := mediaURL + uploadImage("snap.jpg")
	sendMMS("test2", imageURL)

}
