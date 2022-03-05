package data

type SourceType struct {
	Code  string
	Title string
	Alias string
}

func (SourceType) TableName() string {
	return "source_type"
}

//----------------------------------------------------------------------------------

type Source struct {
	Code   string
	Title  string
	Url    string
	Type   SourceType
	Params []SourceParam
}

type DbSource struct {
	Code  string
	Title string
	Url   string
	Type  string
}

func (DbSource) TableName() string {
	return "source"
}
