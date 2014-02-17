package cornerdetection

import "testing"
import "path/filepath"

func TestLoadImage(t *testing.T) {
	path := filepath.FromSlash("resources/images/5x5-clear.png")
	image := LoadImage(path)
	if image == nil {
		t.Errorf("LoadImage(%v) = %v, want image", path, image)
	}
	bounds := image.Bounds()
	if bounds.Max.X != 5 || bounds.Max.Y != 5 {
		t.Errorf("Unexpected image max bounds: %v", bounds)
	}

}
