package command

import (
	"database/sql"
	"deblog-go/module"
	"errors"
	"log"
	"reflect"
	"strings"
)

type UpdatePostCommand struct{}

func (UpdatePostCommand) CommandName() string {
	return "UpdatePostCommand"
}

func (UpdatePostCommand) Exec(ctx map[string]interface{}) error {
	db, isDB := ctx["db"].(*sql.DB)
	if !isDB {
		return errors.New("db cxn not found")
	}

	reqMap := ctx["data"].(map[string]interface{})

	log.Println("reqMap:", reqMap)

	log.Println("id rtype:", reflect.TypeOf(reqMap["id"]))

	id, isID := reqMap["id"].(float64)
	if !isID {
		return errors.New("id not found")
	}

	mod := module.GetPostModule()

	stmt := "UPDATE " + mod.TableName + " SET "
	args := make([]interface{}, 0)
	setTuples := make([]string, 0)
	for _, field := range mod.Fields {
		if reqMap[field.Name] != nil {
			setTuples = append(setTuples, field.Name+" = ? ")
			args = append(args, reqMap[field.Name])
		}
	}

	stmt += strings.Join(setTuples, ", ")
	stmt += " WHERE " + mod.GetIDColumn().DBName + " = ? "
	args = append(args, id)

	log.Println(stmt, args)
	res, err := db.Exec(stmt, args...)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("updated rows: ", ra)

	ctx["status"] = "updation success"
	return nil
}
