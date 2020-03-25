package timestamp

import (
	"encoding/json"
	"time"

	"github.com/marklonquist/protobuf/ptypes"
)

func (m *Timestamp) MarshalJSON() ([]byte, error) {
	t, err := ptypes.Timestamp(m)
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

	m, err = ptypes.TimestampProto(t)
	return err
}
