package uuid

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
)

// UID represents unique identification for db entity
type UID string

// NewUUID creates new formatted uuid for use as db identification
func NewUUID() UID {
	uid := uuid.New()

	s := fmt.Sprintf("%x%x%x%x%x", uid[6:8], uid[4:6], uid[0:4], uid[8:10], uid[10:])

	return UID(s)
}

// Value encode unique identification for store to db
func (t UID) Value() (driver.Value, error) {
	res, err := hex.DecodeString(string(t))
	if err != nil {
		return []byte{}, err
	}
	return res, nil
}

// Scan decode unique identification to get from db
func (t *UID) Scan(src interface{}) error {
	s, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("error scan uid %v", src)
	}
	*t = UID(hex.EncodeToString(s))
	return nil
}
