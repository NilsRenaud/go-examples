package gotests

import (
	"testing"
)

func TestExportedFunc(t *testing.T) { // We will see * later (pointer !!)
	result := MyExportedFunction("test")

	if result != "OK" {
		t.Errorf("OK expected !")
	}
}

func TestPrivateFunc(t *testing.T) { // We can even test private functions !
	result := myPrivateFunction("test")

	if result != "OK" {
		t.Errorf("OK expected !")
	}
}
