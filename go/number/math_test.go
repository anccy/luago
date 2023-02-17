package number

import (
	"testing"
)

func TestIFloorDiv(t *testing.T) {
	ret := IFloorDiv(10, -4)
	if ret != -3 {
		t.Error("floor div err")
	}
}

func TestFFloorDiv(t *testing.T) {
	ret := FFloorDiv(10.0, -4.0)
	if ret != -3.0 {
		t.Errorf("floor div err, want -2, ret %v", ret)
	}
}
