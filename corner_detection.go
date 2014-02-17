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
	if p == image.ZP {
		return nil, false
	}

	bounds := i.Bounds()

	if p.X - padding < bounds.Min.X || p.X + padding > bounds.Max.X ||
       p.Y - padding < bounds.Min.Y || p.Y + padding > bounds.Max.Y {
		return nil, false
	}

	return []image.Point{image.ZP, image.ZP, image.ZP, image.ZP, image.ZP, image.ZP, image.ZP, image.ZP}, true
}

func Contains(list []image.Point, elem image.Point) bool {
        for _, t := range list { if t == elem { return true } }
        return false
}
