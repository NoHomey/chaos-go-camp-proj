package uuid

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	impl "github.com/google/uuid"
)

//UUID wraps UUID implementation for using as value for MySQL column.
type UUID struct {
	impl.UUID
}

//Value implements the driver.Valuer interface.
func (id UUID) Value() (driver.Value, error) {
	b := [16]byte(id.UUID)
	return b[:], nil
}

//Scan implements the sql.Scanner interface.
func (id *UUID) Scan(src interface{}) error {
	if src == nil {
		return ErrNil{}
	}
	switch val := src.(type) {
	case string:
		parsed, err := Parse(val)
		if err != nil {
			return err
		}
		*id = parsed
		return nil
	case []byte:
		parsed, err := impl.FromBytes(val)
		if err != nil {
			return err
		}
		*id = UUID{parsed}
		return nil
	}
	return ErrInvalType{src}
}

//ErrInvalStr is error for invalid UUID strings.
type ErrInvalStr struct {
	Str string
	Err error
}

func (err ErrInvalStr) Error() string {
	msg := "%s is invalid UUID string, error: %s"
	return fmt.Sprintf(msg, err.Str, err.Err.Error())
}

//ErrNil is used for reporting nil value beein scanned.
type ErrNil struct{}

func (ErrNil) Error() string {
	return "Scan error: got nil value. Value of type UUID can not be NULL"
}

//ErrInvalType is used for reporting that value of wrong type is beeing scaned.
type ErrInvalType struct {
	val interface{}
}

func (err ErrInvalType) Error() string {
	msg := "Scan error: got value: %v of type %T. Scaned value for UUID should be of type []byte or string"
	return fmt.Sprintf(msg, err.val, err.val)
}

//Parse validates and parses UUID strings.
func Parse(s string) (UUID, error) {
	var zero uuid.UUID
	parsed, err := uuid.Parse(s)
	if err != nil {
		return UUID{zero}, ErrInvalStr{s, err}
	}
	return UUID{parsed}, nil
}
