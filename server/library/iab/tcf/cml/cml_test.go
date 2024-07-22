package cml

import "testing"

func Test_GetCML(t *testing.T) {
	cml := getLastCML()

	t.Log(cml.Cmps["10"].Environments)
	t.Log(cml)
}
