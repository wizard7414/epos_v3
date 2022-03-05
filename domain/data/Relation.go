package data

type RelationType struct {
	Code  string
	Title string
	Alias string
}

func (RelationType) TableName() string {
	return "rel_type"
}

//--------------------------------------------------------------------------------

type Relation struct {
	Code        string
	Title       string
	Description string
	Object      string
	Type        RelationType
}

type DbRelation struct {
	Code        string
	Title       string
	Description string
	Type        string
}

func (DbRelation) TableName() string {
	return "relation"
}
