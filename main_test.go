package main

import (
	"reflect"
	"testing"
)

func TestEqualPalyers(t *testing.T) {
	expected := Player{
		Name: "John",
		Hp:   100,
	}
	have := Player{
		Name: "John2",
		Hp:   100,
	}

	if !reflect.DeepEqual(expected, have) {
		t.Errorf("expected %v, have %v", expected, have)
	}

}

func TestCalculateValues(t *testing.T) {
	var (
		expected = 10
		a        = 5
		b        = 5
	)
	have := calculateValues(a, b)
	if have != expected {
		t.Errorf("expected %d, have %d", expected, have)
	}

}
