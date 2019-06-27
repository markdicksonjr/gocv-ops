package contours

import (
	"gocv.io/x/gocv"
	"image"
)

func GetContoursWithinAreaRange(contours [][]image.Point, lowerLimit int, upperLimit int) [][]image.Point {
	var matchingContours [][]image.Point
	for _, shape := range contours {
		rect := gocv.BoundingRect(shape)
		area := rect.Dx() * rect.Dy()

		// if the shape is between the upper and lower bound of acceptable numbers
		if area >= lowerLimit && area <= upperLimit && rect.Dy() > rect.Dx() {
			matchingContours = append(matchingContours, shape)
		}
	}

	return matchingContours
}

func GetContoursIntersectingRectangle(contours [][]image.Point, rect image.Rectangle) [][]image.Point {
	var matchingContours [][]image.Point
	for _, shape := range contours {
		contourRect := gocv.BoundingRect(shape)

		if rect.Overlaps(contourRect) {
			matchingContours = append(matchingContours, shape)
		}
	}
	return matchingContours
}

func GetContoursBeyondAspectRatio(contours [][]image.Point, aspect float64) [][]image.Point {
	var matchingContours [][]image.Point
	for _, shape := range contours {
		rect := gocv.BoundingRect(shape)

		if (float64(rect.Dy()) / float64(rect.Dx())) > aspect {
			matchingContours = append(matchingContours, shape)
		}
	}

	return matchingContours
}

func GetBoundingBoxes(contours [][]image.Point) []image.Rectangle {
	var rects []image.Rectangle
	for _, shape := range contours {
		rect := gocv.BoundingRect(shape)
		rects = append(rects, rect)
	}
	return rects
}

func GetBoundingBox(contours [][]image.Point) image.Rectangle {
	var rect image.Rectangle
	for _, shape := range contours {
		thisRect := gocv.BoundingRect(shape)
		rect = rect.Union(thisRect)
	}
	return rect
}

func GetLargestContour(imgPoints [][]image.Point) []image.Point {
	var maxArea float64
	var maxContour []image.Point
	for _, shape := range imgPoints {
		area := gocv.ContourArea(shape)
		if area > maxArea {
			maxArea = area
			maxContour = shape
		}
	}

	return maxContour
}
