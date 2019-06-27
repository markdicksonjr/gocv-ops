package geometry

import (
	"image"
	"math"
)

// TODO: ONLY COMBINES INTO PAIRS
func CombineOverlappingXRectangles(boundingRectangles []image.Rectangle) []image.Rectangle {
	var combinedRectangles []image.Rectangle
	var globbedRectangleIndices = make(map[int]int)
	for i := 0; i < len(boundingRectangles); i++ {

		// if we haven't joined this rect yet
		if globbedRectangleIndices[i] == 0 {
			boundingRect1 := boundingRectangles[i]

			rect1 := image.Rectangle{
				Min: image.Pt(boundingRect1.Min.X, 0),
				Max: image.Pt(boundingRect1.Max.X, 1),
			}

			for j := 0; j < len(boundingRectangles); j++ {
				if j != i && globbedRectangleIndices[j] != 1 {
					boundingRect2 := boundingRectangles[j]

					rect2 := image.Rectangle{
						Min: image.Pt(boundingRect2.Min.X, 0),
						Max: image.Pt(boundingRect2.Max.X, 1),
					}

					if rect1.Overlaps(rect2) {
						boundingRect1 = boundingRect1.Union(boundingRect2)
						globbedRectangleIndices[j] = 1
					}
				}
			}

			combinedRectangles = append(combinedRectangles, boundingRect1)
			globbedRectangleIndices[i] = 1
		}
	}

	return combinedRectangles
}

func ExpandRectangle(rect image.Rectangle, rowSize int, colSize int) image.Rectangle {
	rect.Min.X -= colSize
	rect.Min.Y -= rowSize
	rect.Max.X += colSize
	rect.Max.Y += rowSize
	return rect
}

func EnsureRectBounds(rect image.Rectangle, rows int, cols int) image.Rectangle {
	rect.Min.X = int(math.Max(float64(rect.Min.X), 0))
	rect.Min.Y = int(math.Max(float64(rect.Min.Y), 0))
	rect.Max.X = int(math.Min(float64(rect.Max.X), float64(cols)))
	rect.Max.Y = int(math.Min(float64(rect.Max.Y), float64(rows)))

	return rect
}
