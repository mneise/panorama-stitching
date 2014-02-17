package cornerdetection

import "image"
import "path/filepath"
import "testing"

func TestLoadImage(t *testing.T) {

	path := filepath.FromSlash("resources/images/5x5-clear.png")
	i := LoadImage(path)
	if i == nil {
		t.Errorf("LoadImage(%v) = %v, want image", path, i)
	}
	bounds := i.Bounds()
	if bounds.Max.X != 5 || bounds.Max.Y != 5 {
		t.Errorf("Unexpected image max bounds: %v", bounds)
	}
}

func TestWindowAt(t *testing.T) {

	window, success := windowAt(image.ZP)
	if success {
		t.Errorf("window at %v should not exist, but got %v", image.ZP, window)
	}
}

func TestWindowAt2(t *testing.T) {

	window, _ := windowAt(image.Point{1, 1})
	neighbourCount := 8

	if len(window) != neighbourCount {
		t.Errorf("Expected %v image points, but got %v", neighbourCount, len(window))
	}
}
