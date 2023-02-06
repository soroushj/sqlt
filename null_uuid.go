package sqlt

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
)

// NullUUID represents a UUID that may be null.
// NullUUID implements the sql.Scanner interface so
// it can be used as a scan destination.
//
// Note: [github.com/google/uuid.NullUUID] is available since v1.3.0.
type NullUUID struct {
	UUID  uuid.UUID
	Valid bool // Valid is true if UUID is not NULL
}

// Scan implements the sql.Scanner interface.
func (nu *NullUUID) Scan(value interface{}) error {
	switch value := value.(type) {
	case nil:
		nu.UUID, nu.Valid = uuid.UUID{}, false
	case string:
		// If the value is empty, return a null UUID
		if value == "" {
			nu.UUID, nu.Valid = uuid.UUID{}, false
			return nil
		}
		// Parse the string
		u, err := uuid.Parse(value)
		if err != nil {
			return fmt.Errorf("Scan: %v", err)
		}
		nu.UUID, nu.Valid = u, true
	case []byte:
		// If the value is empty, return a null UUID
		if len(value) == 0 {
			nu.UUID, nu.Valid = uuid.UUID{}, false
			return nil
		}
		// Assume a simple slice of bytes if 16 bytes
		// Otherwise, parse as string
		if len(value) == 16 {
			copy(nu.UUID[:], value)
			nu.Valid = true
			return nil
		}
		u, err := uuid.Parse(string(value))
		if err != nil {
			return fmt.Errorf("Scan: %v", err)
		}
		nu.UUID, nu.Valid = u, true
	default:
		return fmt.Errorf("Scan: unable to scan type %T into NullUUID", value)
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (nu NullUUID) Value() (driver.Value, error) {
	if !nu.Valid {
		return nil, nil
	}
	return nu.UUID.String(), nil
}
