package level

import (
	"github.com/go-playground/validator/v10"
	"github.com/segmentio/encoding/json"
)

const (
	//OrdNotSelected is the order number for not selected level.
	OrdNotSelected ord = iota
	//OrdBeginner is the order number for the beginner level.
	OrdBeginner
	//OrdIntermediate is the order number for the intermediate level.
	OrdIntermediate
	//OrdAdvanced is the order number for the advanced level.
	OrdAdvanced
	//OrdMaster is the order number for the master level.
	OrdMaster
)

//MaxOrd is the maximal order number.
const MaxOrd = OrdMaster

//MaxNum is the maximal number value for order number.
const MaxNum = uint8(MaxOrd)

//Level represents level enum.
type Level struct {
	ordNum ord
}

//FromOrd returns the level for the given order number.
func FromOrd(ordNum ord) Level {
	return Level{ordNum}
}

//FromNum returns Level for the given number.
func FromNum(val uint8) Level {
	num := MaxOrd
	if val < MaxNum {
		num = ord(val)
	}
	return FromOrd(num)
}

//Ord returns the order number.
func (level Level) Ord() uint8 {
	return uint8(level.ordNum)
}

//MarshalJSON implements json.Marshaler.
func (level Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(level.ordNum)
}

//RegisterValidator registers field validator.
func RegisterValidator(validate *validator.Validate) {
	validate.RegisterValidation("level", func(fl validator.FieldLevel) bool {
		return fl.Field().Interface().(uint8) <= MaxNum
	})
}

type ord uint8
