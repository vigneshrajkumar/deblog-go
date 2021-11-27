package chain

import "deblog-go/command"

type ListAuthorsChain struct {
	BaseChain
}

func GetListAuthorsChain() Chain {
	chn := ListAuthorsChain{}
	chn.AddCommand(&command.ListAuthorsCommand{})
	return &chn
}
