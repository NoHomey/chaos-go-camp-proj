package level

const (
	//OrdBeginner is the order number for the beginner level.
	OrdBeginner ord = iota
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

//FromNum returns Rating for the given number.
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

type ord uint8
