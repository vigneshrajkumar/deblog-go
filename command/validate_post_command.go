package command

import (
	"deblog-go/context"
	"errors"
)

type ValidatePostCommand struct{}

func (ValidatePostCommand) CommandName() string {
	return "ValidatePostCommand"
}

func (ValidatePostCommand) Exec(ctx map[string]interface{}) error {
	post, isPost := ctx["post"].(*context.Post)
	if !isPost {
		return errors.New("not a post")
	}
	if post.Content == "" || post.Title == "" || post.Author == 0 {
		return errors.New("malformed post")
	}
	return nil
}
