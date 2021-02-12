package common

import (
	"time"
)

var Chars = []byte{
	'4', 't', 'u', 'X', 'i', 'j', 'k', 'H', '2', '3', 'r', 's', 'l', 'm', 'n', 'o',
	'A', 'z', '0', '1', 'd', 'K', 'L', 'Q', 'R', 'S', 'T', 'U', 'M', 'v', 'J', '5', 'B', 'C', 'e', '8', '9', 'D', 'E', 'I', 'P',
	'p', 'q', 'y', '6', 'a', 'N', 'O', 'Y', 'Z', '_', 'f', 'g', 'h', 'b', 'V', 'W', 'F', 'G', 'c', '7', 'w', 'x',
}

func GenKey() string {
	key := ""
	ss := time.Now().UnixNano()
	size := int64(len(Chars))
	s := ss
	a := s
	for a > size {
		key = string(Chars[a%size]) + key
		a = a / size
	}
	key = string(Chars[a%size]) + key
	return key
}
