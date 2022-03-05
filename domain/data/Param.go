package data

type ParamType struct {
	Code  string
	Title string
	Alias string
}

func (ParamType) TableName() string {
	return "param_type"
}

//-------------------------------------------------------------------------------

type ParamCode struct {
	Code  string
	Title string
	Alias string
}

func (ParamCode) TableName() string {
	return "param_code"
}

//--------------------------------------------------------------------------------

type EntryParam struct {
	Id    string
	Code  ParamCode
	Type  ParamType
	Value string
	Extra string
}

type DbEntryParam struct {
	Id    string
	Code  string
	Type  string
	Value string
	Extra string
}

func (DbEntryParam) TableName() string {
	return "entry_param"
}

//--------------------------------------------------------------------------------

type ObjectParam struct {
	Id    string
	Code  ParamCode
	Type  ParamType
	Value string
	Extra string
}

type DbObjectParam struct {
	Id    string
	Code  string
	Type  string
	Value string
	Extra string
}

func (DbObjectParam) TableName() string {
	return "object_param"
}

//--------------------------------------------------------------------------------

type SourceParam struct {
	Id    string
	Code  ParamCode
	Type  ParamType
	Value string
	Extra string
}

type DbSourceParam struct {
	Id    string
	Code  string
	Type  string
	Value string
	Extra string
}

func (DbSourceParam) TableName() string {
	return "source_param"
}
