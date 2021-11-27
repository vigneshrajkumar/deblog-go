package command

import (
	"deblog-go/context"
)

type ConvertJsonToPostCommand struct{}

func (ConvertJsonToPostCommand) CommandName() string {
	return "ConvertJsonToPostCommand"
}

func (ConvertJsonToPostCommand) Exec(ctx map[string]interface{}) error {

	reqMap := ctx["data"].(map[string]interface{})

	post := context.Post{
		Title:   reqMap["title"].(string),
		Content: reqMap["content"].(string),
		Author:  int64(reqMap["author"].(float64)),
	}
	ctx["post"] = &post
	return nil
}
