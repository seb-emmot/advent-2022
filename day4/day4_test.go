package day4

import "testing"

func TestContains(t *testing.T) {
	var x Span = Span{1, 5}
	var y Span = Span{2, 3}
	expected := true
	actual := x.Contains(y)

	if expected != actual {
		t.Errorf("Expected %t, got %t", expected, actual)
	}
}

func TestOverlaps(t *testing.T) {
	var x Span = Span{1, 5}
	var y Span = Span{5, 8}
	expected := true
	actual := x.Overlaps(y)

	if expected != actual {
		t.Errorf("Expected %t, got %t", expected, actual)
	}
}

func TestCheckContains(t *testing.T) {
	var spans []Span = []Span{
		{1, 5}, {2, 3},
		{10, 13}, {10, 15}}
	expected := []bool{true, true}
	actual := checkContains(spans)

	for i, v := range expected {
		if v != actual[i] {
			t.Errorf("Expected %t, got %t", v, actual[i])
		}
	}
}
