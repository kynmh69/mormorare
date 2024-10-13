package time

import (
	"encoding/json"
	"time"
)

type DateTime time.Time

func (t *DateTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	pt, err := time.ParseInLocation(time.DateOnly, s, time.Local)
	if err != nil {
		return err
	}
	*t = DateTime(pt)
	return nil
}
