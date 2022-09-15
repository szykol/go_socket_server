package message

import (
	"testing"
)

func Test_VectorPayload_ToBytes(t *testing.T) {
	ping := VectorPayload{
		X: 1,
		Y: 2,
	}

	buf, error := ping.ToBytes()
	if error != nil {
		t.Errorf("Received error when casting to bytes: %v", error)
	}

	resX, _ := DecodeValue[int16](buf[:4])
	if resX != 1 {
		t.Errorf("resX (%v) != 2", resX)
	}

	resY, _ := DecodeValue[int16](buf[4:])
	if resY != 2 {
		t.Errorf("resY (%v) != 2", resY)
	}
}
