package data

type ContentType struct {
	Code  string
	Title string
	Alias string
}

func (ContentType) TableName() string {
	return "content_type"
}

//----------------------------------------------------------------------------------

type Content struct {
	Id     string
	Entry  string
	Title  string
	Type   ContentType
	Data   string
	Access int
}

type DbContent struct {
	Id     string
	Entry  string
	Title  string
	Type   string
	Data   string
	Access int
}

func (DbContent) TableName() string {
	return "content"
}
