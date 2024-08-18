package hashdefs

import (
    "testing"
)

func TestMapCrcs(t *testing.T) {
	crcmap := MapCrcs()
	if crcmap == nil {
		t.Fatalf(`crcmap is nil`)
	}
}
