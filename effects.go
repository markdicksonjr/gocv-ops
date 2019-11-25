package operations

import (
	"gocv.io/x/gocv"
	"image"
)

func PadMat(mat *gocv.Mat, padX, padY int) *gocv.Mat {

	// fill a slightly larger mat with white
	padded := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(255, 255, 255, 255), mat.Rows()+padY*2, mat.Cols()+padX*2, mat.Type())

	// copy the image to the inset of the larger image
	paddedRegion := padded.Region(image.Rect(padX, padY, mat.Cols()+padX, mat.Rows()+padY))
	mat.CopyTo(&paddedRegion)
	return &padded
}

// CLAHEOnNthChannel will do the 2.0x4x4 CLAHE on the nth channel
func CLAHEOnNthChannel(img gocv.Mat, n int, clipLimit float64, tileGridSize image.Point) gocv.Mat {
	r := gocv.Split(img)
	c := gocv.NewCLAHEWithParams(clipLimit, tileGridSize)
	c.Apply(r[n], &r[n])
	gocv.Merge(r, &img)
	_ = c.Close()

	return img
}

func ReduceGlareRGBA(img gocv.Mat) gocv.Mat {

	// RGBA => HSV
	gocv.CvtColor(img, &img, gocv.ColorRGBAToBGR)
	gocv.CvtColor(img, &img, gocv.ColorBGRToHSV)

	img = CLAHEOnNthChannel(img, 2, 2.0, image.Point{X: 4, Y: 4})

	// convert to RGB
	gocv.CvtColor(img, &img, gocv.ColorHSVToRGB)
	return img
}
