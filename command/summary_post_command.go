package command

import (
	"database/sql"
	"deblog-go/context"
	"errors"
	"log"
)

type SummaryPostCommand struct{}

func (SummaryPostCommand) CommandName() string {
	return "SummaryPostCommand"
}

func (SummaryPostCommand) Exec(ctx map[string]interface{}) error {
	db, isDB := ctx["db"].(*sql.DB)
	if !isDB {
		return errors.New("db cxn not found")
	}

	id, isID := ctx["id"].(string)
	if !isID {
		return errors.New("no id found")
	}

	stmt := "SELECT ID, TITLE, CONTENT, AUTHOR FROM POST WHERE ID = ?"
	args := []interface{}{id}
	log.Println(stmt, args)
	rows, err := db.Query(stmt, args...)
	if err != nil {
		return err
	}

	var post context.Post

	for rows.Next() {
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author); err != nil {
			return err
		}
	}

	ctx["summary"] = post
	return nil
}
