package cornerdetection

import "testing"
import "path/filepath"

func TestLoadImage(t *testing.T) {
	path := filepath.FromSlash("resources/images/")
	if image := LoadImage(path); image == nil {
		t.Errorf("LoadImage(%v) = %v, want image", path, image)
	}
}
