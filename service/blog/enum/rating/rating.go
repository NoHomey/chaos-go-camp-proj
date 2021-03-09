package rating

import "encoding/json"

//Ords are the availible order numbers.
const (
	Ord0 ord = iota
	Ord1
	Ord2
	Ord3
	Ord4
	Ord5
	Ord6
	Ord7
	Ord8
	Ord9
	Ord10
	Ord11
	Ord12
	Ord13
	Ord14
	Ord15
)

//MaxOrd is the maximal order number.
const MaxOrd = Ord15

//MaxNum is the maximal number value for order number.
const MaxNum = uint8(MaxOrd)

//Rating represents rating enum.
type Rating struct {
	ordNum ord
}

//FromOrd returns Rating for the given order number.
func FromOrd(ordNum ord) Rating {
	return Rating{ordNum}
}

//FromNum returns Rating for the given number.
func FromNum(val uint8) Rating {
	num := MaxOrd
	if val < MaxNum {
		num = ord(val)
	}
	return FromOrd(num)
}

//Ord returns the order number.
func (rating Rating) Ord() uint8 {
	return uint8(rating.ordNum)
}

//MarshalJSON implements json.Marshaler
func (rating Rating) MarshalJSON() ([]byte, error) {
	return json.Marshal(rating.ordNum)
}

type ord uint8
