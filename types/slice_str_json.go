package types

import (
	"database/sql/driver"
	"encoding/json"
)

type SliceStrJson []string

func (obj *SliceStrJson) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, obj)
}

func (obj SliceStrJson) Value() (driver.Value, error) {
	return json.Marshal(obj)
}
