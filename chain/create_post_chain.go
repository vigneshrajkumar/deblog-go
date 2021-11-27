package chain

import "deblog-go/command"

type SavePostChain struct {
	BaseChain
}

func GetCreatePostChain() Chain {
	chn := SavePostChain{}
	chn.AddCommand(&command.ConvertJsonToPostCommand{})
	chn.AddCommand(&command.ValidatePostCommand{})
	chn.AddCommand(&command.CreatePostCommand{})
	return &chn
}
