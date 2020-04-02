package timestamp

import (
	"github.com/marklonquist/protobuf/jsonpb"
)

var marshaller jsonpb.Marshaler

func init() {
	marshaller = jsonpb.Marshaler{}
}

func (m *Timestamp) MarshalJSON() ([]byte, error) {
	data, err := marshaller.MarshalToString(m)
	if err != nil {
		return nil, err
	}
	return []byte(data), nil
}

func (m *Timestamp) UnmarshalJSON(b []byte) error {
	if len(b) == 2 {
		return nil
	}
	return jsonpb.UnmarshalString(string(b), m)
}
