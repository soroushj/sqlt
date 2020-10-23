package sqlt

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// NullRawMessage represents a RawMessage that may be null.
// NullRawMessage implements the sql.Scanner interface so
// it can be used as a scan destination.
type NullRawMessage struct {
	RawMessage json.RawMessage
	Valid      bool // Valid is true if RawMessage is not NULL
}

// Scan implements the sql.Scanner interface.
func (nm *NullRawMessage) Scan(value interface{}) error {
	switch value := value.(type) {
	case nil:
		nm.RawMessage, nm.Valid = json.RawMessage{}, false
	case string:
		nm.RawMessage, nm.Valid = json.RawMessage(value), true
	case []byte:
		nm.RawMessage, nm.Valid = json.RawMessage(value), true
	default:
		return fmt.Errorf("Scan: unable to scan type %T into NullRawMessage", value)
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (nm NullRawMessage) Value() (driver.Value, error) {
	if !nm.Valid {
		return nil, nil
	}
	return []byte(nm.RawMessage), nil
}
