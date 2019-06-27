package contours

import (
	"github.com/markdicksonjr/gocv-ops"
	"github.com/markdicksonjr/gocv-ops/geometry"
	"gocv.io/x/gocv"
	"image"
	"math"
	"sort"
)

func ThresholdAndFindLargestContour(grayMatIn *gocv.Mat) []image.Point {
	matOut := gocv.NewMat()
	defer matOut.Close()
	gocv.Threshold(*grayMatIn, &matOut, 0, 255, gocv.ThresholdBinary | gocv.ThresholdOtsu)
	return GetLargestContour(gocv.FindContours(matOut, gocv.RetrievalExternal, gocv.ChainApproxSimple))
}

func CropToLargestRectWithInset(grayMatIn *gocv.Mat, inset int) (*gocv.Mat, []image.Point) {
	maxContour := ThresholdAndFindLargestContour(grayMatIn)

	rect := gocv.BoundingRect(maxContour)
	rect = geometry.EnsureRectBounds(geometry.ExpandRectangle(rect, inset, inset), grayMatIn.Rows(), grayMatIn.Cols())

	// snag the region
	matRegion := grayMatIn.Region(rect)
	matRegionClone := matRegion.Clone()
	return &matRegionClone, maxContour
}

func CropBoundingRectangles(img *gocv.Mat, boundingRectangles []image.Rectangle, padX, padY int) []*gocv.Mat {

	// sort the rectangles from left to right
	sort.SliceStable(boundingRectangles, func(i, j int) bool {
		iRect := boundingRectangles[i]
		jRect := boundingRectangles[j]

		return iRect.Min.X < jRect.Min.X
	})

	// find the two closest consecutive bounding rectangles (TODO: do we use this?)
	var minDist = math.MaxInt32
	for i := 0; i < len(boundingRectangles) - 1; i++ {
		dist := boundingRectangles[i + 1].Min.X - boundingRectangles[i].Max.X

		// if the distance to the nearest rectangle is the smallest we've seen
		// make this distance the new "smallest"
		if dist > 0 && dist < minDist {
			minDist = dist
		}
	}

	// put a mat of each contour into the mats array
	var mats []*gocv.Mat
	for i := 0; i < len(boundingRectangles); i++ {

		// get a sub-mat reference to the region within the image
		rect := boundingRectangles[i]
		region := img.Region(image.Rect(rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y))

		// copy the region into a new mat
		regionCopy := operations.CopyMat(&region)

		// get a new mat with white around the original region, clean up the original copy of the region
		paddedMat := operations.PadMat(regionCopy, padX, padY)
		regionCopy.Close()

		// add the padded region to the list of images
		mats = append(mats, paddedMat)
	}

	return mats
}
