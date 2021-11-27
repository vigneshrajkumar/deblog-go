package command

import (
	"database/sql"
	"deblog-go/context"
	"errors"
	"log"
)

type ListPostsCommand struct{}

func (ListPostsCommand) CommandName() string {
	return "ListPostsCommand"
}

func (ListPostsCommand) Exec(ctx map[string]interface{}) error {
	db, isDB := ctx["db"].(*sql.DB)
	if !isDB {
		return errors.New("db cxn not found")
	}

	rows, err := db.Query("SELECT ID, TITLE, CONTENT, AUTHOR FROM POST")
	if err != nil {
		return err
	}

	var posts []*context.Post

	for rows.Next() {
		var post context.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author); err != nil {
			return err
		}
		posts = append(posts, &post)
	}

	log.Println("fetched count: ", len(posts))
	ctx["list"] = posts
	return nil
}
