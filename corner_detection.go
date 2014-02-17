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
