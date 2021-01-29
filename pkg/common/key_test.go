package common

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	start := '0'
	end := '9'
	point := start
	for point <= end {
		fmt.Print((string(point)) + "', '")
		point = point + 1
	}

}

func TestGenKey(t *testing.T) {
	i := 0
	for i < 10000 {
		key := GenKey()
		t.Log(key)
		i = i + 1
	}
}
