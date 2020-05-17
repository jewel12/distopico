package telescreen

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

const squareBaseSize = 512
const squareSize = 512

func genSlogan() *gocv.Mat {
	m := gocv.NewMatWithSize(squareSize, squareSize, gocv.MatTypeCV8UC3)
	pts := [][]image.Point{
		{
			{0, 0},
			{0, squareSize},
			{squareSize, squareSize},
			{squareSize, 0},
		},
	}
	red := color.RGBA{255, 0, 0, 0}
	gocv.FillPoly(&m, pts, red)

	bigBrotherIs := "BIG BROTHER IS"
	bscl := float64(1.71 * (squareSize / squareBaseSize))
	bftn := 5 * (squareSize / squareBaseSize)
	bcolor := color.RGBA{0, 0, 0, 0}
	bsz := gocv.GetTextSize(bigBrotherIs, gocv.FontHersheyDuplex, bscl, bftn)
	bpnt := image.Point{
		(squareSize - bsz.X) / 2,
		squareSize - (squareSize / 3),
	}
	gocv.PutText(&m, bigBrotherIs, bpnt, gocv.FontHersheyDuplex, bscl, bcolor, bftn)

	watchingYou := "WATCHING YOU"
	wscl := float64(1.8 * (squareSize / squareBaseSize))
	wftn := 6 * (squareSize / squareBaseSize)
	wcolor := color.RGBA{255, 255, 255, 0}
	wsz := gocv.GetTextSize(watchingYou, gocv.FontHersheyDuplex, wscl, wftn)
	wpnt := image.Point{
		(squareSize - wsz.X) / 2,
		bpnt.Y + bsz.Y + (squareSize / 35),
	}
	gocv.PutText(&m, watchingYou, wpnt, gocv.FontHersheyDuplex, wscl, wcolor, wftn)

	return &m
}

func genScreen(cap *gocv.Mat) *gocv.Mat {
	img := cap.Clone()
	sz := img.Size()
	gocv.Resize(img, &img, image.Point{0, 0},
		512.0/float64(sz[0]), 512.0/float64(sz[1]), gocv.InterpolationDefault)

	pts := [][]image.Point{
		{{0, 0}, {0, 512}, {512, 512}, {512, 0}},
	}
	gocv.Polylines(&img, pts, true, color.RGBA{0, 255, 0, 0}, 20)

	gocv.Circle(&img, image.Point{60, 48}, 24,
		color.RGBA{255, 0, 0, 0}, -1)

	gocv.PutText(&img, "REC", image.Point{90, 70},
		gocv.FontHersheyDuplex, 2.0, color.RGBA{255, 255, 255, 0}, 5)

	return &img
}
