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

	padding := 1
	i := LoadImage(filepath.FromSlash("resources/images/5x5-clear.png"))
	window, success := windowAt(image.ZP, i, padding)
	if success {
		t.Errorf("window at %v should not exist, but got %v", image.ZP, window)
	}
}

func TestWindowAtSize(t *testing.T) {

	padding := 1
	i := LoadImage(filepath.FromSlash("resources/images/5x5-clear.png"))
	window, _ := windowAt(image.Point{1, 1}, i, padding)
	neighbourCount := 8

	if len(window) != neighbourCount {
		t.Errorf("Expected %v image points, but got %v", neighbourCount, len(window))
	}
}

func TestWindowAtForEdgePoints(t *testing.T) {

	padding := 2
	i := LoadImage(filepath.FromSlash("resources/images/5x5-clear.png"))
	bounds := i.Bounds()

	// top and bottom row
	for _, y := range []int{bounds.Min.Y, bounds.Min.Y + 1, bounds.Max.Y, bounds.Max.Y - 1} {
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			window, success := windowAt(image.Point{x, y}, i, padding)
			if success || len(window) != 0 {
				t.Errorf("Expected 0 image points, but got %v", len(window))
			}
		}
	}

	// left and right edge
	for _, x := range []int{bounds.Min.X, bounds.Min.X + 1, bounds.Max.X, bounds.Max.X - 1} {
		for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
			window, success := windowAt(image.Point{x, y}, i, padding)

			if success || len(window) != 0 {
				t.Errorf("Expected 0 image points, but got %v", len(window))
			}
		}
	}
}
