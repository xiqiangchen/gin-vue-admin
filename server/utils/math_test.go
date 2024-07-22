package utils

import "testing"

func TestCeil2(t *testing.T) {
	t.Log(Ceil(123.4111567, 2))
	t.Log(Ceil(123.4567, 2))
	t.Log(Ceil(0.1234567, 4))
	t.Log(Ceil(123.4567, 0))
	t.Log(Ceil(123.4567, -1))
}
