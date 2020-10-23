package common

import (
	"testing"
)

func TestHashIntersection(t *testing.T) {
	t.Run("Test hash intersection", func(t *testing.T) {
		first := []uint{1,2,3,4}
		second := []uint{1,4,454,5}

		intersection := HashIntersection(first, second)
		if len(intersection) != 2 {
			t.Error("Error in create intersections")
		}
	})
}
