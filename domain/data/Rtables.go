package data

type DbRtEntryParam struct {
	Entry string
	Param string
}

func (DbRtEntryParam) TableName() string {
	return "r_entry_param"
}

//--------------------------------------------------------------------------------

type DbRtObjectParam struct {
	Object string
	Param  string
}

func (DbRtObjectParam) TableName() string {
	return "r_object_param"
}

//--------------------------------------------------------------------------------

type DbRtSourceParam struct {
	Source string
	Param  string
}

func (DbRtSourceParam) TableName() string {
	return "r_source_param"
}

//--------------------------------------------------------------------------------

type DbRtObjectRelation struct {
	ObjectA  string
	Relation string
	ObjectB  string
}

func (DbRtObjectRelation) TableName() string {
	return "r_object_relation"
}

//--------------------------------------------------------------------------------

type DbRtEntryObject struct {
	Entry  string
	Object string
}

func (DbRtEntryObject) TableName() string {
	return "r_entry_object"
}
