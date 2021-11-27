package command

import (
	"database/sql"
	"deblog-go/context"
	"errors"
	"log"
)

type ListAuthorsCommand struct{}

func (ListAuthorsCommand) CommandName() string {
	return "ListAuthorsCommand"
}

func (ListAuthorsCommand) Exec(ctx map[string]interface{}) error {
	db, isDB := ctx["db"].(*sql.DB)
	if !isDB {
		return errors.New("db cxn not found")
	}

	rows, err := db.Query("SELECT ID, NAME FROM AUTHOR")
	if err != nil {
		return err
	}

	var authors []*context.Author

	var author context.Author
	for rows.Next() {
		if err := rows.Scan(&author.ID, &author.Name); err != nil {
			return err
		}
		authors = append(authors, &author)
	}

	log.Println("fetched count: ", len(authors))
	ctx["list"] = authors
	return nil
}
