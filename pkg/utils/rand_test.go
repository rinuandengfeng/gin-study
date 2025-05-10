package utils

import (
	"testing"
)

func TestRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		t.Log(RandInt(10, 13))
	}
}

func TestRandNum(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandNum(4))
	}
}

func TestRandStr(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandStr(6))
	}
}
