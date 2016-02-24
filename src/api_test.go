package main

import (
	"testing"
)

func TestEcho(t *testing.T) {
	res, err := echo(nil)
	if err != nil {
		t.Error("Should not error on echo.")
	} else if res != "Hello" {
		t.Error("Bad Echo!")
	}
}

