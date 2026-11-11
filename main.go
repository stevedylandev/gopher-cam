package main

import (
	"image"
	"log"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"

	gocv "gocv.io/x/gocv"
)

const (
	captureWidth  = 640
	captureHeight = 480
	captureFPS    = 30
	windowSize = 480
)

func main() {
	go func() {
		if err := run(); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func squareRect(cols, rows int) image.Rectangle {
	size := cols
	if rows < size {
		size = rows
	}
	x := (cols - size) / 2
	y := (rows - size) / 2
	return image.Rect(x, y, x+size, y+size)
}

func run() error {
	gocv.SetNumThreads(1)

	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		return err
	}
	defer webcam.Close()

	//webcam.Set(gocv.VideoCaptureFrameWidth, captureWidth)
	//webcam.Set(gocv.VideoCaptureFrameHeight, captureHeight)
	webcam.Set(gocv.VideoCaptureFPS, captureFPS)

	img := gocv.NewMat()
	defer img.Close()
	flipped := gocv.NewMat()
	defer flipped.Close()

	w := new(app.Window)
	w.Option(app.Title("gopher cam"))
	w.Option(app.Size(unit.Dp(windowSize), unit.Dp(windowSize)))

	ops := new(op.Ops)
	var frame paint.ImageOp

	for {
		// some bs so my window looks nicer than the default opencv 
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(ops, e)

			for {
				kev, ok := gtx.Event(key.Filter{Name: key.NameEscape})
				if !ok {
					break
				}
				if k, isKey := kev.(key.Event); isKey && k.State == key.Press {
					return nil
				}
			}

			if ok := webcam.Read(&img); ok && !img.Empty() {
				cropped := img.Region(squareRect(img.Cols(), img.Rows()))
				gocv.Flip(cropped, &flipped, 1)
				cropped.Close()
				if gi, err := flipped.ToImage(); err == nil {
					frame = paint.NewImageOp(gi)
				}
			}

			widget.Image{Src: frame, Fit: widget.Contain}.Layout(gtx)

			e.Frame(gtx.Ops)
			w.Invalidate()
		}
	}
}
