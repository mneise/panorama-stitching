package cornerdetection

import "image"
import "path/filepath"
import "testing"
import "math"

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
	neighbourCount := 9

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

func TestWindowCorrectPoints(t *testing.T) {
	padding := 1
	i := LoadImage(filepath.FromSlash("resources/images/5x5-clear.png"))

	correctPoints := []image.Point{
		image.Point{0, 2}, image.Point{1, 2}, image.Point{2, 2},
		image.Point{0, 1}, image.Point{1, 1}, image.Point{2, 1},
		image.Point{0, 0}, image.Point{1, 0}, image.Point{2, 0}}

	point := image.Point{1, 1}
	window, success := windowAt(point, i, padding)

	if !success {
		t.Errorf("Should successfully return window points for point %v", point)
	}

	for _, neighbourPoint := range correctPoints {
		if !Contains(window, neighbourPoint) {
			t.Errorf("Window doesn't contain point %v", neighbourPoint)
		}
	}
}

func TestSobelBase(t *testing.T) {

	i := LoadImage(filepath.FromSlash("resources/images/5x5-clear.png"))

	point := image.Point{0, 0}
	mag, success := Sobel(point, i)

	if success {
		t.Errorf("Expected no magnitude in 0,0, but got %v", mag)
	}
}

func TestSobelPositive(t *testing.T) {

	i := LoadImage(filepath.FromSlash("resources/images/5x5-centered-dot.png"))

	expectedMagnitude := math.Sqrt(math.Pow(-2*255, 2))
	point := image.Point{1, 2}
	mag, success := Sobel(point, i)

	if !success || mag != expectedMagnitude {
		t.Errorf("Expected magnitude %v in 1,1 but got %v", expectedMagnitude, mag)
	}
}

func TestSobelNegative(t *testing.T) {
	i := LoadImage(filepath.FromSlash("resources/images/5x5-clear.png"))
	bounds := i.Bounds()

	// top and bottom row
	for _, y := range []int{bounds.Min.Y, bounds.Max.Y} {
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			_, success := Sobel(image.Point{x, y}, i)
			if success {
				t.Errorf("Expected no magnitude, but got success for point %v", image.Point{x, y})
			}
		}
	}

	// left and right edge
	for _, x := range []int{bounds.Min.X, bounds.Max.X} {
		for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
			_, success := Sobel(image.Point{x, y}, i)
			if success {
				t.Errorf("Expected no magnitude, but got success for point %v", image.Point{x, y})
			}
		}
	}
}
