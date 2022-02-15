package global

import (
	"encoding/json"
	"errors"
	"strconv"
)

type SnowflakeID int64

func (u *SnowflakeID) UnmarshalJSON(b []byte) error {
	var i int64
	if err := json.Unmarshal(b, &i); err == nil {
		*u = SnowflakeID(i)
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return errors.New("expected a string or an integer")
	}
	if err := json.Unmarshal([]byte(s), &i); err != nil {
		return err
	}
	*u = SnowflakeID(i)
	return nil
}

func (a SnowflakeID) MarshalJSON() ([]byte, error) {
	temp := int64(a)
	return json.Marshal(strconv.FormatInt(temp, 10))
}
