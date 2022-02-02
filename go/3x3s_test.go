package unsigs

import (
	"reflect"
	"testing"
)

func Test_Check3x3(t *testing.T) {
	test3x3 := [9]uint16{13418, 14012, 13922, 13655, 13663, 13721, 13960, 13978, 13878}
	if !Check3x3(test3x3) {
		t.Fatalf("Expected 3x3: %v", test3x3)
	}

	test3x3 = [9]uint16{1507, 144, 662, 148, 6, 60, 0, 1, 2}
	if Check3x3(test3x3) {
		t.Fatalf("Unexpected 3x3: %v", test3x3)
	}
}

func Test_Find3x3s(t *testing.T) {
	horizontalPairs := [][2]uint16{{13418, 14012}, {14012, 13922}, {13655, 13663}, {13663, 13721}, {13960, 13978}, {13978, 13878}}
	expected := [][9]uint16{
		{13418, 14012, 13922, 13655, 13663, 13721, 13960, 13978, 13878},
		{13418, 14012, 13922, 13960, 13978, 13878, 13655, 13663, 13721},
		{13655, 13663, 13721, 13418, 14012, 13922, 13960, 13978, 13878},
		{13655, 13663, 13721, 13960, 13978, 13878, 13418, 14012, 13922},
		{13960, 13978, 13878, 13418, 14012, 13922, 13655, 13663, 13721},
		{13960, 13978, 13878, 13655, 13663, 13721, 13418, 14012, 13922},
	}
	result, err := Find3x3s(horizontalPairs)
	if err != nil {
		t.Fatalf("Unexpected error %s:", err.Error())
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Found: %v. Expected: %v", result, expected)
	}
}
