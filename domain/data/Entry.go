package data

import "time"

type Entry struct {
	Id        string
	Added     time.Time
	Title     string
	URL       string
	Completed bool
	Ignored   bool
	Access    int
	Params    []EntryParam
	Content   []Content
	Objects   []Object
}

type DbEntry struct {
	Id        string
	Added     int64
	Title     string
	URL       string
	Completed int
	Ignored   int
	Access    int
}

func (DbEntry) TableName() string {
	return "entry"
}
