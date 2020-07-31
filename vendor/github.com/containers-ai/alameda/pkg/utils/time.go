package utils

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func Timestamp(timestamp *timestamp.Timestamp) *time.Time {
	ts, _ := ptypes.Timestamp(timestamp)
	return &ts
}

func TimestampProto(ts *time.Time) *timestamp.Timestamp {
	var t *timestamp.Timestamp
	if ts == nil {
		t, _ = ptypes.TimestampProto(time.Unix(0, 0).UTC())
	} else {
		t, _ = ptypes.TimestampProto(*ts)
	}
	return t
}
