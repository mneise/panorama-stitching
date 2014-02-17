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

func windowAt(p image.Point) ([]image.Point, bool) {
	if p == image.ZP {
		return nil, false
	}

	return []image.Point{image.ZP}, true
}
