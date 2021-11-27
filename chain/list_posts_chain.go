package chain

import "deblog-go/command"

type ListPostsChain struct {
	BaseChain
}

func GetListPostsChain() Chain {
	chn := ListPostsChain{}
	chn.AddCommand(&command.ListPostsCommand{})
	return &chn
}
