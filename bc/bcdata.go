package bc

import (
	"strconv"
	"time"
)

// A simple data, ie: Car Ownership
type CarData struct {
	ChangeDate time.Time
	Owner      string
	Value      int
}

func (c *CarData) ToString() string {
	return "Car owned by: " + c.Owner + "\t Value:" + strconv.Itoa(c.Value) + " EUR\t TRX:" + c.ChangeDate.Format(time.RFC3339)
}

// Hashed data
type HashedCarData struct {
	previousHash string
	data         CarData
	currentHash  string
}

func NewBcData(changeDate time.Time,
	owner string,
	value int) *CarData {
	t := &CarData{}
	t.ChangeDate = changeDate
	t.Owner = owner
	t.Value = value
	return t
}
