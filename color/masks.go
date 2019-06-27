package color

import (
	"gocv.io/x/gocv"
	"image"
)

func MaskRed(img gocv.Mat) *gocv.Mat {
	return MaskColorDualRange(img,
		gocv.NewScalar(0, 100, 100, 255),
		gocv.NewScalar(10,255,255,255),
		gocv.NewScalar(165, 50, 50, 255),
		gocv.NewScalar(179,255,255,255))
}

func MaskBlue(img gocv.Mat) *gocv.Mat {
	return MaskColor(img,
		gocv.NewScalar(100, 150, 0, 255),
		gocv.NewScalar(140, 255, 255, 255))
}

func MaskColor(img gocv.Mat, lowerMin gocv.Scalar, lowerMax gocv.Scalar) *gocv.Mat {
	hsv := gocv.NewMat()
	defer hsv.Close()

	// TODO: needed?
	gocv.GaussianBlur(img, &hsv, image.Pt(7, 7), 0, 0, gocv.BorderDefault)

	gocv.CvtColor(hsv, &hsv, gocv.ColorBGRToHSV)

	lowerRedHueRange := gocv.NewMat()
	gocv.InRangeWithScalar(hsv, lowerMin, lowerMax, &lowerRedHueRange)

	return &lowerRedHueRange
}

func MaskColorDualRange(img gocv.Mat, lowerMin gocv.Scalar, lowerMax gocv.Scalar, upperMin gocv.Scalar, upperMax gocv.Scalar) *gocv.Mat {
	hsv := gocv.NewMat()
	defer hsv.Close()

	// TODO: needed?
	gocv.GaussianBlur(img, &hsv, image.Pt(7, 7), 0, 0, gocv.BorderDefault)

	gocv.CvtColor(hsv, &hsv, gocv.ColorBGRToHSV)

	upperRedHueRange := gocv.NewMat()
	lowerRedHueRange := gocv.NewMat()
	redHueImage := gocv.NewMat()
	gocv.InRangeWithScalar(hsv, lowerMin, lowerMax, &lowerRedHueRange)
	gocv.InRangeWithScalar(hsv, upperMin, upperMax, &upperRedHueRange)
	gocv.AddWeighted(lowerRedHueRange, 1.0, upperRedHueRange, 1.0, 0.0, &redHueImage)

	return &redHueImage
}
