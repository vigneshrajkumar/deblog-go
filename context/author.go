package context

import "fmt"

type Author struct {
	ID   int64
	Name string
}

func (a Author) String() string {
	return fmt.Sprint("{", a.ID, ",", a.Name, "}")
}
