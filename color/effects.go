package color

import "gocv.io/x/gocv"

func Brighten(img *gocv.Mat, alpha float64) *gocv.Mat {
	out := gocv.NewMat()
	black := gocv.NewMatWithSize(img.Rows(), img.Cols(), img.Type())
	gocv.BitwiseNot(black, &black)
	beta := 1 - alpha
	gocv.AddWeighted(*img, alpha, black, beta, 0.0, &out)
	return &out
}

// Darken will darken an RGB image, but "fade out" an image with alpha
func Darken(img *gocv.Mat, alpha float64) *gocv.Mat {
	out := gocv.NewMat()
	black := gocv.NewMatWithSize(img.Rows(), img.Cols(), img.Type())
	beta := 1 - alpha
	gocv.AddWeighted(*img, alpha, black, beta, 0.0, &out)
	return &out
}
