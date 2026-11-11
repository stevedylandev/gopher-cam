package main

import (
	"image"
	gocv "gocv.io/x/gocv"
)

func squareCrop(src gocv.Mat) gocv.Mat {
	size := src.Cols()
	if src.Rows() < size {
		size = src.Rows()
	}
	x := (src.Cols() - size) / 2
	y := (src.Rows() - size) / 2
	return src.Region(image.Rect(x, y, x+size, y+size))
}

func main() {
	webcam, _ := gocv.OpenVideoCapture(0)
	window := gocv.NewWindow("gopher cam")
	img := gocv.NewMat()
	flipped := gocv.NewMat()
	defer flipped.Close()

	for {
		webcam.Read(&img)
		cropped := squareCrop(img)
		gocv.Flip(cropped, &flipped, 1)
		window.IMShow(flipped)
		window.WaitKey(1)
	}
}

