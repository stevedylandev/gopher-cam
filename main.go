package main

import (
	"image"
	"log"

	gocv "gocv.io/x/gocv"
)

const (
	captureWidth  = 640
	captureHeight = 480
	captureFPS    = 30
)

func squareRect(cols, rows int) image.Rectangle {
	size := cols
	if rows < size {
		size = rows
	}
	x := (cols - size) / 2
	y := (rows - size) / 2
	return image.Rect(x, y, x+size, y+size)
}

func main() {
	gocv.SetNumThreads(1)

	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		log.Fatalf("open webcam: %v", err)
	}
	defer webcam.Close()

	webcam.Set(gocv.VideoCaptureFrameWidth, captureWidth)
	webcam.Set(gocv.VideoCaptureFrameHeight, captureHeight)
	webcam.Set(gocv.VideoCaptureFPS, captureFPS)

	window := gocv.NewWindow("gopher cam")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()
	flipped := gocv.NewMat()
	defer flipped.Close()

	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			continue
		}

		cropped := img.Region(squareRect(img.Cols(), img.Rows()))
		gocv.Flip(cropped, &flipped, 1)
		cropped.Close()

		window.IMShow(flipped)
		if window.WaitKey(1) == 27 || window.GetWindowProperty(gocv.WindowPropertyVisible) < 1 {
			return
		}
	}
}
