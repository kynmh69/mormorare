package time

import (
	"encoding/json"
	"time"
)

type DateTime time.Time

var _ json.Unmarshaler = &DateTime{}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation(time.DateOnly, s, time.Local)
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}
