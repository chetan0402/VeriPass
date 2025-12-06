package veripass

import (
	"time"

	"github.com/google/uuid"
)

func ToUUIDv7Nil(t time.Time) uuid.UUID {
	u := [16]byte{}

	milli := t.UnixMilli()

	u[0] = byte(milli >> 40)
	u[1] = byte(milli >> 32)
	u[2] = byte(milli >> 24)
	u[3] = byte(milli >> 16)
	u[4] = byte(milli >> 8)
	u[5] = byte(milli)

	u[6] = 0x70

	return u
}

func ToUUIDv7Max(t time.Time) uuid.UUID {
	u := ToUUIDv7Nil(t)

	u[6] = 0x7F
	for i := 7; i < 16; i++ {
		u[i] = 0xFF
	}

	return u
}
