package priority

import (
	"github.com/go-playground/validator/v10"
	"github.com/segmentio/encoding/json"
)

const (
	//OrdNormal is the order number for the normal priority.
	OrdNormal ord = iota
	//OrdSecondary is the order number for the secondary priority.
	OrdSecondary
	//OrdMain is the order number for the main priority.
	OrdMain
)

//MaxOrd is the maximal order number.
const MaxOrd = OrdMain

//MaxNum is the maximal number value for order number.
const MaxNum = uint8(MaxOrd)

//Priority represents priority enum.
type Priority struct {
	ordNum ord
}

//FromOrd returns the priority for the given order number.
func FromOrd(ordNum ord) Priority {
	return Priority{ordNum}
}

//FromNum returns Priority for the given number.
func FromNum(val uint8) Priority {
	num := MaxOrd
	if val < MaxNum {
		num = ord(val)
	}
	return FromOrd(num)
}

//Ord returns the order number.
func (priority Priority) Ord() uint8 {
	return uint8(priority.ordNum)
}

//MarshalJSON implements json.Marshaler.
func (priority Priority) MarshalJSON() ([]byte, error) {
	return json.Marshal(priority.ordNum)
}

//RegisterValidator registers field validator.
func RegisterValidator(validate *validator.Validate) {
	validate.RegisterValidation("priority", func(fl validator.FieldLevel) bool {
		return fl.Field().Interface().(uint8) <= MaxNum
	})
}

type ord uint8
