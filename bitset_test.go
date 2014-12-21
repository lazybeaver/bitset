package bitset

import (
	"testing"
)

func TestNew(t *testing.T) {
	size := 101
	bs := New(size)
	if bs.Size() != size {
		t.Errorf("Unexpected initialization failure")
	}
	for i := 0; i < size; i++ {
		if bs.Get(i) {
			t.Errorf("Newly initialized bitset cannot have true values")
		}
	}
}

func TestParse(t *testing.T) {
	expected := "10010010"
	bs, err := Parse(expected)
	if err != nil {
		t.Errorf("Unexpected parsing failure: %s in %s", err, expected)
	}
	actual := bs.String()
	if actual != expected {
		t.Errorf("Actual: %s | Expected: %s", actual, expected)
	}
}

func TestParseError(t *testing.T) {
	s := "1A"
	_, err := Parse(s)
	if err == nil {
		t.Errorf("Parsing of invalid string (%s) did not fail", s)
	}
}

func TestGet(t *testing.T) {
	bs := New(2)
	bs.Set(0)
	bs.Clear(1)
	if bs.Get(0) != true {
		t.Errorf("Actual: false | Expected: true")
	}
	if bs.Get(1) != false {
		t.Errorf("Actual: true | Expected: false")
	}
}

func TestSet(t *testing.T) {
	bs := New(10)
	bs.Set(2)
	bs.Set(3)
	bs.Set(5)
	bs.Set(7)
	actual := bs.String()
	expected := "0011010100"
	if actual != expected {
		t.Errorf("Actual: %s | Expected: %s", actual, expected)
	}
}

func TestSetAll(t *testing.T) {
	bs := New(10)
	bs.SetAll()
	actual := bs.String()
	expected := "1111111111"
	if actual != expected {
		t.Errorf("Actual: %s | Expected: %s", actual, expected)
	}
}

func TestInvert(t *testing.T) {
	bs := New(10)
	bs.SetAll()
	bs.Invert(2)
	bs.Invert(3)
	bs.Invert(5)
	bs.Invert(7)
	actual := bs.String()
	expected := "1100101011"
	if actual != expected {
		t.Errorf("Actual: %s | Expected: %s", actual, expected)
	}
}

func TestInvertAll(t *testing.T) {
	bs := New(10)
	bs.SetAll()
	bs.InvertAll()
	bs.InvertAll()
	actual := bs.String()
	expected := "1111111111"
	if actual != expected {
		t.Errorf("Actual: %s | Expected: %s", actual, expected)
	}
}

func TestClear(t *testing.T) {
	bs := New(10)
	bs.SetAll()
	bs.Clear(0)
	bs.Clear(3)
	bs.Clear(6)
	bs.Clear(9)
	actual := bs.String()
	expected := "0110110110"
	if actual != expected {
		t.Errorf("Actual: %s | Expected: %s", actual, expected)
	}
}

func TestClearAll(t *testing.T) {
	bs := New(10)
	bs.SetAll()
	bs.ClearAll()
	actual := bs.String()
	expected := "0000000000"
	if actual != expected {
		t.Errorf("Actual: %s | Expected: %s", actual, expected)
	}
}

func BenchmarkGet(b *testing.B) {
	bs := New(b.N)
	for i := 0; i < b.N; i++ {
		_ = bs.Get(i)
	}
}

func BenchmarkSet(b *testing.B) {
	bs := New(b.N)
	for i := 0; i < b.N; i++ {
		bs.Set(i)
	}
}

func BenchmarkInvert(b *testing.B) {
	bs := New(b.N)
	for i := 0; i < b.N; i++ {
		bs.Invert(i)
	}
}
