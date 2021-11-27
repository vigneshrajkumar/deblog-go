package chain

import "deblog-go/command"

type DeletePostChain struct {
	BaseChain
}

func GetDeletePostChain() Chain {
	chn := DeletePostChain{}
	chn.AddCommand(&command.DeletePostCommand{})
	return &chn
}
