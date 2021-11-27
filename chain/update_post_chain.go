package chain

import "deblog-go/command"

type UpdatePostChain struct {
	BaseChain
}

func GetUpdatePostChain() Chain {
	chn := UpdatePostChain{}
	chn.AddCommand(&command.UpdatePostCommand{})
	return &chn
}
