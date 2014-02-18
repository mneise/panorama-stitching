package cornerdetection

import "image"
import _ "image/png"
import "os"

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

func Sobel(point image.Point, i image.Image) (int, bool){

	if point == image.ZP {
		return 0, false
	}

	return 0, true
}
