package time

import (
	"encoding/json"
	"github.com/kynmh69/mormorare/consts"
	"time"
)

type DateTime time.Time

func (t *DateTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	pt, err := time.ParseInLocation(consts.DateFormat, s, time.Local)
	if err != nil {
		return err
	}
	*t = DateTime(pt)
	return nil
}
