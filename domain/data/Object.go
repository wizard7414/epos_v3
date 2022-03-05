package data

import "time"

type Object struct {
	Id        string
	Added     time.Time
	Title     string
	Completed bool
	Ignored   bool
	Access    int
	Params    []ObjectParam
	Relations []Relation
}

type DbObject struct {
	Id        string
	Added     int64
	Title     string
	Completed int
	Ignored   int
	Access    int
}

func (DbObject) TableName() string {
	return "object"
}
