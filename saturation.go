package operations

import (
	"gocv.io/x/gocv"
)

func DesaturateGray(img *gocv.Mat) *gocv.Mat {
	dest := img.Clone()
	gocv.CvtColor(*img, &dest, gocv.ColorGrayToBGR)
	gocv.CvtColor(dest, &dest, gocv.ColorBGRToHSV)
	DesaturateHSV(&dest)
	gocv.CvtColor(dest, &dest, gocv.ColorHSVToBGR)
	return &dest
}

func DesaturateBGR(img *gocv.Mat) *gocv.Mat {
	dest := gocv.NewMat()
	gocv.CvtColor(*img, &dest, gocv.ColorBGRToHSV)
	DesaturateHSV(&dest)
	gocv.CvtColor(dest, &dest, gocv.ColorHSVToBGR)
	return &dest
}

// TODO: unconfirmed
func DesaturateHSV(img *gocv.Mat) {
	for i := 0; i < img.Rows(); i++ {
		for j := 0; j < img.Cols(); j++ {
			img.SetUCharAt(i, j * img.Channels() + 1, 255)
		}
	}
}
