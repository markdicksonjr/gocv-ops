package operations

import (
	"gocv.io/x/gocv"
	"image"
	"math"
)

func CopyMat(mat *gocv.Mat) *gocv.Mat {
	newMat := gocv.NewMat()
	mat.CopyTo(&newMat)
	return &newMat
}

func ResizeCopy(img gocv.Mat, targetWidth, targetHeight int) *gocv.Mat {
	newMat := gocv.NewMat()
	gocv.Resize(img, &newMat, image.Pt(targetWidth, targetHeight), 0, 0, gocv.InterpolationLanczos4)
	return &newMat
}

func Resize(img *gocv.Mat, targetWidth, targetHeight int) {
	gocv.Resize(*img, img, image.Pt(targetWidth, targetHeight), 0, 0, gocv.InterpolationLanczos4)
}

// ResizeWithoutStretching will chop the input image at the edges to the target aspect ratio, then scale it
func ResizeWithoutStretching(mat *gocv.Mat, targetWidth, targetHeight int) {
	aspectRatio := float64(targetWidth) / float64(targetHeight)
	croppedHeight := int(math.Ceil(float64(mat.Cols()) / aspectRatio))
	croppedWidth := int(math.Ceil(float64(croppedHeight) * aspectRatio))

	x0 := mat.Cols()/2.0 - croppedWidth/2.0
	y0 := mat.Rows()/2.0 - croppedHeight/2.0
	newMat := mat.Region(image.Rect(
		int(math.Max(float64(x0), 0)),
		int(math.Max(float64(y0), 0)),
		int(math.Min(float64(x0+croppedWidth), float64(mat.Cols()))),
		int(math.Min(float64(y0+croppedHeight), float64(mat.Rows())))))
	*mat = newMat

	gocv.Resize(*mat, mat, image.Pt(targetWidth, targetHeight), 0, 0, gocv.InterpolationLanczos4)
}

// ResizeWithoutStretchingCopy will copy the input image before performing ResizeWithoutStretching
func ResizeWithoutStretchingCopy(img gocv.Mat, targetWidth, targetHeight int) *gocv.Mat {
	newMat := CopyMat(&img)
	ResizeWithoutStretching(newMat, targetWidth, targetHeight)
	return newMat
}
