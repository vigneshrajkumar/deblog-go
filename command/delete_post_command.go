package command

import (
	"database/sql"
	"deblog-go/module"
	"errors"
	"log"
	"strconv"
)

type DeletePostCommand struct{}

func (DeletePostCommand) CommandName() string {
	return "DeletePostCommand"
}

func (DeletePostCommand) Exec(ctx map[string]interface{}) error {
	db, isDB := ctx["db"].(*sql.DB)
	if !isDB {
		return errors.New("db cxn not found")
	}

	idStr, isID := ctx["id"].(string)
	if !isID {
		return errors.New("id not found")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.New("id not found")
	}
	mod := module.GetPostModule()
	stmt := "DELETE FROM " + mod.TableName + " WHERE " + mod.GetIDColumn().Name + " = ?"
	log.Println(stmt, id)
	res, err := db.Exec(stmt, id)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("deleted rows: ", ra)

	ctx["status"] = "deletion success"
	return nil
}
