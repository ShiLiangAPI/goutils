package types

import (
	"database/sql/driver"
	"encoding/json"
)

type SliceInt64Json []string

func (obj *SliceInt64Json) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, obj)
}

func (obj SliceInt64Json) Value() (driver.Value, error) {
	return json.Marshal(obj)
}
