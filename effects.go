package operations

import (
	"gocv.io/x/gocv"
	"image"
)

func PadMat(mat *gocv.Mat, padX, padY int) *gocv.Mat {

	// fill a slightly larger mat with white
	padded := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(255,255,255,255), mat.Rows() + padY * 2, mat.Cols() + padX * 2, mat.Type())

	// copy the image to the inset of the larger image
	paddedRegion := padded.Region(image.Rect(padX, padY, mat.Cols() + padX, mat.Rows() + padY))
	mat.CopyTo(&paddedRegion)
	return &padded
}