package testme

import "testing"

func TestS1(t *testing.T) {
	if s1("123") != 3 {
		t.Error(`s1("123") 1 != 3`)
	}
	if s1("") != 0 {
		t.Error(`s1("") != 0`)
	}
}

func TestS2(t *testing.T) {
	if s2("123") != 3 {
		t.Error(`s1("123") 1 != 3`)
	}
	if s2("") != 0 {
		t.Error(`s1("") != 0`)
	}
}

func TestF1(t *testing.T) {
	if f1(0) != 0 {
		t.Error(`f1(0) != 0`)
	}
	if f1(1) != 1 {
		t.Error(`f1(1) != 1`)
	}
	if f1(2) != 1 {
		t.Error(`f1(2) != 1`)
	}
	if f1(10) != 55 {
		t.Error(`f1(10) != 55`)
	}
}
func TestF2(t *testing.T) {
	if f2(0) != 0 {
		t.Error(`f2(0) != 0`)
	}
	if f2(1) != 1 {
		t.Error(`f2(1) != 1`)
	}
	if f2(2) != 1 {
		t.Error(`f2(2) != 1`)
	}
	if f2(10) != 55 {
		t.Error(`f2(10) != 55`)
	}
}
