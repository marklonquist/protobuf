package timestamp

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (m *Timestamp) MarshalJSON() ([]byte, error) {
	t, err := TS(m)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(t)
	return data, err
}

func (m *Timestamp) UnmarshalJSON(b []byte) error {
	var t time.Time
	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	m, err = TSProto(t)
	return err
}

const (
	// Seconds field of the earliest valid Timestamp.
	// This is time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix().
	minValidSeconds = -62135596800
	// Seconds field just after the latest valid Timestamp.
	// This is time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC).Unix().
	maxValidSeconds = 253402300800
)

func validateTimestamp(ts *Timestamp) error {
	if ts == nil {
		return errors.New("timestamp: nil Timestamp")
	}
	if ts.Seconds < minValidSeconds {
		return fmt.Errorf("timestamp: %v before 0001-01-01", ts)
	}
	if ts.Seconds >= maxValidSeconds {
		return fmt.Errorf("timestamp: %v after 10000-01-01", ts)
	}
	if ts.Nanos < 0 || ts.Nanos >= 1e9 {
		return fmt.Errorf("timestamp: %v: nanos not in range [0, 1e9)", ts)
	}
	return nil
}

func TS(ts *Timestamp) (time.Time, error) {
	// Don't return the zero value on error, because corresponds to a valid
	// timestamp. Instead return whatever time.Unix gives us.
	var t time.Time
	if ts == nil {
		t = time.Unix(0, 0).UTC() // treat nil like the empty Timestamp
	} else {
		t = time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
	}
	return t, validateTimestamp(ts)
}

func TSProto(t time.Time) (*Timestamp, error) {
	ts := &Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
	if err := validateTimestamp(ts); err != nil {
		return nil, err
	}
	return ts, nil
}
