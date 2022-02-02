package unsigs

import (
	"reflect"
	"testing"
)

func Test_Check3x2(t *testing.T) {
	test3x2 := [6]uint16{1507, 144, 148, 6, 1380, 137}
	if !Check3x2(test3x2) {
		t.Fatalf("Expected 3x2: %v", test3x2)
	}

	test3x2 = [6]uint16{1507, 148, 1380, 144, 6, 137}
	if Check3x2(test3x2) {
		t.Fatalf("Unexpected 3x2: %v", test3x2)
	}
}

func Test_Find3x2s(t *testing.T) {
	horizontalPairs := [][2]uint16{{1507, 144}, {148, 6}, {148, 6}, {1380, 137}}
	expected := [][6]uint16{{1507, 144, 148, 6, 1380, 137}, {1507, 144, 1380, 137, 148, 6}}
	result, err := Find3x2s(horizontalPairs)
	if err != nil {
		t.Fatalf("Unexpected error %s:", err.Error())
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Found: %v. Expected: %v", result, expected)
	}
}
