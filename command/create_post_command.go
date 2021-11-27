package command

import (
	"database/sql"
	"deblog-go/context"
	"errors"
	"log"
)

type CreatePostCommand struct{}

func (CreatePostCommand) CommandName() string {
	return "CreatePostCommand"
}

func (CreatePostCommand) Exec(ctx map[string]interface{}) error {
	db, isDB := ctx["db"].(*sql.DB)
	if !isDB {
		return errors.New("db cxn not found")
	}

	post, isPost := ctx["post"].(*context.Post)
	if !isPost {
		return errors.New("post not found")
	}

	res, err := db.Exec("INSERT INTO POST (TITLE, CONTENT, AUTHOR) VALUES (?, ?, ?)", post.Title, post.Content, post.Author)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("inserted rows: ", ra)
	return nil
}
