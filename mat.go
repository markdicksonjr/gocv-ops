package operations

import (
	"gocv.io/x/gocv"
	"image"
)

func CopyMat(mat *gocv.Mat) *gocv.Mat {
	newMat := gocv.NewMat()
	mat.CopyTo(&newMat)
	return &newMat
}

func Resize(img gocv.Mat, targetWidth, targetHeight int) *gocv.Mat {
	newMat := gocv.NewMat()
	gocv.Resize(img, &newMat, image.Pt(targetWidth, targetHeight), 0, 0, gocv.InterpolationLanczos4)
	return &newMat
}
