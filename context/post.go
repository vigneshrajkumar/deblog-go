package context

import "fmt"

type Post struct {
	ID      int64
	Title   string
	Content string
	Author  int64
}

func (p Post) String() string {
	return fmt.Sprint("{", p.ID, ",", p.Title, "}")
}
