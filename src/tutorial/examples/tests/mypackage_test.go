package gotests

import (
	"testing"
)

func TestExportedFunc(t *testing.T) {
	result := MyExportedFunction("test")

	if result != "OK" {
		t.Errorf("OK expected !")
	}
}
