package chain

import "deblog-go/command"

type SummaryPostChain struct {
	BaseChain
}

func GetSummaryPostChain() Chain {
	chn := SummaryPostChain{}
	chn.AddCommand(&command.SummaryPostCommand{})
	return &chn
}
