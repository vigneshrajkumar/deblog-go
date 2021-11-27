package module

type Module struct {
	Name      string
	TableName string
	Fields    []*Field
}

func (mod *Module) GetIDColumn() *Field {
	for _, f := range mod.Fields {
		if f.Name == "id" {
			return f
		}
	}
	return nil
}
