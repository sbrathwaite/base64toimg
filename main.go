package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
  fmt.Println("Hello")
}

func Base64ToJpeg(imageData string) {
	options := jpeg.Options{Quality: 100}

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageData))

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	home, _ := os.UserHomeDir()

	jpegFile := fmt.Sprintf("%s/tempimages/%s.jpeg", home, time.Now().Format(dateFormat))

	if _, err := os.Stat(fmt.Sprintf("%s/tempimages", home)); os.IsNotExist(err) {
		os.Mkdir(fmt.Sprintf("%s/tempimages", home), 0644)
	}

	f, err := os.OpenFile(jpegFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	err = jpeg.Encode(f, m, &options)
	if err != nil {
		log.Fatal(err)
	}
}
