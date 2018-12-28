package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"runtime"

	"github.com/machinebox/sdk-go/facebox"
	"gocv.io/x/gocv"
)

var fbox = facebox.New("http://localhost:8080")

func NumberOfCPU() int {
	return runtime.NumCPU()
}

func main() {
	// open webcam
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("data/haarcascade_frontalface_default.xml") {
		fmt.Println("Error reading cascade file: data/haarcascade_frontalface_default.xml")
		return
	}

	fmt.Printf("start reading camera device: %v\n", 0)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %v\n", 0)
			return
		}
		if img.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		// draw a rectangle around each face on the original image
		for _, r := range rects {
			//gocv.Rectangle(&img, r, blue, 3)
			// Save each found face into the file
			imgFace := img.Region(r)
			//imgName := fmt.Sprintf("%d.jpg", time.Now().UnixNano())
			//gocv.IMWrite(imgName, imgFace)
			buf, err := gocv.IMEncode(".jpg", imgFace)
			imgFace.Close()
			if err != nil {
				log.Printf("unable to encode matrix: %v", err)
				continue
			}

			faces, err := fbox.Check(bytes.NewReader(buf))
			if err != nil {
				log.Printf("unable to recognize face: %v", err)
			}

			var caption = "Not found !!!!"
			if len(faces) > 0 {
				caption = fmt.Sprintf("I know you %s", faces[0].Name)
			}

			// draw rectangle for the face
			size := gocv.GetTextSize(caption, gocv.FontHersheyPlain, 3, 2)
			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)

			gocv.PutText(&img, caption, pt, gocv.FontHersheyPlain, 3, blue, 2)
			gocv.Rectangle(&img, r, blue, 3)
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		window.WaitKey(1)

	}
}
