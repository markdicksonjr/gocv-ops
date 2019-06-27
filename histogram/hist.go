package histogram

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"math"
)

func DrawGrayscaleImageHistogram(img *gocv.Mat, w int, h int, numPoints int) *gocv.Mat {
	imgs := make([]gocv.Mat, 1)
	imgs[0] = *img

	hist := gocv.NewMat()

	channels := make([]int, 1)
	channels[0] = 0

	histSize := make([]int, 1)
	histSize[0] = numPoints

	ranges := make([]float64, 2)
	ranges[0] = 0
	ranges[1] = 256

	mask := gocv.NewMat()

	gocv.CalcHist(imgs, channels, mask, &hist, histSize, ranges, true)

	histW := w
	histH := h
	binW := int(math.Floor(float64(histW) / float64(histSize[0])))

	histImage := gocv.NewMatWithSize(histH, histW, gocv.MatTypeCV8UC3)
	gocv.Normalize(hist, &hist, 0, float64(histH), gocv.NormMinMax)

	for i := 1; i <= histSize[0]; i++ {
		atI := hist.GetFloatAt(i, 0)
		atI1 := hist.GetFloatAt(i - 1, 0)

		gocv.Line(&histImage,
			image.Pt(
				binW * (i-1),
				histH - int(atI1),
			),
			image.Pt(
				binW * i,
				histH - int(atI),
			),
			color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 1,
			},
			1,
		)
	}
	return &histImage
}
