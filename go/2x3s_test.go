package unsigs

import (
	"reflect"
	"testing"
)

func Test_Check2x3(t *testing.T) {
	test2x3 := [6]uint16{1507, 144, 662, 148, 6, 60}
	if !Check2x3(test2x3) {
		t.Fatalf("Expected 2x3: %v", test2x3)
	}

	test2x3 = [6]uint16{1507, 148, 144, 6, 662, 60}
	if Check2x3(test2x3) {
		t.Fatalf("Unexpected 2x3: %v", test2x3)
	}
}

func Test_Find2x3s(t *testing.T) {
	verticalPairs := [][2]uint16{{1507, 148}, {144, 6}, {144, 6}, {662, 60}}
	expected := [][6]uint16{{1507, 144, 662, 148, 6, 60}}
	result, err := Find2x3s(verticalPairs)
	if err != nil {
		t.Fatalf("Unexpected error %s:", err.Error())
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Found: %v. Expected: %v", result, expected)
	}
}
