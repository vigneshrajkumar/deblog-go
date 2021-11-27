package module

func GetPostModule() *Module {
	return &Module{
		Name:      "post",
		TableName: "POST",
		Fields: []*Field{{
			Name:     "id",
			DBName:   "ID",
			Datatype: "INT",
		}, {
			Name:     "title",
			DBName:   "TITLE",
			Datatype: "VARCHAR(255)",
		}, {
			Name:     "content",
			DBName:   "CONTENT",
			Datatype: "TEXT",
		}, {
			Name:     "author",
			DBName:   "Author",
			Datatype: "INT",
		},
		},
	}
}
