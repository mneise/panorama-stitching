package cornerdetection

import "image"
import _ "image/png"
import "os"
import "math"
import "image/color"

func LoadImage(path string) image.Image {
	f, _ := os.Open(path)
	defer f.Close()
	i, _, _ := image.Decode(f)

	return i
}

func windowAt(p image.Point, i image.Image, padding int) ([]image.Point, bool) {
	bounds := i.Bounds()

	if p.X-padding < bounds.Min.X || p.X+padding > bounds.Max.X ||
		p.Y-padding < bounds.Min.Y || p.Y+padding > bounds.Max.Y {
		return nil, false
	}

	var points []image.Point
	points = make([]image.Point, 9)

	index := 0
	for y := p.Y - padding; y <= p.Y+padding; y++ {
		for x := p.X - padding; x <= p.X+padding; x++ {
			points[index] = image.Point{x, y}
			index++
		}
	}

	return points, true
}

func Contains(list []image.Point, elem image.Point) bool {
	for _, t := range list {
		if t == elem {
			return true
		}
	}
	return false
}

func Sobel(point image.Point, i image.Image) (float64, bool){

	if point == image.ZP {
		return 0, false
	}

	kernelX := []float64{-1, 0, 1, -2, 0, 2, -1, 0, 1}
	var magX float64
	magX = convolutionWithKernel(point, i, kernelX)

	kernelY := []float64{1, 2, 1, 0, 0, 0, -1, -2, -1}
	var magY float64
	magY = convolutionWithKernel(point, i, kernelY)


	return math.Sqrt(math.Pow(magX, 2) + math.Pow(magY, 2)), true
}

func convolutionWithKernel(point image.Point, i image.Image, kernel []float64) float64 {

	var mag float64
	mag = 0

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			oldColor := i.At(col + point.X - 1, row + point.Y - 1)
            grayColor := color.GrayModel.Convert(oldColor).(color.Gray)
			mag += float64(grayColor.Y) * kernel[row*3 + col]
		}
	}

	return mag
}
